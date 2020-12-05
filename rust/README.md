## WaspLib for Rust

`WaspLib` allows developers to create Rust smart contracts that compile into Wasm. The interface provided by `WaspLib`
hides the underlying complexities of the Iota Smart Contract Protocol (ISCP) as implemented by Iota's ISCP-enabled Wasp
nodes.
`WaspLib` treats the programming of smart contracts as simple access to a key/value data storage where smart contract
properties, request parameters, and the smart contract state can be accessed in a universal, consistent way.

The _wasplib_ folder provides the interface to the ISCP through _ScContext_.

The _contracts_ folder contains a number of example smart contracts that can be used to learn how to use _ScContext_
properly. For more information on how to go about creating your own smart contracts see the README.md in this folder.

