//go:generate protoc -I=. -I=$GOPATH/src --mgo-go_out=. account.proto user.proto chat.proto ship.proto
//go:generate protoc --doc_out=html,index.html:. account.proto user.proto chat.proto ship.proto
package model
