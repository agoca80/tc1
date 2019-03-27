#!/bin/bash

for profile in cpu mem block;do
echo "top10 -cum" | go tool pprof ${profile}.prof
echo "web"        | go tool pprof ${profile}.prof
done
