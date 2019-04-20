#!/bin/bash

DUMP=/tmp/dump
INPUT=/tmp/input
OUTPUT=/tmp/output

go build
./tc1

input  () { sort $INPUT  | uniq ; }
output () { sort $OUTPUT; }

echo "Sent     $(wc -l $INPUT) numbers"
echo "Sent     $(input |  wc -l) unique numbers"
echo "Received $(output | wc -l) unique numbers"
if cmp -s <(input) <(output); then
    echo OK
else
    echo FAIL
fi

rm $INPUT $OUTPUT $DUMP
