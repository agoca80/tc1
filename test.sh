#!/bin/bash

go run main.go client &
testing=true go run main.go server 

input  () { sort input  | uniq ; }
output () { sort numbers.log; }

echo "Sent     $(wc -l input) numbers"
echo "Sent     $(input |  wc -l) unique numbers"
echo "Received $(output | wc -l) unique numbers"
if cmp -s <(input) <(output); then
    echo OK
else
    echo FAIL
fi

