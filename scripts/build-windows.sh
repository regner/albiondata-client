#!/usr/bin/env bash

export CGO_CPPFLAGS="-I thirdpaty/WpdPack/Include/"
export CGO_LDFLAGS="-L thirdpaty/WpdPack/Lib/x64/"
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1
export CXX=x86_64-w64-mingw32-g++
export CC=x86_64-w64-mingw32-gcc
go build -o albionmarket-client.exe cmd/albionmarket-client/albionmarket-client.go
