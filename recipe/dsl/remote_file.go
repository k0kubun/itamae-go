package dsl

import (
	"fmt"

	"github.com/mitchellh/go-mruby"
)

func RemoteFile(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	fmt.Println("remote_file resource (stubbed)")
	return mruby.Nil, nil
}
