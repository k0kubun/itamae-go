package resource

import (
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
	if b.OnlyIf != "" && !utils.Execute(b.OnlyIf) {
		logger.Debug(b.Resource + " Execution skipped because of not_if attribute")
		return true
	}
	if b.NotIf != "" && utils.Execute(b.NotIf) {
		logger.Debug(b.Resource + " Execution skipped because of only_if attribute")
		return true
	}
	return false
}
