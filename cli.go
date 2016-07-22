package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/k0kubun/itamae-go/command"
	"github.com/mitchellh/cli"
)

func Run(args []string) int {
	// Meta-option for executables.
	// It defines output color and its stdout/stderr stream.
	meta := &command.Meta{
		Ui: &cli.ColoredUi{
			InfoColor:  cli.UiColorBlue,
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		},
	}

	return RunCustom(args, Commands(meta))
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {
	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:       args,
		Commands:   commands,
		Version:    Version,
		HelpFunc:   BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}

func BasicHelpFunc(app string) cli.HelpFunc {
	return func(commands map[string]cli.CommandFactory) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf(
			"usage: %s [--version] [--help] <command> [<args>]\n\n",
			app))
		buf.WriteString("Available commands are:\n")

		// Get the list of keys so we can sort them, and also get the maximum
		// key length so they can be aligned properly.
		keys := make([]string, 0, len(commands))
		maxKeyLen := 0
		for key, _ := range commands {
			if len(key) > maxKeyLen {
				maxKeyLen = len(key)
			}

			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			commandFunc, ok := commands[key]
			if !ok {
				// This should never happen since we JUST built the list of
				// keys.
				panic("command not found: " + key)
			}

			command, err := commandFunc()
			if err != nil {
				log.Printf("[ERR] cli: Command '%s' failed to load: %s",
					key, err)
				continue
			}

			key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
			buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
		}

		return buf.String()
	}
}
