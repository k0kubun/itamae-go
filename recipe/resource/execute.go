package resource

import (
	"os"

	"github.com/k0kubun/itamae-go/logger"
)

type Execute struct {
	Base
	Command string
}

func (e *Execute) Apply() {
	e.notifyApply()
	if !e.execute(e.Command) {
		logger.Error(e.Resource + " Failed.")
		os.Exit(1)
	}
}

func (e *Execute) DryRun() {
	e.notifyApply()
}
