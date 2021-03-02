BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)
LDFLAGS=-ldflags "-w -s -X 'main.Version=${VERSION}' -X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommit=$(GIT_COMMIT)'"

.PHONY: audit
audit:
	go list -m all | nancy sleuth

.PHONY: build
build: generate-prod
	go build $(LDFLAGS) -tags 'production' -o $(BINPATH)/dp-render

.PHONY: debug
debug: generate-debug
	go build $(LDFLAGS) -race -tags 'debug' -o build/dp-render

.PHONY: generate-debug
generate-debug:
	# build the dev version
	cd assets; go run github.com/jteeuwen/go-bindata/go-bindata -debug -o data.go -pkg assets locales/...
	{ echo "// +build debug\n"; cat assets/data.go; } > assets/debug.go.new
	mv assets/debug.go.new assets/data.go

.PHONY: generate-prod
generate-prod:
	# build the production version
	cd assets; go run github.com/jteeuwen/go-bindata/go-bindata -o data.go -pkg assets locales/...
	{ echo "// +build production\n"; cat assets/data.go; } > assets/data.go.new
	mv assets/data.go.new assets/data.go

.PHONY: test
test: generate-prod
	go test -race -cover -tags 'production' ./...
