package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func Git(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("git resource (stubbed)")
	return mruby.Nil, nil
}
