package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestLocalCommand_implement(t *testing.T) {
	var _ cli.Command = &LocalCommand{}
}
