package resource

import (
	"github.com/k0kubun/itamae-go/logger"
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
	if utils.Execute(specinfra.CheckPackageIsInstalled(p.Name, p.Version)) {
		logger.Debug("package[" + p.Name + "] will not change")
	} else {
		logger.Color(logger.Green, func() {
			logger.Info("package[" + p.Name + "] will change")
		})
		utils.Run(specinfra.Install(p.Name, p.Version, p.Options))
	}
}

func (p *Package) DryRun() {
	p.notifyApply()
}
