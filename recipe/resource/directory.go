package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Directory struct {
	Base
	Path  string
	Mode  string
	Owner string
	Group string
}

func (d *Directory) Apply() {
	logger.Debug("directory[" + d.Path + "] exist will change from 'false' to 'true'")
}

func (d *Directory) DryRun() {
	logger.Color(logger.Green, func() {
		logger.Info(d.Resource + " executed will change from 'false' to 'true'")
	})
}
