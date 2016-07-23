package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Define(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("define (stubbed)")
	return mruby.Nil, nil
}
