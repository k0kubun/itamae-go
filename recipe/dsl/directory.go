package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Directory(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("path", args[0], mruby.TypeString)

	dir := resource.Directory{
		Path: args[0].String(),
	}

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "create")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	dir.Action = parser.GetArray("action")
	dir.User = parser.GetString("user")
	dir.Cwd = parser.GetString("cwd")
	dir.OnlyIf = parser.GetString("only_if")
	dir.NotIf = parser.GetString("not_if")
	dir.Mode = parser.GetString("mode")
	dir.Owner = parser.GetString("owner")
	dir.Group = parser.GetString("group")

	resource.Register(&dir)
	return mruby.Nil, nil
}
