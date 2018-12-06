package masterdns

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"time"

	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic"
	kubeinformers "k8s.io/client-go/informers"
	kubeclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	clusterclientset "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	clusterinformers "sigs.k8s.io/cluster-api/pkg/client/informers_generated/externalversions"

	operatorv1 "github.com/openshift/api/operator/v1"
	installertypes "github.com/openshift/installer/pkg/types"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/resource/resourcemerge"
	"github.com/openshift/library-go/pkg/operator/resource/resourceread"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	masterdnsv1 "github.com/openshift/master-dns-operator/pkg/apis/masterdns/v1alpha1"
	"github.com/openshift/master-dns-operator/pkg/operator/assets"
)

const (
	clusterAPINamespace        = "openshift-cluster-api"
	operatorNamespace          = "openshift-master-dns-operator"
	workloadFailingCondition   = "WorkloadFailing"
	masterLabelSelector        = "sigs.k8s.io/cluster-api-machine-role=master"
	endpointsResourceName      = "masters"
	endpointsResourceNamespace = "openshift-master-dns"
)

const (
	// ClusterConfigNamespace is the namespace containing the cluster config
	ClusterConfigNamespace = "kube-system"
	// ClusterConfigName is the name of the cluster config configmap
	ClusterConfigName = "cluster-config-v1"
	// InstallConfigKey is the key in the cluster config configmap containing yaml installConfig data
	InstallConfigKey = "install-config"
)

var (
	configName           = types.NamespacedName{Name: "instance", Namespace: ""}
	endpointsName        = types.NamespacedName{Name: endpointsResourceName, Namespace: endpointsResourceNamespace}
	requeueInterval      = 10 * time.Second
	masterMachinePattern = regexp.MustCompile(`(.+)-master-([0-9]+)`)
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
	mdnsReconciler.clusterAPIClient, err = clusterclientset.NewForConfig(mgr.GetConfig())
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

	clusterAPIInformerFactory := clusterinformers.NewSharedInformerFactoryWithOptions(mdnsReconciler.clusterAPIClient, 10*time.Minute, clusterinformers.WithNamespace(clusterAPINamespace))
	machineInformer := clusterAPIInformerFactory.Cluster().V1alpha1().Machines().Informer()
	mgr.Add(&informerRunnable{informer: machineInformer})

	configInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(mdnsReconciler.kubeClient, 10*time.Minute, kubeinformers.WithNamespace(ClusterConfigNamespace))
	configInformer := configInformerFactory.Core().V1().ConfigMaps().Informer()
	mgr.Add(&informerRunnable{informer: configInformer})

	// Watch for changes to machines in cluster API namespace
	err = c.Watch(&source.Informer{Informer: machineInformer}, &handler.EnqueueRequestsFromMapFunc{ToRequests: configHandlerFunc})
	if err != nil {
		return err
	}

	// Watch for Deployments
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: configHandlerFunc})
	if err != nil {
		return err
	}

	// Watch for Service Accounts
	err = c.Watch(&source.Kind{Type: &corev1.ServiceAccount{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: configHandlerFunc})
	if err != nil {
		return err
	}

	// Watch for Install Config
	err = c.Watch(&source.Informer{Informer: configInformer}, &handler.EnqueueRequestsFromMapFunc{ToRequests: installConfigHandlerFunc})
	if err != nil {
		return err
	}

	// Watch for DNS endpoints
	err = c.Watch(&source.Kind{Type: &masterdnsv1.DNSEndpoint{}}, &handler.EnqueueRequestsFromMapFunc{ToRequests: configHandlerFunc})
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
	client           client.Client
	scheme           *runtime.Scheme
	kubeClient       kubeclient.Interface
	clusterAPIClient clusterclientset.Interface
	eventRecorder    events.Recorder
	imagePullSpec    string
}

// Reconcile reads that state of the cluster for a MasterDNSOperatorConfig object and makes changes based on the state read
// and what is in the MasterDNSOperatorConfig.Spec
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
// +kubebuilder:rbac:groups=masterdns.operator.openshift.io,resources=masterdnsoperatorconfig;masterdnsoperatorconfig/status,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles;clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts;configmaps,verbs=get;list;watch;create;update;patch;delete
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

	installConfig, err := getInstallConfig(r.kubeClient)
	if err != nil {
		return reconcile.Result{}, err
	}

	if !isConfigSupported(installConfig) {
		log.Debug("Installed provider not supported. Nothing to do.")
		return reconcile.Result{}, nil
	}

	providerArgs := getProviderArgs(installConfig)

	// Ensure that External DNS is deployed
	err = r.syncExternalDNS(operatorConfig, providerArgs)
	if err != nil {
		log.WithError(err).Error("An error occurred synchronizing the external DNS deployment")
		return reconcile.Result{}, err
	}

	// Ensure a dnsendpoint resource for masters exists with entries for every master machine
	err = r.syncDNSEndpointResource(installConfig)
	if err != nil {
		log.WithError(err).Error("An error occurred synchronizing the masters endpoint instance")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileMasterDNS) syncExternalDNS(config *masterdnsv1.MasterDNSOperatorConfig, providerArgs []string) error {

	errors := []error{}
	originalConfig := config.DeepCopy()

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
	actualDeployment, _, err := r.syncExternalDNSDeployment(config, providerArgs, forceDeployment)
	if err != nil {
		log.WithError(err).WithField("file", "config/deployment.yaml").Error("Apply error")
		errors = append(errors, fmt.Errorf("%q: %v", "config/deployment.yaml", err))
	}

	if actualDeployment.Status.ReadyReplicas > 0 {
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:               operatorv1.OperatorStatusTypeAvailable,
			Status:             operatorv1.ConditionTrue,
			LastTransitionTime: metav1.Now(),
		})
	} else {
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:               operatorv1.OperatorStatusTypeAvailable,
			Status:             operatorv1.ConditionFalse,
			Reason:             "NoPodsAvailable",
			Message:            "no deployment pods available on any node.",
			LastTransitionTime: metav1.Now(),
		})
	}

	/*
		if actualDeployment.ObjectMeta.Generation == config.Status.ObservedGeneration {
			v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
				Type:               operatorv1.OperatorStatusTypeProgressing,
				Status:             operatorv1.ConditionFalse,
				LastTransitionTime: metav1.Now(),
			})
		} else {
			v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
				Type:               operatorv1.OperatorStatusTypeProgressing,
				Status:             operatorv1.ConditionTrue,
				Reason:             "DesiredStateNotYetAchieved",
				Message:            fmt.Sprintf("Generation: expected: %v, actual: %v", config.Status.ObservedGeneration, actualDeployment.ObjectMeta.Generation),
				LastTransitionTime: metav1.Now(),
			})
		}
	*/

	config.Status.ObservedGeneration = config.ObjectMeta.Generation
	config.Status.ReadyReplicas = actualDeployment.Status.ReadyReplicas
	resourcemerge.SetDeploymentGeneration(&config.Status.Generations, actualDeployment)
	if len(errors) > 0 {
		message := ""
		for _, err := range errors {
			message = message + err.Error() + "\n"
		}
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:               workloadFailingCondition,
			Status:             operatorv1.ConditionTrue,
			Message:            message,
			Reason:             "SyncError",
			LastTransitionTime: metav1.Now(),
		})
	} else {
		v1helpers.SetOperatorCondition(&config.Status.Conditions, operatorv1.OperatorCondition{
			Type:               workloadFailingCondition,
			Status:             operatorv1.ConditionFalse,
			LastTransitionTime: metav1.Now(),
		})
	}
	if !equality.Semantic.DeepEqual(config.Status, originalConfig.Status) {
		if err := r.client.Status().Update(context.TODO(), config); err != nil {
			log.WithError(err).Error("Status update error")
			return err
		}
	}

	if len(errors) > 0 {
		log.WithField("Errors", errors).Error("errors occurred, requeuing")
		return utilerrors.NewAggregate(errors)
	}
	return nil
}

func (r *ReconcileMasterDNS) syncDNSEndpointResource(installConfig *installertypes.InstallConfig) error {
	masterMachines, err := r.clusterAPIClient.ClusterV1alpha1().Machines(clusterAPINamespace).List(metav1.ListOptions{LabelSelector: masterLabelSelector})
	if err != nil {
		log.WithError(err).Infof("Error fetching master machines")
		return err
	}
	endpoints := []*masterdnsv1.Endpoint{}
	for _, machine := range masterMachines.Items {
		endpoint := &masterdnsv1.Endpoint{}
		if !masterMachinePattern.MatchString(machine.Name) {
			log.WithField("machine", machine.Name).Warning("Machine name does not match expected pattern")
			continue
		}
		address := getInternalIP(&machine)
		if len(address) == 0 {
			log.WithField("machine", machine.Name).Warning("Skipping machine because it doesn't currently have an internal IP")
			continue
		}
		nameParts := masterMachinePattern.FindStringSubmatch(machine.Name)
		endpoint.DNSName = fmt.Sprintf("%s-etcd-%s.%s", nameParts[1], nameParts[2], installConfig.BaseDomain)
		endpoint.RecordType = "A"
		endpoint.RecordTTL = 60
		endpoint.Targets = masterdnsv1.Targets{address}

		endpoints = append(endpoints, endpoint)
	}
	sort.Sort(endpointsByName(endpoints))

	dnsEndpoint := &masterdnsv1.DNSEndpoint{}
	err = r.client.Get(context.TODO(), endpointsName, dnsEndpoint)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	// If instance has not been created, create it
	if err != nil {
		dnsEndpoint.Name = endpointsResourceName
		dnsEndpoint.Namespace = endpointsResourceNamespace
		dnsEndpoint.Spec.Endpoints = endpoints
		return r.client.Create(context.TODO(), dnsEndpoint)
	}

	if !reflect.DeepEqual(dnsEndpoint.Spec.Endpoints, endpoints) {
		dnsEndpoint.Spec.Endpoints = endpoints
		return r.client.Update(context.TODO(), dnsEndpoint)
	}

	log.Info("DNSEndpoint instance is up to date")

	return nil
}

func (r *ReconcileMasterDNS) syncExternalDNSDeployment(config *masterdnsv1.MasterDNSOperatorConfig, providerArgs []string, forceDeployment bool) (*appsv1.Deployment, bool, error) {
	deployment := resourceread.ReadDeploymentV1OrDie(assets.MustAsset("config/deployment.yaml"))
	container := &deployment.Spec.Template.Spec.Containers[0]
	container.Image = r.imagePullSpec
	container.ImagePullPolicy = corev1.PullAlways
	for _, arg := range providerArgs {
		container.Args = append(container.Args, arg)
	}
	if len(config.Spec.LogLevel) > 0 {
		container.Args = append(container.Args, fmt.Sprintf("--log-level=%s", config.Spec.LogLevel))
	}
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

var installConfigHandlerFunc configHandler = func(o handler.MapObject) []reconcile.Request {
	if o.Meta.GetName() == ClusterConfigName && o.Meta.GetNamespace() == ClusterConfigNamespace {
		return []reconcile.Request{
			{
				NamespacedName: configName,
			},
		}
	}
	return []reconcile.Request{}
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

func getInstallConfig(client kubeclient.Interface) (*installertypes.InstallConfig, error) {
	cm, err := client.CoreV1().ConfigMaps(ClusterConfigNamespace).Get(ClusterConfigName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed getting clusterconfig %s/%s: %v", ClusterConfigNamespace, ClusterConfigName, err)
	}

	return getInstallConfigFromClusterConfig(cm)
}

// getInstallConfigFromClusterConfig builds an install config from the cluster config.
func getInstallConfigFromClusterConfig(clusterConfig *corev1.ConfigMap) (*installertypes.InstallConfig, error) {
	installConfigData, ok := clusterConfig.Data[InstallConfigKey]
	if !ok {
		return nil, fmt.Errorf("missing %q in configmap", InstallConfigKey)
	}
	installConfig := &installertypes.InstallConfig{}
	if err := yaml.Unmarshal([]byte(installConfigData), installConfig); err != nil {
		return nil, fmt.Errorf("invalid InstallConfig: %v yaml: %s", err, installConfigData)
	}
	return installConfig, nil
}

func isConfigSupported(config *installertypes.InstallConfig) bool {
	// Currently only AWS is supported
	return config.Platform.AWS != nil // || config.Platform.OtherPlatform != nil ..., etc
}

func getProviderArgs(config *installertypes.InstallConfig) []string {
	switch {
	case config.Platform.AWS != nil:
		return []string{
			"--provider=aws",
			fmt.Sprintf("--domain-filter=%s", config.BaseDomain),
			"--aws-zone-type=private",
			fmt.Sprintf("--aws-zone-tags=tectonicClusterID=%s", config.ClusterID),
		}
	}
	return nil
}

func getInternalIP(machine *clusterv1.Machine) string {
	for _, address := range machine.Status.Addresses {
		if address.Type == corev1.NodeInternalIP {
			return address.Address
		}
	}
	return ""
}

type endpointsByName []*masterdnsv1.Endpoint

func (s endpointsByName) Len() int           { return len(s) }
func (s endpointsByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s endpointsByName) Less(i, j int) bool { return s[i].DNSName < s[j].DNSName }
