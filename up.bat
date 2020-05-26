@set WORK_DIR=%~dp0/bin

@set PREFIX=mlgs-
@set SUFFIX=.exe
@set EXENAME=%PREFIX%%date:~0,4%%date:~5,2%%date:~8,2%_%time:~0,2%%time:~3,2%%time:~6,2%%SUFFIX%
@set EXEPATH=bin/%EXENAME%

@IF "%1" == "" call :mod-tidy & call :build & cd %WORK_DIR% & call :run & goto :exit

@IF "%1" == "mod-tidy" call :mod-tidy & cd %WORK_DIR% & goto :exit

@IF "%1" == "kill" call :kill & goto :exit

@echo unsupported operate [%1]

@goto :exit


:build
go build -race -o %EXEPATH% src/main.go
@goto :exit

:run
start /b %EXENAME%
@goto :exit

:mod-tidy
go mod tidy
@goto :exit

:kill
taskkill /f /t /im %PREFIX%*
@goto :exit

:exit
