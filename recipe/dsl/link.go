package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Link(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("link", args[0], mruby.TypeString)

	link := resource.Link{
		Link: args[0].String(),
	}

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "create")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	link.Action = parser.GetArray("action")
	link.User = parser.GetString("user")
	link.Cwd = parser.GetString("cwd")
	link.OnlyIf = parser.GetString("only_if")
	link.NotIf = parser.GetString("not_if")
	link.To = parser.GetString("to")
	link.Force = parser.GetString("force")

	resource.Register(&link)
	return mruby.Nil, nil
}
