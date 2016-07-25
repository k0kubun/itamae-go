package resource

import (
	"os"

	"github.com/k0kubun/itamae-go/logger"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
)

type Execute struct {
	Base
	Command string
}

func (e *Execute) Apply() {
	logger.Color(logger.Green, func() {
		logger.Info(e.Resource + " executed will change from 'false' to 'true'")
		if !utils.Execute(e.Command) {
			logger.Error(e.Resource + " Failed.")
			os.Exit(1)
		}
	})
}

func (e *Execute) DryRun() {
	logger.Color(logger.Green, func() {
		logger.Info(e.Resource + " executed will change from 'false' to 'true'")
	})
}
