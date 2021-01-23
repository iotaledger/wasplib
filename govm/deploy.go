package govm

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address/signaturescheme"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/wasmproc"
	"github.com/iotaledger/wasplib/contracts/dividend"
	"github.com/iotaledger/wasplib/contracts/donatewithfeedback"
	"github.com/iotaledger/wasplib/contracts/dummy"
	"github.com/iotaledger/wasplib/contracts/erc20"
	"github.com/iotaledger/wasplib/contracts/example1"
	"github.com/iotaledger/wasplib/contracts/fairauction"
	"github.com/iotaledger/wasplib/contracts/fairroulette"
	"github.com/iotaledger/wasplib/contracts/hellonewworld"
	"github.com/iotaledger/wasplib/contracts/helloworld"
	"github.com/iotaledger/wasplib/contracts/inccounter"
	"github.com/iotaledger/wasplib/contracts/testcore"
	"github.com/iotaledger/wasplib/contracts/tokenregistry"
)

var ScForGoVM = map[string]func(){
	"dividend":           dividend.OnLoad,
	"donatewithfeedback": donatewithfeedback.OnLoad,
	"dummy":              dummy.OnLoad,
	"erc20":              erc20.OnLoad,
	"example1":           example1.OnLoad,
	"fairauction":        fairauction.OnLoad,
	"fairroulette":       fairroulette.OnLoad,
	"helloworld":         helloworld.OnLoad,
	"hellonewworld":      hellonewworld.OnLoad,
	"inccounter":         inccounter.OnLoad,
	"testcore":           testcore.OnLoad,
	"tokenregistry":      tokenregistry.OnLoad,
}

func DeployGoContract(chain *solo.Chain, sigScheme signaturescheme.SignatureScheme, name string, contractName string) error {
	wasmproc.GoWasmVM = NewGoVM(ScForGoVM)
	hprog, err := chain.UploadWasm(sigScheme, []byte("go:"+contractName))
	if err != nil {
		return err
	}
	return chain.DeployContract(sigScheme, name, hprog)
}
