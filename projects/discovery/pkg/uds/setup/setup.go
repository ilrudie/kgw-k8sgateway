package setup

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/utils/setuputils"
	"github.com/solo-io/solo-kit/projects/discovery/pkg/uds/syncer"
)

func Main(settingsDir string) error {
	return setuputils.Main("fds", syncer.NewSetupSyncer(memory.NewInMemoryResourceCache(), kube.NewKubeCache()), settingsDir)

}
