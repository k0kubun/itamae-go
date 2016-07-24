package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Service(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("service resource (stubbed)")
	return mruby.Nil, nil
}
