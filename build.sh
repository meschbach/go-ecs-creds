#!/bin/bash -xe

workspace=$PWD
vendor=$PWD/vendor
if [ -z "$GOPATH"] ; then
	export GOPATH=$vendor:$workspace
else
	export GOPATH=$GOPATH:$vendor:$workspace
fi

go get github.com/meschbach/go-ecs-creds
go build github.com/meschbach/go-ecs-creds

