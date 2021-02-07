@echo off
if not exist ..\wasm md ..\wasm
for /d %%f in (*.) do call build_rust.bat %%f

