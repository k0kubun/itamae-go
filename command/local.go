package command

import (
	"strings"
)

type LocalCommand struct {
	Meta
}

func (c *LocalCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *LocalCommand) Synopsis() string {
	return ""
}

func (c *LocalCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
