package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type File struct {
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
}

func (f *File) Apply() {
	logger.Debug("file[" + f.Path + "] will not change")
}

func (f *File) DryRun() {
	// TODO: do some checks...
	logger.Debug("file[" + f.Path + "] will be applied")
}
