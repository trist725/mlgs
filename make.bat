@set WORK_DIR=%~dp0
@set OUTPUT_DIR=%WORK_DIR%\bin
@set LOG_DIR=%WORK_DIR%\log

@set GOPATH=%WORK_DIR%

@set RPC_DIR=%WORK_DIR%\src\rpc
@set MODEL_DIR=%WORK_DIR%\src\model
@set GAME_DIR=%WORK_DIR%\src\game
@set LOGIN_DIR=%WORK_DIR%\src\login
@set WORLD_DIR=%WORK_DIR%\src\world

@IF "%1" == "" call :all & cd %cd% & goto exit
@IF "%1" == "clean" call :clean & goto exit
@IF "%1" == "clean-log" call :clean-log & goto exit
@IF "%1" == "glide-up" @call make_go.bat glide-up & cd %cd% & goto exit
@IF "%1" == "zip-vendor" call :zip-vendor & goto exit
@IF "%1" == "unzip-vendor" @call make_go.bat unzip-vendor & goto exit
@IF "%1" == "rpc" call :rpc & cd %cd% & goto exit
@IF "%1" == "model" call :model & cd %cd% & goto exit
@IF "%1" == "gateway" call :gateway & cd %cd% & goto exit
@IF "%1" == "game" call :game & cd %cd% & goto exit
@IF "%1" == "game-msg" call :game-msg & cd %cd% & goto exit
@IF "%1" == "game-sd" call :game-sd & cd %cd% & goto exit
@IF "%1" == "game-cache" call :game-cache & cd %cd% & goto exit
@IF "%1" == "login-msg" call :login-msg & cd %cd% & goto exit
@IF "%1" == "login" make_go build login & cd %cd% & goto exit
@IF "%1" == "robot" call :robot & cd %cd% & goto exit

@echo unsupported operate [%1]

@goto exit


:clean
@call make_go.bat clean
del "%WORK_DIR%\src\game\debug"
del "%WORK_DIR%\src\gateway\debug"
del "%WORK_DIR%\src\robot\debug"
del "%WORK_DIR%\src\login\debug"
@goto exit


:clean-log
@call make_go.bat clean-log
rmdir /q /s "%LOG_DIR%"
rmdir /q /s "%WORK_DIR%\src\gateway\log\*.log"
rmdir /q /s "%WORK_DIR%\src\game\log\*.log"
rmdir /q /s "%WORK_DIR%\src\robot\log\*.log"
rmdir /q /s "%WORK_DIR%\src\login\log\*.log"
@goto exit


:zip-vendor
del "%WORK_DIR%\src\vendor\gitee.com\lwj8507\nggs\vendor.zip"
del "%WORK_DIR%\src\vendor\gitee.com\lwj8507\light-protoactor-go\vendor.zip"
rmdir /q /s "%WORK_DIR%\src\vendor\github.com\coreos\etcd\cmd"
@call make_go.bat zip-vendor
@goto exit


:rpc
@echo build [rpc] begin
cd %RPC_DIR%
go generate
go test
@cd %WORK_DIR%
@echo build [rpc] end
@goto exit


:model
@echo build [model begin
cd %MODEL_DIR%
go generate
go test
@cd %WORK_DIR%
@echo build [model] end
@goto exit


:game-msg
@echo build [game/msg] begin
cd %GAME_DIR%\msg
go generate
go test
@cd %WORK_DIR%
@echo build [game/msg] end
@goto exit


:game-sd
@echo build [game/sd] begin
cd %GAME_DIR%\sd
go generate
go test
@cd %WORK_DIR%
@echo build [game/sd] end
@goto exit


:game-cache
@echo build [game/cache begin
cd %GAME_DIR%\cache
go generate
go test
cd %GAME_DIR%\cache\battle
go generate
go test
@cd %WORK_DIR%
@echo build [game/cache] end
@goto exit


:game
@cd %WORK_DIR%
@call make_go.bat build game
@goto exit


:robot
@cd %WORK_DIR%
@call make_go.bat build robot
@goto exit


:gateway
@cd %WORK_DIR%
@call make_go.bat build gateway
@goto exit


:login-msg
@echo build [login/msg] begin
cd %LOGIN_DIR%\msg
go generate
go test
@cd %WORK_DIR%
@echo build [login/msg] end
@goto exit


:login
@cd %WORK_DIR%
@call make_go.bat build login
@goto exit


:all
@echo build [all] begin
@call :rpc
@call :model
@call :gateway
@call :login-msg
@call :login
@call :game-msg
@call :game-sd
@call :game-cache
@call :game
@call :robot
@echo build [all] end
@goto exit


:exit
