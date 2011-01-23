#!/bin/bash

for cpp in *.cpp
do
    out="${cpp%.cpp}.out"
    if test ${cpp} -nt ${out}
    then
	printf "g++ -o %s %s\n" "${out}" "${cpp}"
	g++ -o ${out} ${cpp}
    fi
done
