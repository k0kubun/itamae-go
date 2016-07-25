package resource

import (
	"fmt"

	"github.com/k0kubun/itamae-go/logger"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
)

type Base struct {
	Resource string
	Action   []string
	User     string
	Cwd      string
	OnlyIf   string
	NotIf    string
}

func (b *Base) ShouldSkip() bool {
	if b.OnlyIf != "" && !b.execute(b.OnlyIf) {
		logger.Debug(b.Resource + " Execution skipped because of not_if attribute")
		return true
	}
	if b.NotIf != "" && b.execute(b.NotIf) {
		logger.Debug(b.Resource + " Execution skipped because of only_if attribute")
		return true
	}
	return false
}

func (b *Base) execute(str string) bool {
	// XXX: Correctly escape shell....
	if b.Cwd != "" {
		str = fmt.Sprintf("cd %s && %s", b.Cwd, str)
	}
	if b.User != "" {
		str = fmt.Sprintf("sudo -H -u %s -- /bin/sh -c '%s'", b.User, str)
	}
	return utils.Execute(str)
}

func (b *Base) notifyApply() {
	logger.Color(logger.Green, func() {
		logger.Info(b.Resource + " executed will change from 'false' to 'true'")
	})
}
