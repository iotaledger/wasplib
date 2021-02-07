@echo off
cd %1
if not exist wasmmain\%1.go goto :xit
echo Building %1
schema
echo compiling %1_go.wasm
if not exist pkg md pkg
tinygo build -o pkg/%1_go.wasm -target wasm wasmmain/%1.go
if exist pkg/%1_go.wasm if exist ..\..\wasm xcopy /y /q pkg\%1_go.wasm ..\..\wasm
:xit
cd ..
