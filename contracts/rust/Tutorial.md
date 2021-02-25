## How to write Smart Contracts for ISCP

The Iota Smart Contracts Protocol (ISCP) provides us with a very flexible way of
programming smart contracts. It does this by providing a sandboxed API that
allows you to interact with the ISCP without any security risks. The actual
implementation of the Virtual Machine (VM) that runs in the sandbox environment
is left to whomever wants to create one. Of course, we are providing an example
implementation of such a VM which allows anyone to get a taste of what it is
like to program a smart contract for the ISCP.

Our particular VM uses WebAssembly (Wasm) as in intermediate language and uses
the open source Wasmtime runtime environment to run the Wasm code. Because Wasm
code runs in its own memory space and cannot access anything outside that memory
by design, Wasm code is ideally suited for secure smart contracts. The Wasm
runtime will provide access to functionality that is needed for the smart
contracts to be able to do their thing, but nothing more. In our case all we do
is provide access to the ISCP sandbox environment.

The ISCP sandbox environment enables the following:

- Access to smart contract meta data
- Access to parameter data for smart contract functions
- Access to the smart contract state data
- A way to return data to the caller of the smart contract function
- Access to tokens owned by the smart contract and ability to move them
- Ability to call other smart contract functions
- Access to logging functionality
- Access to a number of utility functions provided by the host

Our choice of Wasm was guided by the desire to be able to program smart
contracts from any language. Since more and more languages are becoming capable
of generating the intermediate Wasm code this will eventually allow developers
to choose a language they are familiar with. To that end we designed the
interface to the ISCP sandboxed environment as a simple library that enables
access to the ISCP sandbox from within the Wasm environment. This library, for
obvious reasons, has been named WasmLib for now.

So why do we need a library to access the sandbox functionality? Why can't we
call the ISCP sandbox functions directly? The reason for that is same reason
that Wasm is secure. There is no way for the Wasm code to access any memory
outside its own memory space. Therefore, any data that is governed by the ISCP
sandbox has to be copied in and out of that memory space through well-defined
channels in the Wasm runtime. To make this whole process as seamless as possible
the WasmLib interface provides proxy objects to hide the underlying data
transfers between the separate systems.

We tried to keep things as simple and understandable as possible, and therefore
decided upon two kinds of proxy objects. Arrays and maps. The underlying ISCP
sandbox provides access to its data in the form of key/value stores that can
have arbitrary byte data for both key and value. The proxy objects channel those
in easier to use data types, with the type conversions hidden within WasmLib,
while still keeping the option open to use arbitrary byte strings for keys and
values.

Our initial implementation of WasmLib has been created for the Rust programming
language, because this language had the most advanced and stable support for
generating Wasm code at the time when we started implementing our Wasm VM
environment.

### WasmLib for Rust

The implementation of WasmLib for Rust provides direct support for the following
value data types:

- Int64 - An integer value is currently represented as a 64-bit integer.
- Bytes - An arbitrary-length byte array
- String - An UTF-8 encoded string value.
- ScAddress - A 33-byte Tangle address
- ScAgentId - A 37-byte ISCP Agent id
- ScChainId - A 33-byte ISCP Chain id
- ScColor - A 32-byte token color id
- ScContractId - A 37-byte ISCP smart contract id
- ScHash - A 32-byte hash values
- ScHname - A 4-byte unsigned integer hash value derived from a name string
- ScRequestId - A 34-byte transaction request id

The ScXxx data types are ISCP-specific and more detailed explanations will be
provided in those instances where they are used. Each of the ScXxx value data
types has at a minimum the ability to convert itself to a byte array, construct
itself from a byte array, and provide a human-readable representation of its
data, in most cases a base58 encoded string. Each value data type also
implements a trait that allows it to be used as a key with the map proxy object.

Since the smart contract data lives on the host, and we cannot simply copy all
data to the Wasm client because it could be prohibitively large, we also use
proxy objects to access values. Another thing to consider it that some data 
provided by the host is immutable, whereas some may be mutable. To 
facilitate this whe introduce two proxy objects per value type.

