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
