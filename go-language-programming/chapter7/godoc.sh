#!/bin/bash

go doc -all ./

godoc -http=:76 -index_files=github.com/xieydd/Go_Zh/go-language-programming/chapter7 

# link: http://localhost:76/pkg/github.com/xieydd/Go_Zh/go-language-programming/chapter7/
