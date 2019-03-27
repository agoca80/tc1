#!/bin/bash

for profile in cpu mem block; do
    go run main.go client &
    go test -${profile}profile $profile.prof -run=Server

    echo "top10 -cum" | go tool pprof $profile.prof
    echo "web"        | go tool pprof $profile.prof
done
