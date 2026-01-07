#!/usr/bin/env bash

# watch code changes, trigger re-build, and kill process 
while true; do
    go build -o build/pug cmd/tui/main.go && pkill -f 'build/pug'
    inotifywait -e modify,attrib $(find . -name '*.go') || exit
done
