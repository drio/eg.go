#!/bin/bash
#

WATCH_THIS="*.go egenotype/*.go"
filewatcher "$WATCH_THIS" 'go test;echo "\n"'
