//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com\gogo\protobuf\protobuf --mgo-go_out=. *.proto
//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com\gogo\protobuf\protobuf --gogofaster_out=. *.proto
//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com\gogo\protobuf\protobuf --doc_out=. *.proto
package model
