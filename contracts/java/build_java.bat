for %%g in (out\production\contracts\org\iota\wasp\contracts\%1\lib\*Thunk.class) do java -jar bytecoder.jar -classpath=out\production\contracts -mainclass=org.iota.wasp.contracts.%1.lib.%%~ng -builddirectory=out -backend=wasm -minify=false
copy /y out\bytecoder.wasm ..\rust\%1\pkg\%1_bg.wasm
