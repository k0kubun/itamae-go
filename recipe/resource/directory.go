package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Directory struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Path  string
	Mode  string
	Owner string
	Group string
}

func (d *Directory) Apply() {
	logger.Debug("directory[" + d.Path + "] will not change")
}

func (d *Directory) DryRun() {
	// TODO: do some checks...
	logger.Debug("directory[" + d.Path + "] will be applied")
}
