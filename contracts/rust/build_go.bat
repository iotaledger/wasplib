@echo off
cd %1
if not exist wasmmain\main.go goto :xit
echo Building %1
schema -go %2
echo compiling %1_go.wasm
if not exist pkg md pkg
tinygo build -o pkg/%1_bg.wasm -target wasm wasmmain/main.go
:xit
cd ..
