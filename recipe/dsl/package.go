package dsl

import (
	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func Package(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := m.GetArgs()
	utils.AssertType("name", args[0], mruby.TypeString)

	resource.Register(&resource.Package{
		Name: args[0].String(),
	})
	return mruby.Nil, nil
}
