@echo off
cd %1
if not exist src\lib.rs goto :xit
rem echo Building %1
rem schema
echo compiling %1_bg.wasm
wasm-pack build
:xit
cd ..
