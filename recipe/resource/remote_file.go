package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type RemoteFile struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Path    string
	Content string
	Mode    string
	Owner   string
	Group   string
	Source  string
}

func (r *RemoteFile) Apply() {
	logger.Debug("remote_file[" + r.Path + "] will not change")
}

func (r *RemoteFile) DryRun() {
	// TODO: do some checks...
	logger.Debug("remote_file[" + r.Path + "] will be applied")
}
