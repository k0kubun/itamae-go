package resource

import (
	"fmt"

	"github.com/k0kubun/itamae-go/specinfra"
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
	if g.execute(specinfra.CheckFileIsDirectory(g.Destination)) {
		return
	}

	g.notifyApply()
	cmd := "git clone "
	if g.Recursive == "true" {
		cmd += "--recursive "
	}
	cmd = fmt.Sprintf("%s %s %s", cmd, g.Repository, g.Destination)
	g.run(cmd)
}

func (g *Git) DryRun() {
	for _, action := range g.Action {
		if action == "sync" {
			if !g.execute(specinfra.CheckFileIsDirectory(g.Destination)) {
				g.notifyApply()
			}
		}
	}
}
