package masterdns

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
	apiextclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/resource/resourceread"
	"github.com/openshift/library-go/pkg/operator/v1alpha1helpers"

	masterdnsv1 "github.com/openshift/master-dns-operator/pkg/apis/masterdns/v1alpha1"
	"github.com/openshift/master-dns-operator/pkg/operator/assets"
)

const (
	clusterAPINamespace = "openshift-cluster-api"
)

var (
	configName = types.NamespacedName{Name: "instance", Namespace: ""}
)

// Add creates a new MasterDNS Operator and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	log.Info("Adding MasterDNS operator to manager")
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMasterDNS{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Ensure that a config instance exists
	dynamicClient, err := dynamic.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	v1alpha1helpers.EnsureOperatorConfigExists(
		dynamicClient,
		assets.MustAsset("config/operator-config.yaml"),
		schema.GroupVersionResource{Group: masterdnsv1.SchemeGroupVersion.Group, Version: "v1alpha1", Resource: "masterdnsoperatorconfigs"},
		v1alpha1helpers.GetImageEnv,
	)

	r.(*ReconcileMasterDNS).kubeClient, err = kubeclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}
	r.(*ReconcileMasterDNS).apiExtClient, err = apiextclient.NewForConfig(mgr.GetConfig())
	if err != nil {
		return err
	}

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
	client       client.Client
	scheme       *runtime.Scheme
	kubeClient   kubeclient.Interface
	apiExtClient apiextclient.Interface
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
	err = r.ensureExternalDNS(operatorConfig)
	if err != nil {
		log.WithError(err).Error("Cannot ensure external DNS deployment is created")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileMasterDNS) ensureExternalDNS(config *masterdnsv1.MasterDNSOperatorConfig) error {
	eventsClient := r.kubeClient.CoreV1().Events(config.Namespace)
	recorder := events.NewRecorder(
		eventsClient,
		"openshift-master-dns-operator",
		&corev1.ObjectReference{
			Name:      config.Name,
			Namespace: config.Namespace})

	// Apply CRD
	crd := resourceread.ReadCustomResourceDefinitionV1Beta1OrDie(assets.MustAsset("config/01_dns_crd.yaml"))
	_, crdModified, err := resourceapply.ApplyCustomResourceDefinition(r.apiExtClient.ApiextensionsV1beta1(), recorder, crd)
	if err != nil {
		log.WithError(err).Error("error applying DNS crd")
		return err
	}

	// Apply Cluster Role
	role := resourceread.ReadClusterRoleV1OrDie(assets.MustAsset("config/02_dns_rbac.yaml"))
	_, clusterRoleModified, err := resourceapply.ApplyClusterRole(r.kubeClient.RbacV1(), recorder, role)
	if err != nil {
		log.WithError(err).Error("error appplying DNS cluster role")
		return err
	}

	// Apply Service Account
	sa := resourceread.ReadServiceAccountV1OrDie(assets.MustAsset("config/03_dns_sa.yaml"))
	_, saModified, err := resourceapply.ApplyServiceAccount(r.kubeClient.CoreV1(), recorder, sa)
	if err != nil {
		log.WithError(err).Error("error appplying DNS service account")
		return err
	}

	// Apply Cluster Role Binding
	binding := resourceread.ReadClusterRoleBindingV1OrDie(assets.MustAsset("config/04_dns_binding.yaml"))
	_, roleBindingModified, err := resourceapply.ApplyClusterRoleBinding(r.kubeClient.RbacV1(), recorder, binding)
	if err != nil {
		log.WithError(err).Error("error appplying DNS cluster role binding")
		return err
	}

	forceDeployment := config.Generation != config.Status.ObservedGeneration
	if crdModified || clusterRoleModified || saModified || roleBindingModified {
		forceDeployment = true
	}

	_ = forceDeployment

	// Apply External DNS deployment
	/*
		deployment := resourceread.ReadDeploymentV1OrDie(assets.MustAsset("config/04_dns_deployment.yaml"))
		deployment.Spec.Template.Spec.Containers[0].Image = config.Spec.ImagePullSpec
		deployment.Spec.Template.Spec.Containers[0].ImagePullPolicy = corev1.PullPolicy(config.Spec.ImagePullPolicy)
		_, _, err = resourceapply.ApplyDeployment(r.kubeClient.AppsV1().Deployments(config.Namespace), recorder, deployment)
	*/

	return nil
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
