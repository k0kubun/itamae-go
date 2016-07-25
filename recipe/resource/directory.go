package resource

import (
	"github.com/k0kubun/itamae-go/specinfra"
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
	// XXX: Consider mode, owner, group...
	d.notifyApply()
	d.run(specinfra.CreateAsDirectory(d.Path))
}

func (d *Directory) actionDelete() {
	if !d.execute(specinfra.CheckFileIsDirectory(d.Path)) {
		return
	}

	d.notifyApply()
	d.run(specinfra.RemoveFile(d.Path))
}

func (d *Directory) DryRun() {
	for _, action := range d.Action {
		if action == "create" {
			d.notifyApply()
		} else if action == "delete" {
			if d.execute(specinfra.CheckFileIsDirectory(d.Path)) {
				d.notifyApply()
			}
		}
	}
}
