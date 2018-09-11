#!/usr/bin/env bash

PROTO_FILES=(*.proto)

protoc -I=. --go_out=. ${PROTO_FILES[*]}