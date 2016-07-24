package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Link(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("link resource (stubbed)")
	return mruby.Nil, nil
}
