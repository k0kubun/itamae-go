package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Service(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("name", args[0], mruby.TypeString)

	srv := resource.Service{
		Name: args[0].String(),
	}

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "create")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	srv.Action = parser.GetArray("action")
	srv.User = parser.GetString("user")
	srv.Cwd = parser.GetString("cwd")
	srv.OnlyIf = parser.GetString("only_if")
	srv.NotIf = parser.GetString("not_if")
	srv.Provider = parser.GetString("provider")

	resource.Register(&srv)
	return mruby.Nil, nil
}
