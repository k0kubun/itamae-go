package main

import (
	"github.com/mitchellh/cli"
	"github.com/Takashi Kokubun/itamae-go/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"local": func() (cli.Command, error) {
			return &command.LocalCommand{
				Meta: *meta,
			}, nil
		},
		
        "version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
                Revision: GitCommit,
                Name: Name,
			}, nil
		},
	}
}
