#!/bin/bash
export GOPATH=/Users/xieyuandong/Workspace/go/src/github.com/xieydd/Go_Zh/go-language-programming/chapter2/sorter && go build sorter
./sorter -a bubblesort -i ../input.dat -o ../output.dat
