#!/bin/bash

i=0
while read -r line; do
    # echo $i $line
    i=$((i+1))
    if [ $i -eq 10 ]
    then
        echo $line
    fi
done < $1