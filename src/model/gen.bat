@setlocal enabledelayedexpansion enableextensions
@set PROTO_FILES=
@for %%x in (*.proto) do @set PROTO_FILES=!PROTO_FILES! %%x
@set PROTO_FILES=%PROTO_FILES:~1%

protoc -I=. --mgo-go_out=. %PROTO_FILES%
protoc -I=. --gogofaster-go_out=. %PROTO_FILES%