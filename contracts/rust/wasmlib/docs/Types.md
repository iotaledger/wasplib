## WasmLib Data Types

WasmLib provides direct support for the following ISCP value data types:

- `Int64` - 64-bit integer value.
- `Bytes` - An arbitrary-length byte array.
- `String` - An UTF-8 encoded string value.
- `Address` - A 33-byte Tangle address.
- `AgentId` - A 37-byte ISCP Agent id.
- `ChainId` - A 33-byte ISCP Chain id.
- `Color` - A 32-byte token color id.
- `ContractId` - A 37-byte ISCP smart contract id.
- `Hash` - A 32-byte hash values.
- `HName` - A 4-byte unsigned integer hash value derived from a name string.
- `RequestId` - A 34-byte transaction request id.

The first three are basic value data types found in all programming languages,
whereas the other types are ISCP-specific value data types. More detailed
explanations about their specific use can be found in
the [documentation of the ISCP](../../../../docs/docs/coretypes.md). WasmLib
provides implementations for each of these value data types. They can all be
serialized into and deserialized from a byte array. Each value data type can
also be used as a key in key/value maps.

WasmLib implements value proxies for each value type. WasmLib also implements a
set of container proxies. There is a map proxy that allows the value types to be
used as key and/or stored value. There are also array proxies for arrays of each
of these value types and for arrays of maps.

Another thing we need to consider is that some data provided by the host is
mutable, whereas other data may be immutable. To facilitate this each proxy type
comes in two flavors that reflect this and make sure the data is only used as
intended. The rule is that an immutable container proxy can only produce
immutable container and value proxies. The contents of these containers can
never be changed through these proxies. Separating these constraints for types
into separate proxy types allows us to use a compiler's static type checking to
enforce these constraints. The ISCP sandbox will also check these constraints at
runtime on the host to guard against client code that bypasses them.

Here is the full matrix of WasmLib types (excluding array proxies):

| ISCP type  | WasmLib type | Mutable proxy type  | Immutable proxy type  |
| ---------- | ------------ | ------------------- | --------------------- |
| Address    | Sc**Address**    | ScMutable**Address**    | ScImmutable**Address**    |
| AgentId    | Sc**AgentId**    | ScMutable**AgentId**    | ScImmutable**AgentId**    |
| Bytes      | *byte array*     | ScMutable**Bytes**      | ScImmutable**Bytes**      |
| ChainId    | Sc**ChainId**    | ScMutable**ChainId**    | ScImmutable**ChainId**    |
| Color      | Sc**Color**      | ScMutable**Color**      | ScImmutable**Color**      |
| ContractId | Sc**ContractId** | ScMutable**ContractId** | ScImmutable**ContractId** |
| HName      | Sc**HName**      | ScMutable**HName**      | ScImmutable**HName**      |
| Hash       | Sc**Hash**       | ScMutable**Hash**       | ScImmutable**Hash**       |
| Int64      | *64-bit integer* | ScMutable**Int64**      | ScImmutable**Int64**      |
| Map        | Sc**Map**        | ScMutable**Map**        | ScImmutable**Map**        |
| RequestId  | Sc**RequestId**  | ScMutable**RequestId**  | ScImmutable**RequestId**  |
| String     | *UTF-8 string*   | ScMutable**String**     | ScImmutable**String**     |

Note how consistent naming makes it easy to remember the type names and how
Bytes, Int64, and String are the odd ones out in that regard as they are
implemented in WasmLib by the closest equivalents in the chosen implementation
programming language.

Here is the full matrix of WasmLib types for array proxies:

| ISCP type  | Mutable array proxy type  | Immutable array proxy type  |
| ---------- | ------------------- | --------------------- |
| Address    | ScMutable**Address**Array    | ScImmutable**Address**Array    |
| AgentId    | ScMutable**AgentId**Array    | ScImmutable**AgentId**Array    |
| Bytes      | ScMutable**Bytes**Array      | ScImmutable**Bytes**Array      |
| ChainId    | ScMutable**ChainId**Array    | ScImmutable**ChainId**Array    |
| Color      | ScMutable**Color**Array      | ScImmutable**Color**Array      |
| ContractId | ScMutable**ContractId**Array | ScImmutable**ContractId**Array |
| HName      | ScMutable**HName**Array      | ScImmutable**HName**Array      |
| Hash       | ScMutable**Hash**Array       | ScImmutable**Hash**Array       |
| Int64      | ScMutable**Int64**Array      | ScImmutable**Int64**Array      |
| Map        | ScMutable**Map**Array        | ScImmutable**Map**Array        |
| RequestId  | ScMutable**RequestId**Array  | ScImmutable**RequestId**Array  |
| String     | ScMutable**String**Array     | ScImmutable**String**Array     |

Again, consistency in naming makes them easy to remember.

In the next section we will show how the WasmLib types are used in smart 
contract code.

Next: [Function Call Context](Context.md)
