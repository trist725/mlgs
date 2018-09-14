//go:generate PROTO_FILES=(*.proto) protoc -I=. --go_out=. ${PROTO_FILES[*]}
package msg
