package resource

import (
	"github.com/k0kubun/itamae-go/logger"
)

type Link struct {
	Action []string
	User   string
	Cwd    string
	OnlyIf string
	NotIf  string

	Link  string
	To    string
	Force string
}

func (l *Link) Apply() {
	logger.Debug("link[" + l.Link + "] will not change")
}
