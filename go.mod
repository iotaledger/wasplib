module github.com/iotaledger/wasplib

go 1.16

require (
	github.com/bytecodealliance/wasmtime-go v0.21.0
	github.com/iotaledger/goshimmer v0.7.2-0.20210628074845-4f90850164d1
	github.com/iotaledger/hive.go v0.0.0-20210623095912-c1c6f098a6db
	github.com/iotaledger/wart v0.2.2
	github.com/iotaledger/wasp v0.1.1-0.20210707131205-efdee6249c0e
	github.com/stretchr/testify v1.7.0
	github.com/wasmerio/wasmer-go v1.0.3
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
)

replace github.com/iotaledger/wart v0.2.2 => ../wart

replace github.com/iotaledger/wasp v0.1.1-0.20210707131205-efdee6249c0e => ../wasp
