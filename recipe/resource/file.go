package resource

import (
	"io/ioutil"

	"github.com/k0kubun/itamae-go/specinfra"
)

type File struct {
	Base
	Path    string
	Content string
	Mode    string
	Owner   string
	Group   string
}

func (f *File) Apply() {
	for _, action := range f.Action {
		if action == "create" {
			f.actionCreate()
		} else if action == "delete" {
			f.actionDelete()
		} else if action == "edit" {
			f.actionEdit()
		}
	}
}

func (f *File) actionCreate() {
	// XXX: Consider mode, owner, group
	f.notifyApply()
	ioutil.WriteFile(f.Path, []byte(f.Content), 0644)
}

func (f *File) actionDelete() {
	if !f.execute(specinfra.CheckFileIsFile(f.Path)) {
		return
	}

	f.notifyApply()
	f.run(specinfra.RemoveFile(f.Path))
}

func (f *File) actionEdit() {
	// XXX: Consider mode, owner, group
	f.notifyApply()
	ioutil.WriteFile(f.Path, []byte(f.Content), 0644)
}

func (f *File) DryRun() {
	for _, action := range f.Action {
		if action == "create" {
			f.notifyApply()
		} else if action == "delete" {
			if f.execute(specinfra.CheckFileIsFile(f.Path)) {
				f.notifyApply()
			}
		} else if action == "edit" {
			f.notifyApply()
		}
	}
}
