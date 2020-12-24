package wasmhost

// @formatter:off
const (
	KeyAgent       = int32(-1)
	KeyAmount      = KeyAgent       -1
	KeyBalances    = KeyAmount      -1
	KeyBase58      = KeyBalances    -1
	KeyCaller      = KeyBase58      -1
	KeyCalls       = KeyCaller      -1
	KeyChain       = KeyCalls       -1
	KeyColor       = KeyChain       -1
	KeyContract    = KeyColor       -1
	KeyData        = KeyContract    -1
	KeyDelay       = KeyData        -1
	KeyDescription = KeyDelay       -1
	KeyError       = KeyDescription -1
	KeyExports     = KeyError       -1
	KeyFunction    = KeyExports     -1
	KeyHash        = KeyFunction    -1
	KeyId          = KeyHash        -1
	KeyIncoming    = KeyId          -1
	KeyIota        = KeyIncoming    -1
	KeyLength      = KeyIota        -1
	KeyLog         = KeyLength      -1
	KeyLogs        = KeyLog         -1
	KeyName        = KeyLogs        -1
	KeyOwner       = KeyName        -1
	KeyPanic       = KeyOwner       -1
	KeyParams      = KeyPanic       -1
	KeyPosts       = KeyParams      -1
	KeyRandom      = KeyPosts       -1
	KeyResults     = KeyRandom      -1
	KeyState       = KeyResults     -1
	KeyTimestamp   = KeyState       -1
	KeyTrace       = KeyTimestamp   -1
	KeyTraceAll    = KeyTrace       -1
	KeyTransfers   = KeyTraceAll    -1
	KeyUtility     = KeyTransfers   -1
	KeyViews       = KeyUtility     -1
	KeyWarning     = KeyViews       -1
	KeyZzzzzzz     = KeyWarning     -1
)
// @formatter:on

var keyMap = map[string]int32{
	"agent":       KeyAgent,
	"amount":      KeyAmount,
	"balances":    KeyBalances,
	"base58":      KeyBase58,
	"caller":      KeyCaller,
	"calls":       KeyCalls,
	"chain":       KeyChain,
	"color":       KeyColor,
	"contract":    KeyContract,
	"data":        KeyData,
	"delay":       KeyDelay,
	"description": KeyDescription,
	"error":       KeyError,
	"exports":     KeyExports,
	"function":    KeyFunction,
	"hash":        KeyHash,
	"id":          KeyId,
	"incoming":    KeyIncoming,
	"iota":        KeyIota,
	"length":      KeyLength,
	"log":         KeyLog,
	"logs":        KeyLogs,
	"name":        KeyName,
	"owner":       KeyOwner,
	"panic":       KeyPanic,
	"params":      KeyParams,
	"posts":       KeyPosts,
	"random":      KeyRandom,
	"results":     KeyResults,
	"state":       KeyState,
	"timestamp":   KeyTimestamp,
	"trace":       KeyTrace,
	"traceAll":    KeyTraceAll,
	"transfers":   KeyTransfers,
	"utility":     KeyUtility,
	"views":       KeyViews,
	"warning":     KeyWarning,
}
