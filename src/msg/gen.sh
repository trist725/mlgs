#!/usr/bin/env bash

PROTO_FILES=(*.proto)

protoc -I=. --pbex2-go_out=. ${PROTO_FILES[*]}
protoc -I=. --gogofaster_out=. ${PROTO_FILES[*]}