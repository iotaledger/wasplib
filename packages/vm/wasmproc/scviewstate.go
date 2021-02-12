package wasmproc

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
)

// Wraps immutable view state into a mutable KVStore
// with the KVStoreWriter functions implemented as panics
// WaspLib already takes care of the immutability aspect,
// so these panics should never trigger and we can avoid
// a much more drastic refactoring for now
type ScViewState struct {
	ctxView   coretypes.SandboxView
	viewState kv.KVStoreReader
}

func NewScViewState(ctxView coretypes.SandboxView) kv.KVStore {
	return &ScViewState{ctxView: ctxView, viewState: ctxView.State()}
}

func (s ScViewState) Set(key kv.Key, value []byte) {
	s.ctxView.Log().Panicf("ScViewState.Set")
}

func (s ScViewState) Del(key kv.Key) {
	s.ctxView.Log().Panicf("ScViewState.Del")
}

func (s ScViewState) Get(key kv.Key) ([]byte, error) {
	return s.viewState.Get(key)
}

func (s ScViewState) Has(key kv.Key) (bool, error) {
	return s.viewState.Has(key)
}

func (s ScViewState) Iterate(prefix kv.Key, f func(key kv.Key, value []byte) bool) error {
	return s.viewState.Iterate(prefix, f)
}

func (s ScViewState) IterateKeys(prefix kv.Key, f func(key kv.Key) bool) error {
	return s.viewState.IterateKeys(prefix, f)
}

func (s ScViewState) MustGet(key kv.Key) []byte {
	return s.viewState.MustGet(key)
}

func (s ScViewState) MustHas(key kv.Key) bool {
	return s.viewState.MustHas(key)
}

func (s ScViewState) MustIterate(prefix kv.Key, f func(key kv.Key, value []byte) bool) {
	s.viewState.MustIterate(prefix, f)
}

func (s ScViewState) MustIterateKeys(prefix kv.Key, f func(key kv.Key) bool) {
	s.viewState.MustIterateKeys(prefix, f)
}
