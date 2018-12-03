package masterdns

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apiextclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	clusterclientset "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	clusterinformers "sigs.k8s.io/cluster-api/pkg/client/informers_generated/externalversions"

	operatorv1 "github.com/openshift/api/operator/v1"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/resource/resourcemerge"
	"github.com/openshift/library-go/pkg/operator/resource/resourceread"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	masterdnsv1 "github.com/openshift/master-dns-operator/pkg/apis/masterdns/v1alpha1"
	"github.com/openshift/master-dns-operator/pkg/operator/assets"
)

const (
	clusterAPINamespace      = "openshift-cluster-api"
	operatorNamespace        = "openshift-master-dns-operator"
	workloadFailingCondition = "WorkloadFailing"
)

var (
	configName      = types.NamespacedName{Name: "instance", Namespace: ""}
	requeueInterval = 10 * time.Second
)

// Add creates a new MasterDNS Operator and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	log.Info("Adding MasterDNS operator to manager")
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMasterDNS{
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Ensure that a config instance exists
	dynamicClient, err := dynamic.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	v1helpers.EnsureOperatorConfigExists(
		dynamicClient,
		assets.MustAsset("config/operator-config.yaml"),
		schema.GroupVersionResource{Group: masterdnsv1.SchemeGroupVersion.Group, Version: "v1alpha1", Resource: "masterdnsoperatorconfigs"},
	)
	mdnsReconciler := r.(*ReconcileMasterDNS)

	mdnsReconciler.imagePullSpec = os.Getenv("IMAGE")
	if len(mdnsReconciler.imagePullSpec) == 0 {
		return fmt.Errorf("no image specified, specify one via the IMAGE env var")
	}

	mdnsReconciler.kubeClient, err = kubeclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	mdnsReconciler.apiExtClient, err = apiextclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}

	mdnsReconciler.eventRecorder = setupEventRecorder(mdnsReconciler.kubeClient)

	// Create a new controller
	c, err := controller.New("master-dns-operator", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MasterDNSOperatorConfig
	err = c.Watch(&source.Kind{Type: &masterdnsv1.MasterDNSOperatorConfig{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	clusterAPIClient, err := clusterclientset.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	clusterAPIInformerFactory := clusterinformers.NewSharedInformerFactoryWithOptions(clusterAPIClient, 10*time.Minute, clusterinformers.WithNamespace(clusterAPINamespace))
	machineInformer := clusterAPIInformerFactory.Cluster().V1alpha1().Machines().Informer()
	mgr.Add(&informerRunnable{informer: machineInformer})

	// Watch for changes to machines in cluster API namespace
	err = c.Watch(&source.Informer{Informer: machineInformer}, &handler.EnqueueRequestsFromMapFunc{ToRequests: configHandlerFunc})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMasterDNS{}

// ReconcileMasterDNS reconciles a MasterDNSOperatorConfig object
type ReconcileMasterDNS struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client        client.Client
	scheme        *runtime.Scheme
	kubeClient    kubeclient.Interface
	apiExtClient  apiextclient.Interface
	eventRecorder events.Recorder
	imagePullSpec string
}

// Reconcile reads that state of the cluster for a MasterDNSOperatorConfig object and makes changes based on the state read
// and what is in the MasterDNSOperatorConfig.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMasterDNS) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	log.Printf("Reconciling MasterDNS %s\n", request.Name)

	// Fetch the MasterDNSOperatorConfig instance
	operatorConfig := &masterdnsv1.MasterDNSOperatorConfig{}
	err := r.client.Get(context.TODO(), request.NamespacedName, operatorConfig)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Warning("Master DNS operator config not found. Probing until one is created.")
			// Config not found. Keep retrying until one exists.
			return reconcile.Result{Requeue: true, RequeueAfter: 2 * time.Minute}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Ensure that External DNS is deployed
	return r.syncExternalDNS(operatorConfig)
}

func (r *ReconcileMasterDNS) syncExternalDNS(config *masterdnsv1.MasterDNSOperatorConfig) (reconcile.Result, error) {

	errors := []error{}

	crd := resourceread.ReadCustomResourceDefinitionV1Beta1OrDie(assets.MustAsset("config/crd.yaml"))
	_, _, err := resourceapply.ApplyCustomResourceDefinition(r.apiExtClient.ApiextensionsV1beta1(), r.eventRecorder, crd)

	results := resourceapply.ApplyDirectly(r.kubeClient, r.eventRecorder, assets.Asset,
		"config/role.yaml",
		"config/binding.yaml",
		"config/sa.yaml",
	)
	resourcesThatForceRedeployment := sets.NewString("config/sa.yaml")
	forceDeployment := config.Generation != config.Status.ObservedGeneration

	for _, currResult := range results {
		if currResult.Error != nil {
			log.WithError(currResult.Error).WithField("file", currResult.File).Error("Apply error")
			errors = append(errors, fmt.Errorf("%q (%T): %v", currResult.File, currResult.Type, currResult.Error))
			continue
		}

		if currResult.Changed && resourcesThatForceRedeployment.Has(currResult.File) {
			forceDeployment = true
		}
	}
	actualDeployment, _, err := r.syncExternalDNSDeployment(config, forceDeployment)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "deployments", err))
	}

	config.Status.ObservedGeneration = config.ObjectMeta.Generation
	resourcemerge.SetDeploymentGeneration(&config.Status.Generations, actualDeployment)
	if len(errors) > 0 {
		message := ""
		for _, err := range errors {
			message = message + err.Error() + "\n"
		}
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:    workloadFailingCondition,
			Status:  operatorv1.ConditionTrue,
			Message: message,
			Reason:  "SyncError",
		})
	} else {
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:   workloadFailingCondition,
			Status: operatorv1.ConditionFalse,
		})
	}
	if err := r.client.Status().Update(context.TODO(), config); err != nil {
		return reconcile.Result{}, err
	}

	if len(errors) > 0 {
		return reconcile.Result{Requeue: true, RequeueAfter: requeueInterval}, nil
	}
	return reconcile.Result{}, nil
}

func (r *ReconcileMasterDNS) syncExternalDNSDeployment(config *masterdnsv1.MasterDNSOperatorConfig, forceDeployment bool) (*appsv1.Deployment, bool, error) {
	deployment := resourceread.ReadDeploymentV1OrDie(assets.MustAsset("config/deployment.yaml"))
	deployment.Spec.Template.Spec.Containers[0].Image = r.imagePullSpec
	deployment.Spec.Template.Spec.Containers[0].ImagePullPolicy = corev1.PullAlways
	return resourceapply.ApplyDeployment(
		r.kubeClient.AppsV1(),
		r.eventRecorder,
		deployment,
		resourcemerge.ExpectedDeploymentGeneration(deployment, config.Status.Generations),
		forceDeployment)
}

type configHandler func(handler.MapObject) []reconcile.Request

func (r configHandler) Map(o handler.MapObject) []reconcile.Request {
	return r(o)
}

var configHandlerFunc configHandler = func(o handler.MapObject) []reconcile.Request {
	return []reconcile.Request{
		{
			NamespacedName: configName,
		},
	}
}

type informerRunnable struct {
	informer cache.SharedIndexInformer
}

func (r *informerRunnable) Start(stopch <-chan struct{}) error {
	r.informer.Run(stopch)
	cache.WaitForCacheSync(stopch, r.informer.HasSynced)
	return nil
}

func getLogRecorder() events.Recorder {
	return &logRecorder{}
}

type logRecorder struct{}

func (logRecorder) Event(reason, message string) {
	log.WithField("reason", reason).Info(message)
}

func (logRecorder) Eventf(reason, messageFmt string, args ...interface{}) {
	log.WithField("reason", reason).Infof(messageFmt, args...)
}

func (logRecorder) Warning(reason, message string) {
	log.WithField("reason", reason).Warning(message)
}

func (logRecorder) Warningf(reason, messageFmt string, args ...interface{}) {
	log.WithField("reason", reason).Warningf(messageFmt, args...)
}

func setupEventRecorder(kubeClient kubeclient.Interface) events.Recorder {
	controllerRef, err := events.GetControllerReferenceForCurrentPod(kubeClient, operatorNamespace, nil)
	if err != nil {
		log.WithError(err).Warning("Cannot determine pod name for event recorder. Using logger.")
		return getLogRecorder()
	}

	eventsClient := kubeClient.CoreV1().Events(controllerRef.Namespace)
	return events.NewRecorder(eventsClient, "openshift-master-dns-operator", controllerRef)
}
