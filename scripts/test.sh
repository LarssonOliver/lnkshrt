#!/usr/bin/env sh
SCRIPT_DIR=$(dirname "$0")
cd $SCRIPT_DIR/..

go test -race -v ./...
