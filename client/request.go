package client

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScBaseInfo struct {
	request ScMutableMap
}

func (ctx ScBaseInfo) Contract(contract string) ScBaseInfo {
	ctx.request.GetString(Key("contract")).SetValue(contract)
	return ctx
}

func (ctx ScBaseInfo) exec(delay int64) {
	ctx.request.GetInt(Key("delay")).SetValue(delay)
}

func (ctx ScBaseInfo) Params() ScMutableMap {
	return ctx.request.GetMap(Key("params"))
}

func (ctx ScBaseInfo) results() ScImmutableMap {
	return ctx.request.GetMap(Key("results")).Immutable()
}

func (ctx ScBaseInfo) transfer(color *ScColor, amount int64) {
	transfers := ctx.request.GetMap(Key("transfers"))
	transfers.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	ScBaseInfo
}

func (ctx ScCallInfo) Call() ScCallInfo {
	ctx.exec(-1)
	return ctx
}

func (ctx ScCallInfo) Results() ScImmutableMap {
	return ctx.ScBaseInfo.results()
}

func (ctx ScCallInfo) Transfer(color *ScColor, amount int64) {
	ctx.ScBaseInfo.transfer(color, amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostInfo struct {
	ScBaseInfo
}

func (ctx ScPostInfo) Chain(chain *ScAddress) ScPostInfo {
	ctx.request.GetAddress(Key("chain")).SetValue(chain)
	return ctx
}

func (ctx ScPostInfo) Post(delay int64) {
	ctx.exec(delay)
}

func (ctx ScPostInfo) Transfer(color *ScColor, amount int64) {
	ctx.ScBaseInfo.transfer(color, amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewInfo struct {
	ScBaseInfo
}

func (ctx ScViewInfo) Results() ScImmutableMap {
	return ctx.ScBaseInfo.results()
}

func (ctx ScViewInfo) View() ScViewInfo {
	ctx.exec(-2)
	return ctx
}
