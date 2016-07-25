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
	for _, action := range g.Action {
		if action == "sync" {
			g.actionSync()
		}
	}
}

func (g *Git) actionSync() {
	logger.Debug("git[" + g.Destination + "] will not change")
}

func (g *Git) DryRun() {
	for _, action := range g.Action {
		if action == "sync" {
			g.notifyApply()
		}
	}
}
