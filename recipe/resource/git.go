package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Git struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Destination string
	Repository  string
	Revision    string
	Recursive   string
}

func (g *Git) Apply() {
	logger.Debug("git[" + g.Destination + "] will not change")
}

func (g *Git) DryRun() {
	// TODO: do some checks...
	logger.Debug("git[" + g.Destination + "] will be applied")
}
