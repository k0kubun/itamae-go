package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Service struct {
	Base
	Name     string
	Provider string
}

func (s *Service) Apply() {
	logger.Debug("file[" + s.Name + "] will not change")
}

func (s *Service) DryRun() {
	s.notifyApply()
}
