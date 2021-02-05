@echo off
if not exist wasm md wasm
cd contracts
for %%f in (*.go) do tinygo build -o ../wasm/%%~nf_go.wasm -target wasm ./%%~nf.go
cd ..
cd rust\contracts
for /d %%f in (*.) do if exist %%f\test\%%f\wasm\%%f.go tinygo build -o ../../wasm/%%f_go.wasm -target wasm %%f\test\%%f\wasm\%%f.go
cd ..\..
