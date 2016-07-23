package recipe

import (
	"fmt"
)

type EvalContext struct {
}

func NewContext() *EvalContext {
	return &EvalContext{}
}

func (c *EvalContext) LoadRecipe(src string) {
	fmt.Println(src)
}
