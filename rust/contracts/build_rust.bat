@echo off
if "%1"=="examples" goto :xit
cd %1
echo Building %1
wasm-pack build
if exist ..\..\..\wasm copy pkg\*.wasm ..\..\..\wasm
if exist ..\..\..\..\wasm copy pkg\*.wasm ..\..\..\..\wasm
cd ..
:xit

