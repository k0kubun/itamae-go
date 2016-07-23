package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Execute(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("execute resource (stubbed)")
	return mruby.Nil, nil
}
