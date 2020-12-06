package client

func makeRequest(key Key, function string) ScMutableMap {
	root := ScMutableMap{objId: 1}
	requests := root.GetMapArray(key)
	request := requests.GetMap(requests.Length())
	request.GetString(Key("function")).SetValue(function)
	return request
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	call ScMutableMap
}

func (ctx ScCallInfo) Call() ScCallInfo {
	ctx.call.GetInt(Key("delay")).SetValue(-1)
	return ctx
}

func (ctx ScCallInfo) Contract(contract string) ScCallInfo {
	ctx.call.GetString(Key("contract")).SetValue(contract)
	return ctx
}

func (ctx ScCallInfo) Params() ScMutableMap {
	return ctx.call.GetMap(Key("params"))
}

func (ctx ScCallInfo) Results() ScImmutableMap {
	return ctx.call.GetMap(Key("results")).Immutable()
}

func (ctx ScCallInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.call.GetMap(Key("transfers"))
	transfers.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostInfo struct {
	post ScMutableMap
}

func (ctx ScPostInfo) Chain(chain *ScAddress) ScPostInfo {
	ctx.post.GetAddress(Key("chain")).SetValue(chain)
	return ctx
}

func (ctx ScPostInfo) Contract(contract string) ScPostInfo {
	ctx.post.GetString(Key("contract")).SetValue(contract)
	return ctx
}

func (ctx ScPostInfo) Params() ScMutableMap {
	return ctx.post.GetMap(Key("params"))
}

func (ctx ScPostInfo) Post(delay int64) {
	ctx.post.GetInt(Key("delay")).SetValue(delay)
}

func (ctx ScPostInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.post.GetMap(Key("transfers"))
	transfers.GetInt(color).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewInfo struct {
	view ScMutableMap
}

func (ctx ScViewInfo) Contract(contract string) ScViewInfo {
	ctx.view.GetString(Key("contract")).SetValue(contract)
	return ctx
}

func (ctx ScViewInfo) Params() ScMutableMap {
	return ctx.view.GetMap(Key("params"))
}

func (ctx ScViewInfo) Results() ScImmutableMap {
	return ctx.view.GetMap(Key("results")).Immutable()
}

func (ctx ScViewInfo) View() ScViewInfo {
	ctx.view.GetInt(Key("delay")).SetValue(-2)
	return ctx
}
