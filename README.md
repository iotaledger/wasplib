## WaspLib for Go

`NOTE: The Go version is mostly for educational purposes until we have a better
Go to Wasm compiler. It allows us to compare identical Rust/Go/Java smart 
contracts to highlight the simplicity and consistency along different programming
languages.`

`NOTE 2: Due to the limited Wasm capabilities of Go we currently use TinyGo to
compile the Go smart contracts into functioning Wasm code. The resulting Wasm
files are still much larger than their Rust equivalent, and one needs to take
care to avoid certain unsupported packages for now. Having said that, all the
smart contracts provided work 100% identical to their Rust equivalents.`

`WaspLib` allows developers to create Go smart contracts that compile into Wasm.
The interface provided by `WaspLib` hides the underlying complexities of the 
Iota Smart Contract Protocol (ISCP) as implemented by Iota's ISCP-enabled Wasp nodes. 
`WaspLib` treats the programming of smart contracts as simple access to a key/value
data storage where smart contract properties, request parameters, and the smart
contract state can be accessed in a universal, consistent way.

The _wasplib_ folder provides the interface to the ISCP through _ScContext_.

The _contracts_ folder contains a number of example smart contracts that can 
be used to learn how to use _ScContext_ properly. For more information on how
to go about creating your own smart contracts see the README.md in this folder.

