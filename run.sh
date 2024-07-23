#!/bin/bash
export GOPATH=/d/me/go
export GOMODCACHE=$GOPATH/pkg/mod
export GOCACHE=/d/me/go/cache

go build -o bin/go-ecom cmd/main.go
if [ $? -eq 0 ]; then
    ./bin/go-ecom
else
    echo "Build failed"
fi

# If Makefile way is falied you need to follow this approach
# To use this script:

# Save it as run.sh in your project's root directory.
# !!! Make it executable (you only need to do this once) !!!
# +x run.sh

# Run the script:
# ./run.sh