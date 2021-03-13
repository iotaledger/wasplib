## WasmLib Overview

WasmLib provides direct support for the following value data types:

- `Int64` - We currently only directly support 64-bit integer values.
- `Bytes` - An arbitrary-length byte array.
- `String` - An UTF-8 encoded string value.
- `Address` - A 33-byte Tangle address.
- `AgentId` - A 37-byte ISCP Agent id.
- `ChainId` - A 33-byte ISCP Chain id.
- `Color` - A 32-byte token color id.
- `ContractId` - A 37-byte ISCP smart contract id.
- `Hash` - A 32-byte hash values.
- `Hname` - A 4-byte unsigned integer hash value derived from a name string.
- `RequestId` - A 34-byte transaction request id.

The first three are basic value data types found in all programming languages,
whereas the other types are ISCP-specific value data types. More detailed
explanations about their specific use can be found in the [documentation of the
ISCP](../../../../docs/docs/coretypes.md). Each of these value data types has the ability to serialize into and
deserialize from a byte array. Each value data type can also be used as a key to
our key/value proxy objects.

Since the smart contract data lives on the host, and we cannot simply copy all
data to the Wasm client because it could be prohibitively large, we also use
proxy objects to access values. Another thing we need to consider is that some
data provided by the host is immutable, whereas some may be mutable. To
facilitate this whe introduce a number of proxy objects for each of the value
types.

Next: [Proxy Objects](Proxies.md)
