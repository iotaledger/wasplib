@echo off
if not exist wasm md wasm
cd rust\contracts
for /D %%f in (*.) do call build_rust.bat %%f
cd examples
for /D %%f in (*.) do call ..\build_rust.bat %%f
cd ..\..\..


