MRUBY_COMMIT ?= 1.2.0
GO_MRUBY_DIR := "../../mitchellh/go-mruby"
.PHONY: all itamae-go clean

all: itamae-go

itamae-go: libmruby.a ${GO_MRUBY_DIR}/libmruby.a
	go get github.com/mitchellh/cli
	go build

clean:
	if [ -e vendor/mruby ]; then \
		cd vendor/mruby && ${MAKE} clean ; \
	fi
	if [ -e ${GO_MRUBY_DIR}/vendor/mruby ]; then \
		cd ${GO_MRUBY_DIR}/vendor/mruby && ${MAKE} clean ; \
	fi
	rm -f libmruby.a
	rm -f ${GO_MRUBY_DIR}/libmruby.a

libmruby.a: vendor/mruby/build/host/lib/libmruby.a
	cp vendor/mruby/build/host/lib/libmruby.a .

vendor/mruby/build/host/lib/libmruby.a: vendor/mruby
	cd vendor/mruby && MRUBY_CONFIG=../../build_config.rb ${MAKE}

vendor/mruby:
	mkdir -p vendor
	git clone https://github.com/mruby/mruby.git vendor/mruby
	cd vendor/mruby && git reset --hard && git clean -fdx
	cd vendor/mruby && git checkout ${MRUBY_COMMIT}

${GO_MRUBY_DIR}/libmruby.a: ${GO_MRUBY_DIR}/vendor/mruby/build/host/lib/libmruby.a
	cp ${GO_MRUBY_DIR}/vendor/mruby/build/host/lib/libmruby.a ${GO_MRUBY_DIR}

${GO_MRUBY_DIR}/vendor/mruby/build/host/lib/libmruby.a:
	if [ ! -e ${GO_MRUBY_DIR}/libmruby.a ]; then \
	  cd ${GO_MRUBY_DIR} && MRUBY_COMMIT=${MRUBY_COMMIT} MRUBY_CONFIG=../../../../k0kubun/itamae-go/build_config.rb ${MAKE} ; \
	fi
