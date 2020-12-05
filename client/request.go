package client

func makeRequest(key string, function string) ScMutableMap {
	root := ScMutableMap{objId: 1}
	requests := root.GetMapArray(key)
	request := requests.GetMap(requests.Length())
	request.GetString("function").SetValue(function)
	return request
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScCallInfo struct {
	call ScMutableMap
}

func (ctx ScCallInfo) Call() ScCallInfo {
	ctx.call.GetInt("delay").SetValue(-1)
	return ctx
}

func (ctx ScCallInfo) Contract(contract string) ScCallInfo {
	ctx.call.GetString("contract").SetValue(contract)
	return ctx
}

func (ctx ScCallInfo) Params() ScMutableMap {
	return ctx.call.GetMap("params")
}

func (ctx ScCallInfo) Results() ScImmutableMap {
	return ctx.call.GetMap("results").Immutable()
}

func (ctx ScCallInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.call.GetKeyMap("transfers")
	transfers.GetInt(color.Bytes()).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostInfo struct {
	post ScMutableMap
}

func (ctx ScPostInfo) Chain(chain *ScAddress) ScPostInfo {
	ctx.post.GetAddress("chain").SetValue(chain)
	return ctx
}

func (ctx ScPostInfo) Contract(contract string) ScPostInfo {
	ctx.post.GetString("contract").SetValue(contract)
	return ctx
}

func (ctx ScPostInfo) Params() ScMutableMap {
	return ctx.post.GetMap("params")
}

func (ctx ScPostInfo) Post(delay int64) {
	ctx.post.GetInt("delay").SetValue(delay)
}

func (ctx ScPostInfo) Transfer(color *ScColor, amount int64) {
	transfers := ctx.post.GetKeyMap("transfers")
	transfers.GetInt(color.Bytes()).SetValue(amount)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScViewInfo struct {
	view ScMutableMap
}

func (ctx ScViewInfo) Contract(contract string) ScViewInfo {
	ctx.view.GetString("contract").SetValue(contract)
	return ctx
}

func (ctx ScViewInfo) Params() ScMutableMap {
	return ctx.view.GetMap("params")
}

func (ctx ScViewInfo) Results() ScImmutableMap {
	return ctx.view.GetMap("results").Immutable()
}

func (ctx ScViewInfo) View() ScViewInfo {
	ctx.view.GetInt("delay").SetValue(-2)
	return ctx
}
