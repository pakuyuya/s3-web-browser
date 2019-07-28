@echo off

setlocal

rd /Q /S .\dist

mkdir .\dist
mkdir .\dist\static
mkdir .\dist\templates
mkdir .\dist\testdatas

copy /Y .\setting.yml .\dist\

dep ensure
go build -o dist/server.exe main.go

set BUILD_RESULT=%errorlevel%
if %BUILD_RESULT% neq 0 (
  echo `go build` is failed.
  exit /b %BUILD_RESULT%
)

echo server/build.bat is successful

endlocal
