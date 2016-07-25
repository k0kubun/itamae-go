package recipe

import (
	"log"

	"github.com/k0kubun/itamae-go/recipe/dsl"
	"github.com/mitchellh/go-mruby"
)

// Define Ruby modules to execute itamae recipes
func prelude(mrb *mruby.Mrb) {
	defineDSL(mrb)
	defineDefine(mrb)
	defineNode(mrb)
	defineResourceEvalContext(mrb)
	defineTemplateResource(mrb)
}

func defineDSL(mrb *mruby.Mrb) {
	mItamae := mrb.DefineModule("Itamae")
	cRecipe := mrb.DefineClassUnder("Recipe", nil, mItamae)
	cEvalContext := mrb.DefineClassUnder("EvalContext", nil, cRecipe)

	cEvalContext.DefineMethod("directory", dsl.Directory, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("execute", dsl.Execute, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("file", dsl.File, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("git", dsl.Git, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("include_recipe", dsl.IncludeRecipe, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("link", dsl.Link, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("package", dsl.Package, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("remote_file", dsl.RemoteFile, mruby.ArgsReq(1))
	cEvalContext.DefineMethod("service", dsl.Service, mruby.ArgsReq(1))

	_, err := mrb.LoadString("ITAMAE_CONTEXT = Itamae::Recipe::EvalContext.new")
	assertError(err)
}

func defineDefine(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
		module Itamae::Recipe::Define
		  class DefineContext
				attr_reader :params

				def initialize(defaults)
					@params = defaults
					@attributes = defaults.keys
				end

				def method_missing(method, *args)
				  if @attributes.include?(method)
						if args.length == 1
							@params[method] = args.first
						else
							raise "unexpected #{method} attribute specification: #{args.inspect}"
						end
					else
						super
					end
				end
			end

		  attr_reader :params

			def define(name, options = {}, &define_block)
			  scope = self
			  self.class.define_method(name.to_sym) do |name, &arg_block|
				  context = DefineContext.new(options)
					context.instance_exec(&arg_block) if arg_block
					@params = context.params.merge(name: name)
					scope.instance_exec(&define_block)
				end
			end
		end
		Itamae::Recipe::EvalContext.prepend(Itamae::Recipe::Define)
	`)
	assertError(err)
}

func defineNode(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
		module Itamae::NodeLoader
		  def self.load_json(json)
			  @node = JSON.load(json)
			end

			def self.loaded_node
			  @node
			end
		end

		Kernel.class_eval do
		  def node
		    Itamae::NodeLoader.loaded_node || {}
		  end
		end
	`)
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
				source
				content
				mode
				owner
				group
				repository
				revision
				recursive
				to
				force
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
					# Show error before super because go-mruby's LoadString crashes on raise...
					puts "Undefined method '#{method}'"
					super
				end
			end
		end
	`)
	assertError(err)
}

func defineTemplateResource(mrb *mruby.Mrb) {
	_, err := mrb.LoadString(`
		module Itamae::Recipe::TemplateResource
			def template(name, &block)
			  context = Itamae::Resource::Base::EvalContext.new
				context.instance_exec(&block)
				src_file = context.attributes.delete(:source)
				raise "source must be specified for template" if src_file.nil?

				# We should read path relative from recipe and pass variables as binding
				raise "template source not found (currently it's not relative path from recipe)" unless File.exist?(src_file)
				erb_result = ERB.new(File.read(src_file)).result
				attrs      = context.attributes

			  file name do
				  content erb_result
				  attrs.each do |key, value|
						send(key, value)
					end
				end
			end
		end
		Itamae::Recipe::EvalContext.prepend(Itamae::Recipe::TemplateResource)
	`)
	assertError(err)
}

func assertError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
