## WasmLib for Rust

`WasmLib` allows developers to use Rust to create smart contracts for ISCP that compile into Wasm
and can run directly ISCP-enabled Wasp nodes and on Solo environment.

`WasmLib` treats the programming of smart contracts as simple access to a key/value data and token
storage where smart contract properties, request parameters, token balances and the smart contract
state can be accessed in a universal, consistent way.

The _wasmlib_ folder provides the interface to the VM sandbox provided by the Wasp node through _
ScFuncContext_ and _
ScViewContext_.

The folder also contains a number of example smart contracts that can be used to learn how to use _
ScFuncContext_ and _
ScViewContext_ properly.

