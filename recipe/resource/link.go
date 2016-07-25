package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Link struct {
	Base
	Link  string
	To    string
	Force string
}

func (l *Link) Apply() {
	for _, action := range l.Action {
		if action == "run" {
			l.actionCreate()
		}
	}
}

func (l *Link) actionCreate() {
	logger.Debug("link[" + l.Link + "] will not change")
}

func (l *Link) DryRun() {
	for _, action := range l.Action {
		if action == "run" {
			l.notifyApply()
		}
	}
}
