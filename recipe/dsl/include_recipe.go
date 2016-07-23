package dsl

import (
	"fmt"

	"github.com/k0kubun/itamae-go/recipe/resource/utils"
	"github.com/mitchellh/go-mruby"
)

func IncludeRecipe(m *mruby.Mrb, self *mruby.MrbValue) (mruby.Value, mruby.Value) {
	args := m.GetArgs()
	utils.AssertType("path", args[0], mruby.TypeString)

	recipePath := args[0].String()
	fmt.Println("include_recipe: " + recipePath)
	return mruby.Nil, nil
}

// FIXME: This should be in recipe.RecipeContext
var recipeStack = make([]string, 0)

func PushRecipe(file string) {
	recipeStack = append(recipeStack, file)
}

func PopRecipe() {
	recipeStack = recipeStack[0 : len(recipeStack)-1]
}
