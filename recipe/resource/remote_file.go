package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type RemoteFile struct {
	Base
	Path    string
	Content string
	Mode    string
	Owner   string
	Group   string
	Source  string
}

func (r *RemoteFile) Apply() {
	for _, action := range r.Action {
		if action == "create" {
			r.actionCreate()
		} else if action == "delete" {
			r.actionDelete()
		} else if action == "edit" {
			r.actionEdit()
		}
	}
}

func (r *RemoteFile) actionCreate() {
	logger.Debug("remote_file[" + r.Path + "] will not change")
}

func (r *RemoteFile) actionDelete() {
}

func (r *RemoteFile) actionEdit() {
}

func (r *RemoteFile) DryRun() {
	for _, action := range r.Action {
		if action == "create" {
			r.notifyApply()
		} else if action == "delete" {
			r.notifyApply()
		} else if action == "edit" {
			r.notifyApply()
		}
	}
}
