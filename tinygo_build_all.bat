@echo off
if not exist wasm md wasm
cd contracts
for %%f in (*.go) do tinygo build -o ../wasm/%%~nf_go.wasm -target wasm ./%%~nf.go
cd ..
