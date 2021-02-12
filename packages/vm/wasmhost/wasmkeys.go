package wasmhost

const (
	KeyAddress         = int32(-1)
	KeyAggregateBls    = int32(-2)
	KeyBalances        = int32(-3)
	KeyBase58Bytes     = int32(-4)
	KeyBase58String    = int32(-5)
	KeyCall            = int32(-6)
	KeyCaller          = int32(-7)
	KeyChainOwnerId    = int32(-8)
	KeyColor           = int32(-9)
	KeyContractCreator = int32(-10)
	KeyContractId      = int32(-11)
	KeyDeploy          = int32(-12)
	KeyEvent           = int32(-13)
	KeyExports         = int32(-14)
	KeyHashBlake2b     = int32(-15)
	KeyHashSha3        = int32(-16)
	KeyHname           = int32(-17)
	KeyIncoming        = int32(-18)
	KeyLength          = int32(-19)
	KeyLog             = int32(-20)
	KeyMaps            = int32(-21)
	KeyName            = int32(-22)
	KeyPanic           = int32(-23)
	KeyParams          = int32(-24)
	KeyPost            = int32(-25)
	KeyRandom          = int32(-26)
	KeyResults         = int32(-27)
	KeyReturn          = int32(-28)
	KeyState           = int32(-29)
	KeyTimestamp       = int32(-30)
	KeyTrace           = int32(-31)
	KeyTransfers       = int32(-32)
	KeyUtility         = int32(-33)
	KeyValid           = int32(-34)
	KeyValidBls        = int32(-35)
	KeyValidEd25519    = int32(-36)

	// Treat this one like a version number. When anything changes
	// to the keys give this one a different value and make sure
	// the client side in wasplib is updated accordingly
	KeyZzzzzzz = int32(-37)
)

var keyMap = map[string]int32{
	"address":         KeyAddress,
	"aggregateBls":    KeyAggregateBls,
	"balances":        KeyBalances,
	"base58Bytes":     KeyBase58Bytes,
	"base58String":    KeyBase58String,
	"call":            KeyCall,
	"caller":          KeyCaller,
	"chainOwnerId":    KeyChainOwnerId,
	"color":           KeyColor,
	"contractCreator": KeyContractCreator,
	"contractId":      KeyContractId,
	"deploy":          KeyDeploy,
	"event":           KeyEvent,
	"exports":         KeyExports,
	"hashBlake2b":     KeyHashBlake2b,
	"hashSha3":        KeyHashSha3,
	"hname":           KeyHname,
	"incoming":        KeyIncoming,
	"length":          KeyLength,
	"log":             KeyLog,
	"maps":            KeyMaps,
	"name":            KeyName,
	"panic":           KeyPanic,
	"params":          KeyParams,
	"post":            KeyPost,
	"random":          KeyRandom,
	"results":         KeyResults,
	"return":          KeyReturn,
	"state":           KeyState,
	"timestamp":       KeyTimestamp,
	"trace":           KeyTrace,
	"transfers":       KeyTransfers,
	"utility":         KeyUtility,
	"valid":           KeyValid,
	"validBls":        KeyValidBls,
	"validEd25519":    KeyValidEd25519,
}
