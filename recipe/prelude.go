package recipe

import (
	"log"

	"github.com/k0kubun/itamae-go/recipe/dsl"
	"github.com/mitchellh/go-mruby"
)

// Define Ruby modules to execute itamae recipes
func prelude(mrb *mruby.Mrb) {
	defineNameSpaces(mrb)
	defineDSL(mrb)
	defineResourceEvalContext(mrb)
}

func defineDSL(mrb *mruby.Mrb) {
	kernel := mrb.KernelModule()
	kernel.DefineMethod("define", dsl.Define, mruby.ArgsReq(1))
	kernel.DefineMethod("directory", dsl.Directory, mruby.ArgsReq(1))
	kernel.DefineMethod("execute", dsl.Execute, mruby.ArgsReq(1))
	kernel.DefineMethod("file", dsl.File, mruby.ArgsReq(1))
	kernel.DefineMethod("git", dsl.Git, mruby.ArgsReq(1))
	kernel.DefineMethod("include_recipe", dsl.IncludeRecipe, mruby.ArgsReq(1))
	kernel.DefineMethod("link", dsl.Link, mruby.ArgsReq(1))
	kernel.DefineMethod("package", dsl.Package, mruby.ArgsReq(1))
	kernel.DefineMethod("remote_file", dsl.RemoteFile, mruby.ArgsReq(1))
	kernel.DefineMethod("run_command", dsl.RunCommand, mruby.ArgsReq(1))
	kernel.DefineMethod("service", dsl.Service, mruby.ArgsReq(1))
	kernel.DefineMethod("template", dsl.Template, mruby.ArgsReq(1))
}

func defineNameSpaces(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
	  module Itamae
		  class Resource
			  class Base
				end
			end
		end
	`)
	assertError(err)
}

func defineResourceEvalContext(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
		class Itamae::Resource::Base::EvalContext
		  attr_reader :attributes

			# FIXME: This should be resource-specific.
			ATTRIBUTES = %i[
			  action
				user
				cwd
				only_if
				not_if
				version
				options
			]

		  def initialize(*)
			  @attributes = {}
			end

			# FIXME: Should have type assertion
			def method_missing(method, *args)
			  if ATTRIBUTES.include?(method)
				  if args.length == 1
						@attributes[method] = args.first
					else
						raise "unexpected #{method} attribute specification: #{args.inspect}"
					end
				else
				  super
				end
			end
		end
	`)
	assertError(err)
}

func assertError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
