#!/usr/bin/env sh
SCRIPT_DIR=$(dirname "$0")
cd $SCRIPT_DIR/..

go build -v -o bin/lnkshrt cmd/lnkshrt/main.go
