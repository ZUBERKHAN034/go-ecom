APP_NAME = go-ecom
BIN_DIR = bin

ifeq ($(RENDER),true)
    GOOS = linux
    EXT =
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOOS = linux
        EXT =
    else ifeq ($(UNAME_S),Darwin)
        GOOS = darwin
        EXT =
    else
        GOOS = windows
        EXT = .exe
    endif
endif

build:
	GOOS=$(GOOS) GOARCH=amd64 go build -o $(BIN_DIR)/$(APP_NAME)$(EXT) cmd/main.go

run: build
	./$(BIN_DIR)/$(APP_NAME)$(EXT)

# Run tests
test:
	go test ./...

# Run the dev server (only for local use)
run-dev:
	air -c .air.toml

# Generate Swagger docs
run-swagger:
	cd cmd/ && swag init -g main.go
