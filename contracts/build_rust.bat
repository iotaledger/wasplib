@echo off
cd %1
if not exist src\lib.rs goto :xit
echo Building %1
schema
echo compiling %1_bg.wasm
wasm-pack build
if exist pkg/%1_go.wasm if exist ..\..\wasm xcopy /y /q pkg\%1_bg.wasm ..\..\wasm
:xit
cd ..
