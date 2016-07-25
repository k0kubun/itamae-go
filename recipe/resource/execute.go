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
	for _, action := range e.Action {
		if action == "run" {
			e.actionRun()
		}
	}
}

func (e *Execute) actionRun() {
	e.notifyApply()
	if !e.execute(e.Command) {
		logger.Error(e.Resource + " Failed.")
		os.Exit(1)
	}
}

func (e *Execute) DryRun() {
	for _, action := range e.Action {
		if action == "run" {
			e.notifyApply()
		}
	}
}
