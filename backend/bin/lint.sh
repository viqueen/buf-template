#!/usr/bin/env bash

docker run --rm \
  --volume "${PWD}":/go/src/backend \
  --volume "${PWD}/../api":/go/src/api \
  --workdir /go/src/backend \
  golangci/golangci-lint:v2.1.2 \
  golangci-lint run "$@"
