package dsl

import (
	"fmt"

	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/mitchellh/go-mruby"
)

func Package(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("package resource (stubbed)")
	resource.Register(&resource.Package{
		Name: "vim",
	})
	return mruby.Nil, nil
}
