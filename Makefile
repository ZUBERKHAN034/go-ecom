build:
	@go build -o bin/go-ecom.exe cmd/main.go

run: build
	@./bin/go-ecom.exe

test:
	@go test ./...

run-dev:
	air -c .air.toml

run-swagger:
	cd cmd/ && swag init -g main.go	