package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Execute(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("command", args[0], mruby.TypeString)

	exec := resource.Execute{
		Command: args[0].String(),
	}

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "run")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	exec.Action = parser.GetArray("action")
	exec.User = parser.GetString("user")
	exec.Cwd = parser.GetString("cwd")
	exec.OnlyIf = parser.GetString("only_if")
	exec.NotIf = parser.GetString("not_if")

	resource.Register(&exec)
	return mruby.Nil, nil
}
