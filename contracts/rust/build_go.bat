@echo off
cd %1
if not exist wasmmain\%1.go goto :xit
echo Building %1
schema -go
echo compiling %1_go.wasm
if not exist pkg md pkg
tinygo build -o pkg/%1_bg.wasm -target wasm wasmmain/%1.go
:xit
cd ..
