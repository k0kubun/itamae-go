package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Service struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Name     string
	Provider string
}

func (s *Service) Apply() {
	logger.Debug("file[" + s.Name + "] will not change")
}

func (s *Service) DryRun() {
	// TODO: do some checks...
	logger.Debug("file[" + s.Name + "] will be applied")
}
