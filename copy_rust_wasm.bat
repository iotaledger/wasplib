@echo off
if not exist wasm md wasm
cd rust
for /D %%f in (*.) do if exist %%f\pkg\%%f_bg.wasm copy /Y %%f\pkg\%%f_bg.wasm ..\wasm
cd ..

