#!/usr/bin/env bash

PROTO_FILES=(*.proto)

protoc -I=. -I=%GOPATH%/src -I=%GOPATH%/src/github.com/gogo/protobuf/protobuf --gogofaster_out=. ${PROTO_FILES[*]}
protoc -I=. -I=%GOPATH%/src -I=%GOPATH%/src/github.com/gogo/protobuf/protobuf --mgo-go_out=. ${PROTO_FILES[*]}
protoc -I=. -I=%GOPATH%/src -I=%GOPATH%/src/github.com/gogo/protobuf/protobuf --doc_out=html,index.html:. ${PROTO_FILES[*]}

