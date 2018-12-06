package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"

	clusterapis "sigs.k8s.io/cluster-api/pkg/apis"

	"github.com/openshift/master-dns-operator/pkg/apis"
	"github.com/openshift/master-dns-operator/pkg/operator"
)

func main() {
	rand.
		Seed(time.Now().UTC().UnixNano())

	pflag.CommandLine.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	logs.InitLogs()
	defer logs.FlushLogs()

	command := NewOperatorManagerCommand()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

const defaultLogLevel = "info"

// NewOperatorManagerCommand returns a command for the Manager
func NewOperatorManagerCommand() *cobra.Command {
	var logLevel string
	cmd := &cobra.Command{
		Use:   "manager",
		Short: "OpenShift Master DNS operator",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set log level
			level, err := log.ParseLevel(logLevel)
			if err != nil {
				log.WithError(err).Fatal("Cannot parse log level")
			}
			log.SetLevel(level)
			log.Debug("debug logging enabled")

			return runOperator()
		},
	}
	cmd.PersistentFlags().StringVar(&logLevel, "log-level", defaultLogLevel, "Log level (debug,info,warn,error,fatal)")
	goflag.CommandLine.Parse([]string{})
	if stderrFlag := pflag.CommandLine.Lookup("logtostderr"); stderrFlag != nil {
		stderrFlag.Value.Set("true")
	}

	return cmd
}

func runOperator() error {
	printVersion()

	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.WithError(err).Error("failed to get watch namespace")
		return err
	}

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.WithError(err).Error("error getting client config")
		return err
	}

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{Namespace: namespace})
	if err != nil {
		log.WithError(err).Error("error creating controller manager")
		return err
	}

	log.Info("Registering Components")

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.WithError(err).Error("error registering master-dns-operator types")
		return err
	}

	if err := clusterapis.AddToScheme(mgr.GetScheme()); err != nil {
		log.WithError(err).Error("error registering cluster-api types")
		return err
	}

	// Setup all Controllers
	if err := operator.AddToManager(mgr); err != nil {
		log.WithError(err).Error("error setting up controllers")
		return err
	}

	log.Info("Starting the Operator")

	// Start the Cmd
	return mgr.Start(signals.SetupSignalHandler())
}

func printVersion() {
	log.Infof("Go Version: %s", runtime.Version())
	log.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	log.Infof("operator-sdk Version: %v", sdkVersion.Version)
}
