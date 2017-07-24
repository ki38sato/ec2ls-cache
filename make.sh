#!/bin/bash

rm -rf vendor/*
##https://github.com/Masterminds/glide/issues/654

echo "execute build.sh using golang:1.8.3"
docker run --rm -v "$(pwd)":/go/src/github.com/ki38sato/ec2ls-cache -w /go/src/github.com/ki38sato/ec2ls-cache golang:1.8.3 bash build.sh
