@echo off
cd client
set GOOS=js
set GOARCH=wasm
go build -o html-content/wasm/main.wasm main.go

REM tinygo build -target wasm -o html-content/wasm/main.wasm -size full main.go
