java -jar ..\..\..\bytecoder.jar -classpath=. -mainclass=org.iota.wasp.contracts.%1.lib.%2Thunk -builddirectory=. -backend=wasm -minify=false
del %1_bg.wasm
ren bytecoder.wasm %1_bg.wasm
