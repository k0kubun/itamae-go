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
	for _, action := range d.Action {
		if action == "create" {
			d.actionCreate()
		} else if action == "delete" {
			d.actionDelete()
		}
	}
}

func (d *Directory) actionCreate() {
	logger.Debug("directory[" + d.Path + "] exist will change from 'false' to 'true'")
}

func (d *Directory) actionDelete() {
	logger.Debug("directory[" + d.Path + "] exist will change from 'false' to 'true'")
}

func (d *Directory) DryRun() {
	for _, action := range d.Action {
		if action == "create" || action == "delete" {
			d.notifyApply()
		}
	}
}
