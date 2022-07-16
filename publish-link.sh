#!/bin/bash
add_link() {
    echo - Problem \[$1\]\("https://leetcode.com/problems/"$1\)  >> README.md
}

# Then  ${f} | rev | cut -c2- | rev    -> remove last `/` 
for f in `ls -d */`; do t=$(echo ${f} | rev | cut -c2- | rev) && add_link "${t}" ; done