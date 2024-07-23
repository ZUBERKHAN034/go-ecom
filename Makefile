export GOPATH=/d/me/go
export GOMODCACHE=$(GOPATH)/pkg/mod
export GOCACHE=$(GOPATH)/cache
export GOTMPDIR=$(GOPATH)/tmp

build:
	@go build -o bin/go-ecom cmd/main.go

run: build
	@./bin/go-ecom

test:
	@go test ./...