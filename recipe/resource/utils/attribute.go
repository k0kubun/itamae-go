package utils

import (
	"log"

	"github.com/mitchellh/go-mruby"
)

type AttributeParser struct {
	mrb         *mruby.Mrb
	evalContext *mruby.MrbValue
}

func NewAttributeParser(mrb *mruby.Mrb) *AttributeParser {
	evalContext, err := mrb.LoadString("Itamae::Resource::Base::EvalContext.new")
	assertError(err)
	return &AttributeParser{
		mrb:         mrb,
		evalContext: evalContext,
	}
}

func (a *AttributeParser) Parse(block *mruby.MrbValue) {
	_, err := a.evalContext.CallBlock("instance_eval", block)
	assertError(err)
}

func (a *AttributeParser) SetDefaultString(name string, str string) {
	value := a.mrb.StringValue(str)

	_, err := a.attributes().Call("[]=", a.toSym(name), value)
	assertError(err)
}

func (a *AttributeParser) GetString(name string) string {
	value, err := a.attributes().Call("[]", a.toSym(name))
	assertError(err)

	if value.Type() != mruby.TypeNil {
		AssertType(name, value, mruby.TypeString)
	}
	return value.String()
}

func (a *AttributeParser) GetArray(name string) []string {
	value, err := a.attributes().Call("[]", a.toSym(name))
	assertError(err)
	switch value.Type() {
	case mruby.TypeArray:
		arr := value.Array()
		ret := make([]string, 0, arr.Len())
		for i := 0; i < arr.Len(); i++ {
			val, err := arr.Get(i)
			assertError(err)
			ret = append(ret, val.String())
		}
		return ret
	default:
		return []string{value.String()}
	}
}

func (a *AttributeParser) toSym(str string) *mruby.MrbValue {
	sym, err := a.mrb.LoadString(":" + str)
	assertError(err)
	return sym
}

func (a *AttributeParser) attributes() *mruby.MrbValue {
	attributes, err := a.evalContext.Call("attributes")
	assertError(err)
	return attributes
}

func assertError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
