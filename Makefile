MRUBY_COMMIT ?= 1.2.0

all: itamae-go

itamae-go: libmruby.a
	go get -d github.com/mitchellh/go-mruby
	go build

clean:
	rm -rf vendor
	rm -f libmruby.a

libmruby.a: vendor/mruby/build/host/lib/libmruby.a
	cp vendor/mruby/build/host/lib/libmruby.a .

vendor/mruby/build/host/lib/libmruby.a: vendor/mruby
	cd vendor/mruby && MRUBY_CONFIG=../../build_config.rb ${MAKE}

vendor/mruby:
	mkdir -p vendor
	git clone https://github.com/mruby/mruby.git vendor/mruby
	cd vendor/mruby && git reset --hard && git clean -fdx
	cd vendor/mruby && git checkout ${MRUBY_COMMIT}

.PHONY: all clean libmruby.a test
