package wasmlib

type ScContractFunc struct {
	ctx      ScFuncContext
	chainId  ScChainId
	contract ScHname
	delay    int32
	post     bool
	results  ScImmutableMap
}

func NewScContractFunc(ctx ScFuncContext, contract ScHname) ScContractFunc {
	return ScContractFunc{ctx: ctx, chainId: ctx.ChainId(), contract: contract}
}

func (f *ScContractFunc) Delay(seconds int32) {
	f.delay = seconds
}

func (f *ScContractFunc) OfContract(contract ScHname) {
	f.contract = contract
}

func (f *ScContractFunc) Post() {
	f.post = true
}

func (f *ScContractFunc) PostToChain(chainId ScChainId) {
	f.post = true
	f.chainId = chainId
}

func (f *ScContractFunc) ResultMapId() int32 {
	mapId := f.results.MapId()
	f.ctx.Require(mapId != 0, "Cannot get results from asynchronous post")
	return mapId
}

func (f *ScContractFunc) Run(function ScHname, paramsId int32, transfer *ScTransfers) {
	params := &ScMutableMap{objId: paramsId}
	if f.post {
		f.ctx.Require(transfer != nil, "Cannot post to view")
		f.ctx.Post(f.chainId, f.contract, function, params, *transfer, f.delay)
		return
	}

	f.results = f.ctx.Call(f.contract, function, params, transfer)
}

type ScContractView struct {
	ctx      ScViewContext
	contract ScHname
	results  ScImmutableMap
}

func NewScContractView(ctx ScViewContext, contract ScHname) ScContractView {
	return ScContractView{ctx: ctx, contract: contract}
}

func (v *ScContractView) OfContract(contract ScHname) {
	v.contract = contract
}

func (v *ScContractView) ResultMapId() int32 {
	return v.results.MapId()
}

func (v *ScContractView) Run(function ScHname, paramsId int32) {
	params := &ScMutableMap{objId: paramsId}
	v.results = v.ctx.Call(v.contract, function, params)
}
