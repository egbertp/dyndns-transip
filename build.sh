#!/bin/sh

if [[ -e dyndns-transip ]]; then
    rm dyndns-transip
 fi

BUILD_VERSION=`git describe --exact-match --tags 2> /dev/null`
[ -z "$BUILD_VERSION" ] && BUILD_VERSION="v0.0.0"
COMMIT_HASH=`git rev-parse --short=8 HEAD 2>/dev/null`
BUILD_TIME=`date -u +%Y-%m-%dT%TZ`
BUILT_BY=`hostname`

go build -o dyndns-transip -ldflags "-X main.version=$BUILD_VERSION -X main.commit=$COMMIT_HASH -X main.date=$BUILD_TIME -X main.builtBy=$BUILT_BY" ./cmd/