## WaspLib for Rust/Go

```
NOTE 1: This library allows us to compare identical Rust/Go/Java smart contracts
        to highlight the simplicity and consistency of our programming model
        along different programming languages.
      
NOTE 2: We currently use TinyGo to compile the Go smart contracts into
        functioning Wasm code. The resulting Wasm files are still about double
        the size of their Rust equivalent, and one must avoid a few unsupported
        packages for now. Having said that, all the smart contracts provided
        work 100% identical to their Rust equivalents.
```

`WaspLib` allows developers to create Rust or Go smart contracts that compile
into Wasm. The interface provided by `WaspLib`
hides the underlying complexities of the Iota Smart Contract Protocol (ISCP) as
implemented by Iota's ISCP-enabled Wasp nodes.
`WaspLib` treats the programming of smart contracts as simple access to a
key/value data storage where smart contract properties, request parameters, and
the smart contract state can be accessed in a universal, consistent way.

The _wasplib_ folder provides the interface to the ISCP through _ScFuncContext_
and _ScViewContext_.

The _contracts_ folder contains a number of example smart contracts that can be
used to learn how to use _ScFuncContext_ and _ScViewContext_ properly. For more
information on how to go about creating your own smart contracts see the
README.md in this folder.

### Prerequisites for using TinyGo

* Install Go version 1.15.8. Do *NOT* install Go version 1.16, because it is not
  yet supported by TinyGo.
* Install the latest TinyGo. You can find instructions on the
  [TinyGo website](https://tinygo.org/getting-started/).
  