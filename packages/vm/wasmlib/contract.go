package wasmlib

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScView struct {
	hContract ScHname
	hFunction ScHname
	paramsId  *int32
	resultsId *int32
}

func NewScView(hContract ScHname, hFunction ScHname) *ScView {
	return &ScView{hContract, hFunction, nil, nil}
}

func (v *ScView) SetPtrs(paramsId *int32, resultsId *int32) {
	v.paramsId = paramsId
	v.resultsId = resultsId
	if paramsId != nil {
		*paramsId = NewScMutableMap().MapId()
	}
}

func (v *ScView) Call() {
	v.call(0)
}

func (v *ScView) call(transferId int32) {
	encode := NewBytesEncoder()
	encode.Hname(v.hContract)
	encode.Hname(v.hFunction)
	encode.Int32(paramsId(v.paramsId))
	encode.Int32(transferId)
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	if v.resultsId != nil {
		*v.resultsId = GetObjectId(1, KeyReturn, TYPE_MAP)
	}
}

func (v *ScView) OfContract(hContract ScHname) *ScView {
	v.hContract = hContract
	return v
}

func paramsId(id *int32) int32 {
	if id == nil {
		return 0
	}
	return *id
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScFunc struct {
	ScView
	delay      int32
	transferId int32
}

func NewScFunc(hContract ScHname, hFunction ScHname) *ScFunc {
	return &ScFunc{ScView{hContract, hFunction, nil, nil}, 0, 0}
}

func (f *ScFunc) Call() {
	if f.delay != 0 {
		Panic("cannot delay a call")
	}
	f.call(f.transferId)
}

func (f *ScFunc) Delay(seconds int32) *ScFunc {
	f.delay = seconds
	return f
}

func (f *ScFunc) Post() {
	f.PostToChain(Root.GetChainId(KeyChainId).Value())
}

func (f *ScFunc) PostToChain(chainId ScChainId) {
	if f.transferId == 0 {
		Panic("transfer is required for post")
	}
	encode := NewBytesEncoder()
	encode.ChainId(chainId)
	encode.Hname(f.hContract)
	encode.Hname(f.hFunction)
	encode.Int32(paramsId(f.paramsId))
	encode.Int32(f.transferId)
	encode.Int32(f.delay)
	Root.GetBytes(KeyPost).SetValue(encode.Data())
}

func (f *ScFunc) Transfer(transfer ScTransfers) *ScFunc {
	f.transferId = transfer.transfers.MapId()
	return f
}

func (f *ScFunc) TransferIotas(amount int64) *ScFunc {
	return f.Transfer(NewScTransferIotas(amount))
}
