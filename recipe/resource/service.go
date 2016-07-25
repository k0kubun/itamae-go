package resource

import (
	"github.com/k0kubun/itamae-go/specinfra"
)

type Service struct {
	Base
	Name     string
	Provider string
}

func (s *Service) Apply() {
	for _, action := range s.Action {
		if action == "start" {
			s.actionStart()
		} else if action == "stop" {
			s.actionStop()
		} else if action == "enable" {
			s.actionEnable()
		} else if action == "disable" {
			s.actionDisable()
		}
	}
}

func (s *Service) actionStart() {
	s.notifyApply()
	s.run(specinfra.StartService(s.Name))
}

func (s *Service) actionStop() {
	s.notifyApply()
	s.run(specinfra.StopService(s.Name))
}

func (s *Service) actionEnable() {
	s.notifyApply()
	s.run(specinfra.EnableService(s.Name))
}

func (s *Service) actionDisable() {
	s.notifyApply()
	s.run(specinfra.DisableService(s.Name))
}

func (s *Service) DryRun() {
	for _, action := range s.Action {
		if action == "start" {
			s.notifyApply()
		} else if action == "stop" {
			s.notifyApply()
		} else if action == "enable" {
			s.notifyApply()
		} else if action == "disable" {
			s.notifyApply()
		}
	}
}
