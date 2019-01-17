package setup

import (
	"github.com/solo-io/gloo/pkg/utils/setuputils"
	"github.com/solo-io/gloo/projects/discovery/pkg/uds/syncer"
	gloosyncer "github.com/solo-io/gloo/projects/gloo/pkg/syncer"
)

func Main() error {
	return setuputils.Main(setuputils.SetupOpts{
		LoggingPrefix: "uds",
		SetupFunc:     gloosyncer.NewSetupFuncWithRun(syncer.RunUDS),
		ExitOnError:   true,
	})
}
