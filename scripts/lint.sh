#!/usr/bin/env sh
SCRIPT_DIR=$(dirname "$0")
cd $SCRIPT_DIR/..

go vet ./...

if command -v docker &> /dev/null
then
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v
fi
