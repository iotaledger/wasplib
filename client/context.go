package client

type ScAccount struct {
	account ScImmutableMap
}

func (ctx ScAccount) Balance(color string) int64 {
	return ctx.account.GetMap("balance").GetInt(color).Value()
}

func (ctx ScAccount) Colors() ScImmutableStringArray {
	return ctx.account.GetStringArray("colors")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContract struct {
	contract ScImmutableMap
}

func (ctx ScContract) Address() string {
	return ctx.contract.GetString("address").Value()
}

func (ctx ScContract) Color() string {
	return ctx.contract.GetString("color").Value()
}

func (ctx ScContract) Description() string {
	return ctx.contract.GetString("description").Value()
}

func (ctx ScContract) Id() string {
	return ctx.contract.GetString("id").Value()
}

func (ctx ScContract) Name() string {
	return ctx.contract.GetString("name").Value()
}

func (ctx ScContract) Owner() string {
	return ctx.contract.GetString("owner").Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScEvent struct {
	event ScMutableMap
}

func (ctx ScEvent) Code(code int64) {
	ctx.event.GetInt("code").SetValue(code)
}

func (ctx ScEvent) Contract(contract string) {
	ctx.event.GetString("contract").SetValue(contract)
}

func (ctx ScEvent) Delay(delay int64) {
	ctx.event.GetInt("delay").SetValue(delay)
}

func (ctx ScEvent) Function(function string) {
	ctx.event.GetString("function").SetValue(function)
}

func (ctx ScEvent) Params() ScMutableMap {
	return ctx.event.GetMap("params")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMap
}

func (ctx ScLog) Append(timestamp int64, data []byte) {
	ctx.log.GetInt("timestamp").SetValue(timestamp)
	ctx.log.GetBytes("data").SetValue(data)
}

func (ctx ScLog) Length() int32 {
	return int32(ctx.log.GetInt("length").Value())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequest struct {
	request ScImmutableMap
}

func (ctx ScRequest) Address() string {
	return ctx.request.GetString("address").Value()
}

func (ctx ScRequest) Balance(color string) int64 {
	return ctx.request.GetMap("balance").GetInt(color).Value()
}

func (ctx ScRequest) Colors() ScImmutableStringArray {
	return ctx.request.GetStringArray("colors")
}

func (ctx ScRequest) Hash() string {
	return ctx.request.GetString("hash").Value()
}

func (ctx ScRequest) Id() string {
	return ctx.request.GetString("id").Value()
}

func (ctx ScRequest) Params() ScImmutableMap {
	return ctx.request.GetMap("params")
}

func (ctx ScRequest) Timestamp() int64 {
	return ctx.request.GetInt("timestamp").Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfer struct {
	transfer ScMutableMap
}

func (ctx ScTransfer) Address(address string) {
	ctx.transfer.GetString("address").SetValue(address)
}

func (ctx ScTransfer) Amount(amount int64) {
	ctx.transfer.GetInt("amount").SetValue(amount)
}

func (ctx ScTransfer) Color(color string) {
	ctx.transfer.GetString("color").SetValue(color)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Hash(value string) string {
	hash := ctx.utility.GetString("hash")
	hash.SetValue(value)
	return hash.Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContext struct {
	root ScMutableMap
}

func NewScContext() ScContext {
	return ScContext{root: ScMutableMap{objId: 1}}
}

func (ctx ScContext) Account() ScAccount {
	return ScAccount{ctx.root.GetMap("account").Immutable()}
}

func (ctx ScContext) Contract() ScContract {
	return ScContract{ctx.root.GetMap("contract").Immutable()}
}

func (ctx ScContext) Error() ScMutableString {
	return ctx.root.GetString("error")
}

func (ctx ScContext) Event(contract string, function string, delay int64) ScMutableMap {
	events := ctx.root.GetMapArray("events")
	evt := ScEvent{events.GetMap(events.Length())}
	evt.Contract(contract)
	evt.Function(function)
	evt.Delay(delay)
	return evt.Params()
}

// just for compatibility with old hardcoded SCs
func (ctx ScContext) EventWithCode(contract string, code int64, delay int64) ScMutableMap {
	events := ctx.root.GetMapArray("events")
	evt := ScEvent{events.GetMap(events.Length())}
	evt.Contract(contract)
	evt.Code(code)
	evt.Delay(delay)
	return evt.Params()
}

func (ctx ScContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScContext) Random(max int64) int64 {
	return int64(uint64(ctx.root.GetInt("random").Value()) % uint64(max))
}

func (ctx ScContext) Request() ScRequest {
	return ScRequest{ctx.root.GetMap("request").Immutable()}
}

func (ctx ScContext) State() ScMutableMap {
	return ctx.root.GetMap("state")
}

func (ctx ScContext) TimestampedLog(key string) ScLog {
	return ScLog{ctx.root.GetMap("logs").GetMap(key)}
}

func (ctx ScContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScContext) Transfer(address string, color string, amount int64) {
	transfers := ctx.root.GetMapArray("transfers")
	xfer := ScTransfer{transfers.GetMap(transfers.Length())}
	xfer.Address(address)
	xfer.Color(color)
	xfer.Amount(amount)
}

func (ctx ScContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}
