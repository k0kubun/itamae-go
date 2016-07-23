package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func File(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("file resource (stubbed)")
	return mruby.Nil, nil
}
