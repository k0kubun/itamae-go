package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Execute struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Command string
}

func (e *Execute) Apply() {
	logger.Debug("execute[" + e.Command + "] will not change")
}

func (e *Execute) DryRun() {
	// TODO: do some checks...
	logger.Debug("execute[" + e.Command + "] will be applied")
}
