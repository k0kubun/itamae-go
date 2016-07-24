package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Execute(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("execute resource (stubbed)")
	return mruby.Nil, nil
}
