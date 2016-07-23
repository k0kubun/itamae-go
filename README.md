# itamae-go

itamae-go is Go implementation of [itamae](https://github.com/itamae-kitchen/itamae) embedding mruby, which is a lightweight configuration management tool inspired by Chef.
You can write a configuration recipe in Ruby and apply it without Ruby!

## How to write recipes

See [itamae's reference](https://github.com/itamae-kitchen/itamae/wiki).

### Supported features

You can use only the features listed below.

- Resources
  - [ ] execute resource
  - [ ] package resource
  - [ ] directory resource
  - [ ] git resource
  - [ ] file resource
  - [ ] remote\_file resource
  - [ ] template resource
  - [ ] link resource
  - [ ] service resource
  - [ ] user resource
  - [ ] group resource
  - [ ] remote\_directory resource
  - [ ] gem\_package resource
  - [ ] http\_request resource
  - [ ] local\_ruby\_block resource
- [x] Definitions
- [x] Including Recipes
- [ ] Node Attributes
- [ ] run\_command
- [ ] Host Inventory
- [ ] Plugins

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
