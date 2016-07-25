package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Git(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("destination", args[0], mruby.TypeString)

	git := resource.Git{
		Destination: args[0].String(),
	}
	git.Destination = "git[" + git.Destination + "]"

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "sync")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	git.Action = parser.GetArray("action")
	git.User = parser.GetString("user")
	git.Cwd = parser.GetString("cwd")
	git.OnlyIf = parser.GetString("only_if")
	git.NotIf = parser.GetString("not_if")
	git.Repository = parser.GetString("repository")
	git.Revision = parser.GetString("revision")
	git.Recursive = parser.GetString("recursive")

	resource.Register(&git)
	return mruby.Nil, nil
}
