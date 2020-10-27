@echo off
cd %1
wasm-pack build
copy pkg\*.wasm ..\..\..\wasm
cd ..
