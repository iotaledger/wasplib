package main

import (
	"fmt"
	"github.com/iotaledger/wasplib/host"
	"github.com/iotaledger/wasplib/host/interfaces/objtype"
)

func main() {
	fmt.Println("Hello, WaspLib!")

	ctx := host.NewHostImpl()
	ctx.LoadWasm("wasm/fairroulette_bg.wasm")

	//set up placeBet
	host.EnableImmutableChecks = false
	contract := ctx.Object(nil, "contract", objtype.OBJTYPE_MAP)
	contract.SetString(ctx.GetKeyId("address"), "smartContractAddress")
	request := ctx.Object(nil, "request", objtype.OBJTYPE_MAP)
	request.SetString(ctx.GetKeyId("hash"), "requestTransactionHash")
	request.SetString(ctx.GetKeyId("address"), "requestInitiatorAddress")
	params := ctx.Object(request, "params", objtype.OBJTYPE_MAP)
	params.SetInt(ctx.GetKeyId("color"), 3)
	ctx.AddBalance(request, "iota", 500)
	host.EnableImmutableChecks = true

	ctx.RunWasmFunction("placeBet")
}
