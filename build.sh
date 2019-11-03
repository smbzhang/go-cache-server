#!/bin/sh

export GOPATH=`pwd`
echo $GOPATH
curdir=`pwd`

if [[ ! -e "${curdir}/src/github.com" ]]; then
    ln -s /opt/go/workspace/src/github.com "${curdir}/src"
fi
if [[ ! -e "${curdir}/src/golang.org" ]]; then
    ln -s /opt/go/workspace/src/golang.org "${curdir}/src"
fi
if [[ ! -e "${curdir}/src/stathat.com" ]]; then
    ln -s /opt/go/workspace/src/stathat.com "${curdir}/src"
fi
if [[ ! -e "${curdir}/src/rocksdb" ]]; then
    ln -s "${curdir}/rocksdb" "${curdir}/src"
fi

cd src/server01 && go build


