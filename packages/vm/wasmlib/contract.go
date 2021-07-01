package wasmlib

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScView struct {
	hContract ScHname
	hFunction ScHname
	paramsID  *int32
	resultsID *int32
}

func NewScView(hContract, hFunction ScHname) *ScView {
	return &ScView{hContract, hFunction, nil, nil}
}

func (v *ScView) SetPtrs(paramsID, resultsID *int32) {
	v.paramsID = paramsID
	v.resultsID = resultsID
	if paramsID != nil {
		*paramsID = NewScMutableMap().MapID()
	}
}

func (v *ScView) Call() {
	v.call(0)
}

func (v *ScView) call(transferID int32) {
	encode := NewBytesEncoder()
	encode.Hname(v.hContract)
	encode.Hname(v.hFunction)
	encode.Int32(paramsID(v.paramsID))
	encode.Int32(transferID)
	Root.GetBytes(KeyCall).SetValue(encode.Data())
	if v.resultsID != nil {
		*v.resultsID = GetObjectID(OBJ_ID_ROOT, KeyReturn, TYPE_MAP)
	}
}

func (v *ScView) OfContract(hContract ScHname) *ScView {
	v.hContract = hContract
	return v
}

func paramsID(id *int32) int32 {
	if id == nil {
		return 0
	}
	return *id
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScFunc struct {
	ScView
	delay      int32
	transferID int32
}

func NewScFunc(hContract, hFunction ScHname) *ScFunc {
	return &ScFunc{ScView{hContract, hFunction, nil, nil}, 0, 0}
}

func (f *ScFunc) Call() {
	if f.delay != 0 {
		Panic("cannot delay a call")
	}
	f.call(f.transferID)
}

func (f *ScFunc) Delay(seconds int32) *ScFunc {
	f.delay = seconds
	return f
}

func (f *ScFunc) Post() {
	f.PostToChain(Root.GetChainID(KeyChainID).Value())
}

func (f *ScFunc) PostToChain(chainID ScChainID) {
	if f.transferID == 0 {
		Panic("transfer is required for post")
	}
	encode := NewBytesEncoder()
	encode.ChainID(chainID)
	encode.Hname(f.hContract)
	encode.Hname(f.hFunction)
	encode.Int32(paramsID(f.paramsID))
	encode.Int32(f.transferID)
	encode.Int32(f.delay)
	Root.GetBytes(KeyPost).SetValue(encode.Data())
}

func (f *ScFunc) Transfer(transfer ScTransfers) *ScFunc {
	f.transferID = transfer.transfers.MapID()
	return f
}

func (f *ScFunc) TransferIotas(amount int64) *ScFunc {
	return f.Transfer(NewScTransferIotas(amount))
}
