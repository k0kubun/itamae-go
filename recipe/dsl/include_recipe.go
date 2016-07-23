package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func IncludeRecipe(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("include_recipe (stubbed)")
	return mruby.Nil, nil
}
