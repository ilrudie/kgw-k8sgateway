package setuputils

import (
	"context"
	"flag"
	"sync"
	"time"

	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector"
	kube2 "github.com/solo-io/gloo/pkg/bootstrap/leaderelector/kube"
	"github.com/solo-io/gloo/pkg/bootstrap/leaderelector/singlereplica"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/external/kubernetes/namespace"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type SetupOpts struct {
	LoggerName string
	// logged as the version of Gloo currently executing
	Version     string
	SetupFunc   SetupFunc
	ExitOnError bool
	CustomCtx   context.Context

	// optional - if present, add these values in each JSON log line in the gloo pod.
	// By default, we already log the gloo version.
	LoggingPrefixVals []interface{}
	// optional - if present, report usage with the payload this discovers
	// should really only provide it in very intentional places- in the gloo pod, and in glooctl
	// otherwise, we'll provide redundant copies of the usage data

	ElectionConfig *leaderelector.ElectionConfig
}

var once sync.Once

// there is a setup event loop that will sync with a setup function
// the setup function is usually a setup_syncer. The setup_syner Setup() function
// is called by the event_loop to setup up the environment.  This should contain
// the Gloo Settings.
// Main is the main entrypoint for running Gloo Edge components
// It works by performing the following:
//	1. Initialize a SettingsClient backed either by Kubernetes or a File
// 	2. Run an event loop, watching events on the Settings resource, and executing the
//		opts.SetupFunc whenever settings change
// This allows Gloo components to automatically receive updates to Settings and reload their
// configuration, without needing to restart the container
func Main(opts SetupOpts) error {
	// prevent panic if multiple flag.Parse called concurrently
	once.Do(func() {
		flag.Parse()
	})

	ctx := opts.CustomCtx
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = contextutils.WithLogger(ctx, opts.LoggerName)
	loggingContext := append([]interface{}{"version", opts.Version}, opts.LoggingPrefixVals...)
	ctx = contextutils.WithLoggerValues(ctx, loggingContext...)

	settingsClient, err := fileOrKubeSettingsClient(ctx, setupNamespace, setupDir)
	if err != nil {
		return err
	}

	if err := settingsClient.Register(); err != nil {
		return err
	}

	identity, err := startLeaderElection(ctx, setupDir, opts.ElectionConfig)
	if err != nil {
		return err
	}

	// settings come from the ResourceClient in the settingsClient
	// the eventLoop will Watch the emitter's settingsClient to recieve settings from the ResourceClient
	emitter, err := snapshotEmitter(ctx, settingsClient)
	if err != nil {
		return err
	}
	settingsRef := &core.ResourceRef{Namespace: setupNamespace, Name: setupName}
	eventLoop := v1.NewSetupEventLoop(emitter, NewSetupSyncer(settingsRef, opts.SetupFunc, identity))
	errs, err := eventLoop.Run([]string{setupNamespace}, clients.WatchOpts{
		Ctx:         ctx,
		RefreshRate: time.Second,
	})
	if err != nil {
		return err
	}
	for err := range errs {
		if opts.ExitOnError {
			contextutils.LoggerFrom(ctx).Fatalf("error in setup: %v", err)
		}
		contextutils.LoggerFrom(ctx).Errorf("error in setup: %v", err)
	}
	return nil
}

func fileOrKubeSettingsClient(ctx context.Context, setupNamespace, settingsDir string) (v1.SettingsClient, error) {
	if settingsDir != "" {
		contextutils.LoggerFrom(ctx).Infow("using filesystem for settings", zap.String("directory", settingsDir))
		return v1.NewSettingsClient(ctx, &factory.FileResourceClientFactory{
			RootDir: settingsDir,
		})
	}

	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, err
	}
	return v1.NewSettingsClient(ctx, &factory.KubeResourceClientFactory{
		Crd:                v1.SettingsCrd,
		Cfg:                cfg,
		SharedCache:        kube.NewKubeCache(ctx),
		NamespaceWhitelist: []string{setupNamespace},
	})
}

func startLeaderElection(ctx context.Context, settingsDir string, electionConfig *leaderelector.ElectionConfig) (leaderelector.Identity, error) {
	if electionConfig == nil || settingsDir != "" {
		return singlereplica.NewElectionFactory().StartElection(ctx, electionConfig)
	}

	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return nil, err
	}
	return kube2.NewElectionFactory(cfg).StartElection(ctx, electionConfig)
}

func snapshotEmitter(ctx context.Context, settingsClient v1.SettingsClient) (v1.SetupEmitter, error) {
	var kube kubernetes.Interface
	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		// it is not neccessary to have the config because this will occur at set up
		contextutils.LoggerFrom(ctx).Warnf("setup of kube client config: %v", err)
	} else {
		kube, err = kubernetes.NewForConfig(cfg)
		if err != nil {
			return nil, err
		}
	}
	resourceNamespaceLister := namespace.NewKubeClientResourceNamespaceLister(kube)
	emitter := v1.NewSetupEmitter(settingsClient, resourceNamespaceLister)
	return emitter, nil
}
