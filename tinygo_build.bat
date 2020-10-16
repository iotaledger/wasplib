@echo off
if not exist wasm md wasm
tinygo build -o wasm/%1_go.wasm -target wasm ./contracts/%1.go
