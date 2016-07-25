package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Package(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("name", args[0], mruby.TypeString)

	pkg := resource.Package{
		Name: args[0].String(),
	}
	pkg.Resource = "package[" + pkg.Name + "]"

	parser := utils.NewAttributeParser(mrb)
	parser.SetDefaultString("action", "install")
	if len(args) > 1 {
		utils.AssertType("block", args[1], mruby.TypeProc)
		parser.Parse(args[1])
	}

	pkg.Action = parser.GetArray("action")
	pkg.User = parser.GetString("user")
	pkg.Cwd = parser.GetString("cwd")
	pkg.OnlyIf = parser.GetString("only_if")
	pkg.NotIf = parser.GetString("not_if")
	pkg.Version = parser.GetString("version")
	pkg.Options = parser.GetString("options")

	resource.Register(&pkg)
	return mruby.Nil, nil
}
