package recipe

import (
	"log"

	"github.com/mitchellh/go-mruby"
)

type EvalContext struct {
	mrb *mruby.Mrb
}

func NewContext() *EvalContext {
	return &EvalContext{
		mrb: mruby.NewMrb(),
	}
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
