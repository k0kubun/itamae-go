package resource

import (
	"os"

	"github.com/k0kubun/itamae-go/logger"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
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
	logger.Color(logger.Green, func() {
		logger.Info("execute[" + e.Command + "] will change from 'false' to 'true'")
		if !utils.Execute(e.Command) {
			logger.Error("execute[" + e.Command + "] Failed.")
			os.Exit(1)
		}
	})
}

func (e *Execute) DryRun() {
	// TODO: do some checks...
	logger.Debug("execute[" + e.Command + "] will be applied")
}
