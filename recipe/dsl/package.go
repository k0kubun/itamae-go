package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Package(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("package resource (stubbed)")
	return mruby.Nil, nil
}
