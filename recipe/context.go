package recipe

import (
	"log"

	"github.com/k0kubun/itamae-go/recipe/resource"
	"github.com/mitchellh/go-mruby"
)

type RecipeContext struct {
	mrb *mruby.Mrb
}

func NewContext() *RecipeContext {
	context := &RecipeContext{
		mrb: mruby.NewMrb(),
	}
	prelude(context.mrb)
	return context
}

func (c *RecipeContext) Close() {
	c.mrb.Close()
}

func (c *RecipeContext) LoadRecipe(src string) {
	_, err := c.mrb.LoadString(src)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *RecipeContext) Resources() []resource.Resource {
	return resource.Registered()
}
