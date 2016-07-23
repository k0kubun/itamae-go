package utils

import (
	"log"

	"github.com/mitchellh/go-mruby"
)

func AssertType(attr string, v *mruby.MrbValue, vtype mruby.ValueType) {
	if v.Type() != vtype {
		log.Fatalf("%s attribute should be %s. (Itamae::Resource::InvalidTypeError)", attr, typeString(vtype))
	}
}

func typeString(vtype mruby.ValueType) string {
	switch vtype {
	case mruby.TypeString:
		return "String"
	case mruby.TypeProc:
		return "Proc"
	default:
		return "UnexpectedType"
	}
}
