module github.com/iotaledger/wasplib

go 1.16

require (
	github.com/bytecodealliance/wasmtime-go v0.21.0
	github.com/iotaledger/goshimmer v0.7.4
	github.com/iotaledger/hive.go v0.0.0-20210625103722-68b2cf52ef4e
	github.com/iotaledger/wart v0.2.2
	github.com/iotaledger/wasp v0.1.1-0.20210817223024-51e4eeca5104
	github.com/stretchr/testify v1.7.0
	github.com/wasmerio/wasmer-go v1.0.4-0.20210722072119-4d063a16fde3
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
)

replace github.com/iotaledger/wart v0.2.2 => ../wart

replace github.com/iotaledger/wasp v0.1.1-0.20210817223024-51e4eeca5104 => ../wasp
