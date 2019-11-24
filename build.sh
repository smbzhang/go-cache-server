#!/bin/sh

export GOPATH=`pwd`:$GOPATH
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

cd src/cache-benchmark && go build && cd $curdir
cd src/server01 && go build && cd $curdir
cd src/server02 && go build && cd $curdir
cd src/server03 && go build && cd $curdir

cd src/client && go build && cd $curdir





