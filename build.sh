#!/bin/bash

VERSION=$(cat ./VERSION)
#PACKAGE="github.com/ki38sato/ec2ls-cache"

go get -v github.com/Masterminds/glide
go install github.com/Masterminds/glide
glide up

apt-get update
apt-get install -y zip

GOOS=linux GOARCH=amd64 go build -v -ldflags "-X main.version=${VERSION}" -o build/ec2ls-cache
cd build
tar czf ec2ls-cache_linux_amd64.tar.gz ec2ls-cache
cd ..

GOOS=darwin GOARCH=amd64 go build -v -ldflags "-X main.version=${VERSION}" -o build/ec2ls-cache
cd build
zip ec2ls-cache_darwin_amd64.zip ec2ls-cache
cd ..
