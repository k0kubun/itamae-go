package dsl

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/k0kubun/itamae-go/logger"
	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func IncludeRecipe(mrb *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := mrb.GetArgs()
	utils.AssertType("path", args[0], mruby.TypeString)

	recipe := args[0].String()
	path := recipePath(recipe)
	PushRecipe(recipe)
	loadRecipe(mrb, path)
	PopRecipe()
	return mruby.Nil, nil
}

func loadRecipe(mrb *mruby.Mrb, path string) {
	logger.Info("Recipe: " + path)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = mrb.LoadString("ITAMAE_CONTEXT.instance_exec {" + string(buf) + "}")
	if err != nil {
		log.Fatal(err)
	}
}

func recipePath(target string) string {
	dir := "."
	for _, recipe := range recipeStack {
		dir = filepath.Join(dir, filepath.Dir(recipe))
	}
	return filepath.Join(dir, target)
}

// FIXME: This should be in recipe.RecipeContext
var recipeStack = make([]string, 0)

func PushRecipe(file string) {
	recipeStack = append(recipeStack, file)
}

func PopRecipe() {
	recipeStack = recipeStack[0 : len(recipeStack)-1]
}
