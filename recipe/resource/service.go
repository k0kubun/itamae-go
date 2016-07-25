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
	for _, action := range s.Action {
		if action == "start" {
			s.actionStart()
		} else if action == "stop" {
			s.actionStop()
		} else if action == "restart" {
			s.actionRestart()
		} else if action == "reload" {
			s.actionReload()
		} else if action == "enable" {
			s.actionEnable()
		} else if action == "disable" {
			s.actionDisable()
		}
	}
}

func (s *Service) actionStart() {
	logger.Debug("file[" + s.Name + "] will not change")
}
func (s *Service) actionStop() {
}
func (s *Service) actionRestart() {
}
func (s *Service) actionReload() {
}
func (s *Service) actionEnable() {
}
func (s *Service) actionDisable() {
}

func (s *Service) DryRun() {
	for _, action := range s.Action {
		if action == "start" {
			s.notifyApply()
		} else if action == "stop" {
			s.notifyApply()
		} else if action == "restart" {
			s.notifyApply()
		} else if action == "reload" {
			s.notifyApply()
		} else if action == "enable" {
			s.notifyApply()
		} else if action == "disable" {
			s.notifyApply()
		}
	}
}
