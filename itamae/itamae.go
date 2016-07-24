package itamae

import (
	"github.com/k0kubun/itamae-go/logger"
	. "github.com/k0kubun/itamae-go/recipe/resource"
)

func Apply(resources []Resource) {
	for _, resource := range resources {
		resource.Apply()
	}
}

func DryRun(resources []Resource) {
	logger.Info("itamae dry-run (stubbed)")
}
