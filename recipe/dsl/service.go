package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Service(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("service resource (stubbed)")
	return mruby.Nil, nil
}
