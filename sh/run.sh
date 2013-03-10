#!/bin/bash
export GOROOT=$HOME/go; export PATH=$PATH:$GOROOT/bin
ls  | grep -v go | grep -v sh | xargs rm -f && go run eg.go
