#!/usr/bin/env bash

PROTO_FILES=(*.proto)

protoc -I=. --enum-go_out=. ${PROTO_FILES[*]}
protoc -I=. --mgo-go_out=. ${PROTO_FILES[*]}
protoc --doc_out=html,index.html:. ${PROTO_FILES[*]}

