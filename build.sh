#!/bin/bash -xe

workspace=$PWD
if [ -z "$GOPATH"] ; then
	export GOPATH=$workspace
else
	export GOPATH=$GOPATH:$workspace
fi

go build github.com/meschbach/go-ecs-creds
