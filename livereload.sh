#!/usr/bin/env bash

BUILDRELOADER=build/livereloader
# watch code changes, trigger re-build, and kill process 
while true; do
    go build -o $BUILDRELOADER cmd/tui/main.go && pkill -f $BUILDRELOADER
    inotifywait -e modify,attrib $(find . -name '*.go') || exit
done
