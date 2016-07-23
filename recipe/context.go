package recipe

import (
	"log"

	"github.com/mitchellh/go-mruby"
)

type EvalContext struct {
}

func NewContext() *EvalContext {
	return &EvalContext{}
}

func (c *EvalContext) LoadRecipe(src string) {
	mrb := mruby.NewMrb()
	defer mrb.Close()

	_, err := mrb.LoadString(src)
	if err != nil {
		log.Fatal(err)
	}
}
