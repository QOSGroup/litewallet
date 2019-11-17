#!/bin/bash

env GO111MODULE=off go get -v -d github.com/QOSGroup/litewallet

cd ${GOPATH}/src/github.com/QOSGroup/litewallet/ || echo "get litewallet err!" && exit 1

env GO111MODULE=on go mod tidy
env GO111MODULE=on go mod vendor

if [ ! -d "${GOPATH}/src/github.com/ethereum/go-ethereum" ];then
  env GO111MODULE=off go get -v -d github.com/ethereum/go-ethereum
fi

cp -r \
  "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
  "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"

env GO111MODULE=off gomobile bind -target android -o litewallet.aar github.com/QOSGroup/litewallet/mobile

rm -rf vendor