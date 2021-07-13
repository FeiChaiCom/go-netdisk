#!/bin/bash

VERSION=$1
PUBLISH=$2

docker build -t go-netdisk .
docker tag go-netdisk feichaicom/go-netdisk:${VERSION}

# shellcheck disable=SC2236
if [[ ! -z "$PUBLISH" ]]; then
  docker push feichaicom/go-netdisk:${VERSION}
fi
