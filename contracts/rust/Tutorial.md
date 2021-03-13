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
decided upon two kinds of key/value proxy objects. Arrays and maps. The
underlying ISCP sandbox provides access to its data in the form of key/value
stores that can have arbitrary byte data for both key and value. The proxy
objects channel those in easier to use data types, with the necessary type
conversions hidden within WasmLib, while still keeping the option open to use
arbitrary byte strings for keys and values.

Our initial implementation of WasmLib has been created for the Rust programming
language, because this language had the most advanced and stable support for
generating Wasm code at the time when we started implementing our Wasm VM
environment.

Here is a list of topics this tutorial will cover:

* [WasmLib Overview](wasmlib/docs/Overview.md)
* [Proxy Objects](wasmlib/docs/Proxies.md)
* [Function Call Context](wasmlib/docs/Context.md)
* [Function Parameters](wasmlib/docs/Params.md)
* [Smart Contract State](wasmlib/docs/State.md)
* [Incoming Token Transfers](wasmlib/docs/Incoming.md)
* [Limiting Access](wasmlib/docs/Access.md)
* [View-Only Functions](wasmlib/docs/Views.md)
* [Calling Functions](wasmlib/docs/Views.md)
