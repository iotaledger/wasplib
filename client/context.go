// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

type ScAccount struct {
	account ScImmutableMap
}

func (ctx ScAccount) Balance(color *ScColor) int64 {
	return ctx.account.GetKeyMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScAccount) Colors() ScImmutableColorArray {
	return ctx.account.GetColorArray("colors")
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScContract struct {
	contract ScImmutableMap
}

func (ctx ScContract) Color() *ScColor {
	return ctx.contract.GetColor("color").Value()
}

func (ctx ScContract) Description() string {
	return ctx.contract.GetString("description").Value()
}

func (ctx ScContract) Id() *ScAgent {
	return ctx.contract.GetAgent("id").Value()
}

func (ctx ScContract) Name() string {
	return ctx.contract.GetString("name").Value()
}

func (ctx ScContract) Owner() *ScAgent {
	return ctx.contract.GetAgent("owner").Value()
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

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScLog struct {
	log ScMutableMapArray
}

func (ctx ScLog) Append(timestamp int64, data []byte) {
	logEntry := ctx.log.GetMap(ctx.log.Length())
	logEntry.GetInt("timestamp").SetValue(timestamp)
	logEntry.GetBytes("data").SetValue(data)
}

func (ctx ScLog) Length() int32 {
	return int32(ctx.log.Length())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScPostedRequest struct {
	request ScMutableMap
}

func (ctx ScPostedRequest) Code(code int64) {
	ctx.request.GetInt("code").SetValue(code)
}

func (ctx ScPostedRequest) Contract(contract *ScAgent) {
	ctx.request.GetAgent("contract").SetValue(contract)
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

func (ctx ScRequest) Balance(color *ScColor) int64 {
	return ctx.request.GetKeyMap("balance").GetInt(color.Bytes()).Value()
}

func (ctx ScRequest) Colors() ScImmutableColorArray {
	return ctx.request.GetColorArray("colors")
}

func (ctx ScRequest) From(originator *ScAgent) bool {
	return ctx.Sender().Equals(originator)
}

func (ctx ScRequest) Id() *ScRequestId {
	return ctx.request.GetRequestId("id").Value()
}

func (ctx ScRequest) MintedColor() *ScColor {
	return ctx.request.GetColor("hash").Value()
}

func (ctx ScRequest) Params() ScImmutableMap {
	return ctx.request.GetMap("params")
}

func (ctx ScRequest) Sender() *ScAgent {
	return ctx.request.GetAgent("sender").Value()
}

func (ctx ScRequest) Timestamp() int64 {
	return ctx.request.GetInt("timestamp").Value()
}

func (ctx ScRequest) TxHash() *ScTxHash {
	return ctx.request.GetTxHash("hash").Value()
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfer struct {
	transfer ScMutableMap
}

func (ctx ScTransfer) Agent(address *ScAgent) {
	ctx.transfer.GetAgent("agent").SetValue(address)
}

func (ctx ScTransfer) Amount(amount int64) {
	ctx.transfer.GetInt("amount").SetValue(amount)
}

func (ctx ScTransfer) Color(color *ScColor) {
	ctx.transfer.GetColor("color").SetValue(color)
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScUtility struct {
	utility ScMutableMap
}

func (ctx ScUtility) Base58Decode(value string) []byte {
	//TODO atomic set/get
	decode := ctx.utility.GetString("base58")
	encode := ctx.utility.GetBytes("base58")
	decode.SetValue(value)
	return encode.Value()
}

func (ctx ScUtility) Base58Encode(value []byte) string {
	//TODO atomic set/get
	decode := ctx.utility.GetString("base58")
	encode := ctx.utility.GetBytes("base58")
	encode.SetValue(value)
	return decode.Value()
}

func (ctx ScUtility) Hash(value []byte) []byte {
	//TODO atomic set/get
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

func (ctx ScContext) PostRequest(contract *ScAgent, function string, delay int64) ScMutableMap {
	postedRequests := ctx.root.GetMapArray("postedRequests")
	request := ScPostedRequest{postedRequests.GetMap(postedRequests.Length())}
	request.Contract(contract)
	request.Function(function)
	request.Delay(delay)
	return request.Params()
}

// just for compatibility with old hardcoded SCs
func (ctx ScContext) PostRequestWithCode(contract *ScAgent, code int64, delay int64) ScMutableMap {
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
	return ScLog{ctx.root.GetMap("logs").GetMapArray(key)}
}

func (ctx ScContext) Trace(text string) {
	SetString(1, KeyTrace(), text)
}

func (ctx ScContext) Transfer(agent *ScAgent, color *ScColor, amount int64) {
	transfers := ctx.root.GetMapArray("transfers")
	xfer := ScTransfer{transfers.GetMap(transfers.Length())}
	xfer.Agent(agent)
	xfer.Color(color)
	xfer.Amount(amount)
}

func (ctx ScContext) Utility() ScUtility {
	return ScUtility{ctx.root.GetMap("utility")}
}
