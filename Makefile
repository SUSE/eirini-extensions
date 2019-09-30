ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: test-unit build image

.PHONY: build
build:
	bin/build

image:
	bin/build-image

export NAMESPACE ?= default
up:
	bin/up

gen-kube:
	bin/gen-kube

gen-fakes:
	bin/gen-fakes

verify-gen-kube:
	bin/verify-gen-kube

generate: gen-kube gen-fakes

vet:
	bin/vet

lint:
	bin/lint

test-unit:
	bin/test-unit

test-integration:
	bin/test-integration

test-e2e:
	bin/test-e2e

test: vet lint test-unit test-integration test-e2e

tools:
	bin/tools

check-scripts:
	bin/check-scripts

test-docker:
	docker run -v $(ROOT_DIR):/src/ --workdir /src/ --rm -ti golang make tools test-unit