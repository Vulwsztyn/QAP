#!/bin/bash
for i in {0..9}
do
   rm ./consts.go
   printf "package main\nconst instanceIndex = $i" >> ./consts.go
   go build
   ./QAP
done
