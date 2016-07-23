package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Template(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("template resource (stubbed)")
	return mruby.Nil, nil
}
