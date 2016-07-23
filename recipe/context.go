package recipe

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/itamae-go/recipe/dsl"
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

func (c *RecipeContext) LoadRecipe(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	dsl.PushRecipe(file)
	_, err = c.mrb.LoadString("ITAMAE_CONTEXT.instance_exec {" + string(buf) + "}")
	if err != nil {
		log.Fatal(err)
	}
	dsl.PopRecipe()
}

func (c *RecipeContext) Resources() []resource.Resource {
	return resource.Registered()
}
