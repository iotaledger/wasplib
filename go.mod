module github.com/iotaledger/wasplib

go 1.16

require (
	github.com/bytecodealliance/wasmtime-go v0.21.0
	github.com/iotaledger/goshimmer v0.6.5-0.20210607123300-2b57453f8501
	github.com/iotaledger/hive.go v0.0.0-20210602094236-4305b15fac6e
	github.com/iotaledger/wart v0.2.2
	github.com/iotaledger/wasp v0.1.1-0.20210609110039-af05fed3987f
	github.com/stretchr/testify v1.7.0
	github.com/wasmerio/wasmer-go v1.0.3
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
)

replace github.com/iotaledger/wart v0.2.2 => ../wart

replace github.com/iotaledger/wasp v0.1.1-0.20210609110039-af05fed3987f => ../wasp
