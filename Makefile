export GOPATH=/d/me/go
export GOMODCACHE=$(GOPATH)/pkg/mod
export GOCACHE=$(GOPATH)/cache
export GOTMPDIR=$(GOPATH)/tmp

build:
	@go build -o bin/go-ecom.exe cmd/main.go

run: build
	@./bin/go-ecom.exe

test:
	@go test ./...

run-dev:
	air -c .air.toml

run-swagger:
	cd pkg/ && swag init -g ../cmd/main.go	