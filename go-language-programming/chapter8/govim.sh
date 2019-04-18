#!/bin/bash
mkdir -p $HOME/.vim/ftdetect
mkdir -p $HOME/.vim/syntax
mkdir -p $HOME/.vim/autoload/go
ln -s $GOROOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect/ ln -s $GOROOT/misc/vim/syntax/go.vim $HOME/.vim/syntax
ln -s $GOROOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go

# fuck gfw
#go get -u github.com/mdempsky/gocode

cd $GOPATH/src/github/ && mkdir mdempsky && cd mdempsky && git clone https://github.com/mdempsky/gocode.git
cd $GOPATH/src/github.com/mdempsky/gocode/ && cd vim && sh update.sh
