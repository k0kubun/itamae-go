package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func File(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("path", args[0], mruby.TypeString)

	file := resource.File{
		Path: args[0].String(),
	}

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "create")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	file.Action = parser.GetArray("action")
	file.User = parser.GetString("user")
	file.Cwd = parser.GetString("cwd")
	file.OnlyIf = parser.GetString("only_if")
	file.NotIf = parser.GetString("not_if")
	file.Content = parser.GetString("content")
	file.Mode = parser.GetString("mode")
	file.Owner = parser.GetString("owner")
	file.Group = parser.GetString("group")

	resource.Register(&file)
	return mruby.Nil, nil
}
