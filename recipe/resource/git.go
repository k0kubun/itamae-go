package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Git struct {
	Base
	Destination string
	Repository  string
	Revision    string
	Recursive   string
}

func (g *Git) Apply() {
	logger.Debug("git[" + g.Destination + "] will not change")
}

func (g *Git) DryRun() {
	g.notifyApply()
}
