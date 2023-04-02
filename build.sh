#!/usr/bin/env bash

setContext() {
  goPath="${GOPATH}";
  cd $goPath/blaid;
}

build() {
  go clean;
  go fmt;
  go build;
}

main() {
  setContext;
  build;
}

main;