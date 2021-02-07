@echo off
if not exist ..\wasm md ..\wasm
for /d %%f in (*.) do call build_go.bat %%f
