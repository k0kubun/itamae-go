package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func RunCommand(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("run_command (stubbed)")
	return mruby.Nil, nil
}
