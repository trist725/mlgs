@setlocal enabledelayedexpansion enableextensions
@set PROTO_FILES=
@for %%x in (*.proto) do @set PROTO_FILES=!PROTO_FILES! %%x
@set PROTO_FILES=%PROTO_FILES:~1%

protoc -I=. -I=%GOPATH%/src -I=%GOPATH%/src/github.com\gogo\protobuf\protobuf --mgo-go_out=. %PROTO_FILES%
protoc -I=. -I=%GOPATH%/src -I=%GOPATH%/src/github.com\gogo\protobuf\protobuf --gogofaster_out=. %PROTO_FILES%