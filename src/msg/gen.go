//go:generate protoc -I=. --pbex-go_out=. *.proto
//go:generate protoc -I=. --gogofaster_out=. *.proto
//go:generate protoc -I=. --doc_out=. *.proto
package msg
