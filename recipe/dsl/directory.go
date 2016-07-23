package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Directory(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("directory resource (stubbed)")
	return mruby.Nil, nil
}
