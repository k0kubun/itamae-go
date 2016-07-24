package recipe

import (
	"io/ioutil"

	"github.com/k0kubun/itamae-go/logger"
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

func (c *RecipeContext) LoadJson(file string) {
	if len(file) == 0 {
		return
	}

	jsonBuf, err := ioutil.ReadFile(file)
	assertError(err)

	loader, err := c.mrb.LoadString("Itamae::NodeLoader")
	assertError(err)

	_, err = loader.Call("load_json", c.mrb.StringValue(string(jsonBuf)))
	assertError(err)
}

func (c *RecipeContext) LoadRecipe(file string) {
	buf, err := ioutil.ReadFile(file)
	assertError(err)

	logger.Info("Recipe: " + file)
	dsl.PushRecipe(file)
	logger.WithIndent(func() {
		_, err = c.mrb.LoadString("ITAMAE_CONTEXT.instance_exec {" + string(buf) + "}")
		assertError(err)
	})
	dsl.PopRecipe()
}

func (c *RecipeContext) Resources() []resource.Resource {
	return resource.Registered()
}
