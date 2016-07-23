package recipe

import (
	"log"

	"github.com/k0kubun/itamae-go/recipe/dsl"
	"github.com/mitchellh/go-mruby"
)

type EvalContext struct {
	mrb *mruby.Mrb
}

func NewContext() *EvalContext {
	context := &EvalContext{
		mrb: mruby.NewMrb(),
	}
	context.prelude()
	return context
}

func (c *EvalContext) Close() {
	c.mrb.Close()
}

func (c *EvalContext) LoadRecipe(src string) {
	_, err := c.mrb.LoadString(src)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *EvalContext) prelude() {
	kernel := c.mrb.KernelModule()
	kernel.DefineMethod("directory", dsl.Directory, mruby.ArgsReq(1))
	kernel.DefineMethod("execute", dsl.Execute, mruby.ArgsReq(1))
	kernel.DefineMethod("file", dsl.File, mruby.ArgsReq(1))
	kernel.DefineMethod("git", dsl.Git, mruby.ArgsReq(1))
	kernel.DefineMethod("link", dsl.Link, mruby.ArgsReq(1))
	kernel.DefineMethod("package", dsl.Package, mruby.ArgsReq(1))
	kernel.DefineMethod("remote_file", dsl.RemoteFile, mruby.ArgsReq(1))
	kernel.DefineMethod("service", dsl.Service, mruby.ArgsReq(1))
	kernel.DefineMethod("template", dsl.Template, mruby.ArgsReq(1))
}
