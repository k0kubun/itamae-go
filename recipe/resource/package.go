package resource

import (
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/k0kubun/itamae-go/specinfra"
)

type Package struct {
	Base
	Name    string
	Version string
	Options string
}

func (p *Package) Apply() {
	for _, action := range p.Action {
		if action == "install" {
			p.actionInstall()
		} else if action == "remove" {
			p.actionRemove()
		}
	}
}

func (p *Package) actionInstall() {
	// XXX: Check version...
	if !p.execute(specinfra.CheckPackageIsInstalled(p.Name, p.Version)) {
		p.notifyApply()
		utils.Run(specinfra.Install(p.Name, p.Version, p.Options))
	}
}

func (p *Package) actionRemove() {
}

func (p *Package) DryRun() {
	for _, action := range p.Action {
		if action == "install" {
			p.notifyApply()
		} else if action == "remove" {
			p.notifyApply()
		}
	}
}
