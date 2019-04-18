#!/bin/bash
# linux to windows
GOOS=windows GOARCH=amd64 go build -o hello.exe hello.go

# if your package need cgo
# in debian
sudo apt-get install gcc-multilib
sudo apt-get install gcc-mingw-w64

GOOS=windows GOARCH=386 \
CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
go build -o hello.exe hello.go

# mac to linux
GOOS=linux GOARCH=amd64 go build
