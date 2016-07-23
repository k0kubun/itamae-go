package command

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/k0kubun/itamae-go/itamae"
	"github.com/k0kubun/itamae-go/recipe"
)

type LocalCommand struct {
	Meta

	dryRun   bool
	nodeJson string
	logLevel string
	recipes  []string
}

func (c *LocalCommand) Run(args []string) int {
	flags := flag.NewFlagSet("itamae-go", flag.ContinueOnError)

	flags.BoolVar(&c.dryRun, "n", false, "Dry run")
	flags.BoolVar(&c.dryRun, "dry-run", false, "Dry run")
	flags.StringVar(&c.nodeJson, "j", "", "Node JSON")
	flags.StringVar(&c.nodeJson, "node-json", "", "Node JSON")
	flags.StringVar(&c.logLevel, "l", "", "Log level")
	flags.StringVar(&c.logLevel, "log-level", "", "Log level")

	if err := flags.Parse(args); err != nil {
		return 1
	}
	for 0 < flags.NArg() {
		c.recipes = append(c.recipes, flags.Arg(0))
		flags.Parse(flags.Args()[1:])
	}

	context := recipe.NewContext()
	for _, file := range c.recipes {
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		context.LoadRecipe(string(buf))
	}

	if c.dryRun {
		itamae.DryRun("dry-run (stubbed)")
	} else {
		itamae.Apply("apply (stubbed)")
	}
	return 0
}

func (c *LocalCommand) Synopsis() string {
	return "Run itamae locally"
}

func (c *LocalCommand) Help() string {
	helpText := `
Usage:
  itamae-go local RECIPE [RECIPE...]

Options:
  -j, [--node-json=NODE_JSON]
  -n, [--dry-run]
  -l, [--log-level=LOG_LEVEL]

Run Itamae locally
`
	return strings.TrimSpace(helpText)
}
