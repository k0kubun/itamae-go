# itamae-go

itamae-go is Go implementation of [itamae](https://github.com/itamae-kitchen/itamae) embedding mruby, which is a lightweight configuration management tool inspired by Chef.
You can write a configuration recipe in Ruby and apply it without Ruby.

## Status

Experimental. Internal architecture is poor and many features are omitted or not tested.

## Synopsis

Like original [itamae](https://github.com/itamae-kitchen/itamae), you can manage configuration by Ruby DSL. But itamae-go does not require Ruby to run.

```rb
# cat recipe.rb
include_recipe 'included.rb'

directory '/tmp/etc'

3.times do |i|
  file "/tmp/etc/#{i}"
end

template '/tmp/etc/config.yml' do
  source 'config.yml.erb'
end
```

```rb
# cat included.rb
define :vim, options: '--with-lua --with-luajit' do
  ver  = params[:name]
  opts = params[:options]
  package 'vim' do
    version ver
    options opts
  end
end

vim '7.4.1910-1'

service 'mysqld' do
  action [:start, :enable]
end
```

```bash
# ./itamae-go local -j node.json recipe.rb
 INFO : Starting itamae...
 INFO : Recipe: recipe.rb
 INFO :   Recipe: included.rb
 INFO : package[vim] executed will change from 'false' to 'true'
 INFO : service[mysqld] executed will change from 'false' to 'true'
 INFO : service[mysqld] executed will change from 'false' to 'true'
 INFO : directory[/tmp/etc] executed will change from 'false' to 'true'
 INFO : file[/tmp/etc/0] executed will change from 'false' to 'true'
 INFO : file[/tmp/etc/1] executed will change from 'false' to 'true'
 INFO : file[/tmp/etc/2] executed will change from 'false' to 'true'
 INFO : file[/tmp/etc/config.yml] executed will change from 'false' to 'true'
```

## How to write recipes

See [itamae's reference](https://github.com/itamae-kitchen/itamae/wiki).

### Supported features

You can use only the features listed below.

- Resources
  - [x] execute resource
  - [x] package resource
  - [x] directory resource
  - [x] git resource
  - [x] file resource
  - [x] template resource
  - [x] link resource
  - [x] service resource
  - [ ] user resource
  - [ ] group resource
  - [ ] remote\_file resource
  - [ ] remote\_directory resource
  - [ ] gem\_package resource
  - [ ] http\_request resource
  - [ ] local\_ruby\_block resource
- [x] Definitions
- [x] Including Recipes
- [x] Node Attributes
- [ ] run\_command
- [ ] Host Inventory
- [ ] Plugins

### Supported platforms

- Arch Linux
- Debian
- OSX
- Ubuntu

## Contribution

1. Fork ([https://github.com/k0kubun/itamae-go/fork](https://github.com/k0kubun/itamae-go/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create a new Pull Request

## Notes

Thanks to the original implementation.

https://github.com/itamae-kitchen/itamae

## Author

[Takashi Kokubun](https://github.com/k0kubun)
