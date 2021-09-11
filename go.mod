module github.com/iotaledger/wasplib

go 1.16

require (
	github.com/bytecodealliance/wasmtime-go v0.21.0
	github.com/iotaledger/goshimmer v0.7.5-0.20210811162925-25c827e8326a
	github.com/iotaledger/hive.go v0.0.0-20210625103722-68b2cf52ef4e
	github.com/iotaledger/wart v0.2.3-0.20210824144406-382ad0e0d608
	github.com/iotaledger/wasp v0.1.1-0.20210908202312-ec434a713fad
	github.com/stretchr/testify v1.7.0
)

replace github.com/iotaledger/wasp v0.1.1-0.20210908202312-ec434a713fad => ../wasp
