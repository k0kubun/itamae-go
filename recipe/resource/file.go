package resource

import (
	"github.com/k0kubun/itamae-go/logger"
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
	logger.Debug("file[" + f.Path + "] will not change")
}

func (f *File) actionDelete() {
	logger.Debug("file[" + f.Path + "] will not change")
}

func (f *File) actionEdit() {
	logger.Debug("file[" + f.Path + "] will not change")
}

func (f *File) DryRun() {
	for _, action := range f.Action {
		if action == "create" {
			f.notifyApply()
		} else if action == "delete" {
			f.notifyApply()
		} else if action == "edit" {
			f.notifyApply()
		}
	}
}
