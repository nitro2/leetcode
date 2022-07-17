#!/bin/bash
no_column=$((`head -n 1 file.txt | tr -dc ' ' | wc -c` +1))
# echo $no_column
out=""
for i in `seq 1 $no_column`
do
    out+=`awk -v j=$i '{print $j}' file.txt`"\n"
done

echo -e $out | sed -E -e '/^$/d'