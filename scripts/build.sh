#!/bin/bash

GO_VERSION=1.7.4

PROJECTS=(\
"example/cmd/example-srv" \
)

for PROJECT in "${PROJECTS[@]}"
do
	docker run --rm -it -v "$PWD":/go/src/github.com/coyle/go-starter -w /go/src/github.com/coyle/go-starter/${PROJECT} golang:${GO_VERSION} sh -c 'CGO_ENABLED=0 go build'
done
