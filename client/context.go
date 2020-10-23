package client

type ScAccount struct {
	account ScImmutableMap
}

func (ctx ScAccount) Balance(color *ScColor) int64 {
	return ctx.account.GetMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScAccount) Colors() ScColors {
	return ScColors{colors: ctx.account.GetStringArray("colors")}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColors struct {
	colors ScImmutableStringArray
}

func (ctx ScColors) GetColor(index int32) *ScColor {
	return NewScColor(ctx.colors.GetString(index).Value())
}

func (ctx ScColors) Length() int32 {
	return ctx.colors.Length()
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

type ScExports struct {
	exports ScMutableStringArray
	next    int32
}

func NewScExports() ScExports {
	root := ScMutableMap{objId: 1}
	return ScExports{root.GetStringArray("exports"), 0}
}

func (ctx *ScExports) Add(name string) {
	ctx.next++
	ctx.exports.GetString(ctx.next).SetValue(name)
}

func (ctx *ScExports) AddProtected(name string) {
	ctx.next++
	ctx.exports.GetString(ctx.next | 0x4000).SetValue(name)
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

type ScPostedRequest struct {
	request ScMutableMap
}

func (ctx ScPostedRequest) Code(code int64) {
	ctx.request.GetInt("code").SetValue(code)
}

func (ctx ScPostedRequest) Contract(contract string) {
	ctx.request.GetString("contract").SetValue(contract)
}

func (ctx ScPostedRequest) Delay(delay int64) {
	ctx.request.GetInt("delay").SetValue(delay)
}

func (ctx ScPostedRequest) Function(function string) {
	ctx.request.GetString("function").SetValue(function)
}

func (ctx ScPostedRequest) Params() ScMutableMap {
	return ctx.request.GetMap("params")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequest struct {
	request ScImmutableMap
}

func (ctx ScRequest) Address() string {
	return ctx.request.GetString("address").Value()
}

func (ctx ScRequest) Balance(color *ScColor) int64 {
	return ctx.request.GetMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScRequest) Colors() ScColors {
	return ScColors{colors: ctx.request.GetStringArray("colors")}
}

func (ctx ScRequest) Id() string {
	return ctx.request.GetString("id").Value()
}

func (ctx ScRequest) MintedColor() *ScColor {
	return &ScColor{color: ctx.request.GetString("hash").Value()}
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

func (ctx ScTransfer) Color(color *ScColor) {
	ctx.transfer.GetString("color").SetValue(color.Bytes())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Hash(value []byte) []byte {
	hash := ctx.utility.GetBytes("hash")
	hash.SetValue(value)
	return hash.Value()
}

func (ctx ScUtility) Random(max int64) int64 {
	return int64(uint64(ctx.utility.GetInt("random").Value()) % uint64(max))
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

func (ctx ScContext) Log(text string) {
	SetString(1, KeyLog(), text)
}

func (ctx ScContext) PostRequest(contract string, function string, delay int64) ScMutableMap {
	postedRequests := ctx.root.GetMapArray("postedRequests")
	request := ScPostedRequest{postedRequests.GetMap(postedRequests.Length())}
	request.Contract(contract)
	request.Function(function)
	request.Delay(delay)
	return request.Params()
}

// just for compatibility with old hardcoded SCs
func (ctx ScContext) PostRequestWithCode(contract string, code int64, delay int64) ScMutableMap {
	postedRequests := ctx.root.GetMapArray("postedRequests")
	request := ScPostedRequest{postedRequests.GetMap(postedRequests.Length())}
	request.Contract(contract)
	request.Code(code)
	request.Delay(delay)
	return request.Params()
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

func (ctx ScContext) Transfer(address string, color *ScColor, amount int64) {
	transfers := ctx.root.GetMapArray("transfers")
	xfer := ScTransfer{transfers.GetMap(transfers.Length())}
	xfer.Address(address)
	xfer.Color(color)
	xfer.Amount(amount)
}

func (ctx ScContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}
