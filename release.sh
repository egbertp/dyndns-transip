#!/bin/sh

BUILD_VERSION=`git describe --exact-match --tags 2> /dev/null`
[ -z "$BUILD_VERSION" ] && BUILD_VERSION="v0.0.0"
COMMIT_HASH=`git rev-parse --short=8 HEAD 2>/dev/null`
BUILD_TIME=`date +%FT%T%z`

gox -ldflags "-X main.ApplicationVersion=$BUILD_VERSION -X main.CommitHash=$COMMIT_HASH -X main.BuildTime=$BUILD_TIME" -os="linux darwin windows openbsd" -arch="amd64" -output="dist/dyndns-transip_{{.OS}}_{{.Arch}}" ./cmd/