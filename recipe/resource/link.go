package resource

import (
	"github.com/k0kubun/itamae-go/specinfra"
)

type Link struct {
	Base
	Link  string
	To    string
	Force string
}

func (l *Link) Apply() {
	for _, action := range l.Action {
		if action == "create" {
			l.actionCreate()
		}
	}
}

func (l *Link) actionCreate() {
	l.notifyApply()
	l.run(specinfra.LinkFileTo(l.Link, l.To, l.Force == "true"))
}

func (l *Link) DryRun() {
	for _, action := range l.Action {
		if action == "create" {
			l.notifyApply()
		}
	}
}
