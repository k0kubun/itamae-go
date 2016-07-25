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
	logger.Debug("link[" + l.Link + "] will not change")
}

func (l *Link) DryRun() {
	logger.Color(logger.Green, func() {
		logger.Info(l.Resource + " executed will change from 'false' to 'true'")
	})
}
