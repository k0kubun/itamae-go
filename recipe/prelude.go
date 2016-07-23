package recipe

import (
	"log"

	"github.com/k0kubun/itamae-go/recipe/dsl"
	"github.com/mitchellh/go-mruby"
)

// Define Ruby modules to execute itamae recipes
func prelude(mrb *mruby.Mrb) {
	defineDSL(mrb)
	defineResourceEvalContext(mrb)
}

func defineDSL(mrb *mruby.Mrb) {
	mItamae := mrb.DefineModule("Itamae")
	cRecipe := mrb.DefineClassUnder("Recipe", nil, mItamae)
	cEvalContext := mrb.DefineClassUnder("EvalContext", nil, cRecipe)

	cEvalContext.DefineMethod("define", dsl.Define, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("directory", dsl.Directory, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("execute", dsl.Execute, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("file", dsl.File, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("git", dsl.Git, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("include_recipe", dsl.IncludeRecipe, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("link", dsl.Link, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("package", dsl.Package, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("remote_file", dsl.RemoteFile, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("run_command", dsl.RunCommand, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("service", dsl.Service, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("template", dsl.Template, mruby.ArgsReq(1))

	_, err := mrb.LoadString("ITAMAE_CONTEXT = Itamae::Recipe::EvalContext.new")
	assertError(err)
}

func defineResourceEvalContext(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
		class Itamae::Resource; end
		class Itamae::Resource::Base; end

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
