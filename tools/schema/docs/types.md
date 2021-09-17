## WasmLib Data Types

When creating smart contracts you will want to manipulate some data. WasmLib provides
direct support for the following value data types:

- `Int16` - 16-bit signed integer value.
- `Int32` - 32-bit signed integer value.
- `Int64` - 64-bit signed integer value.
- `Bytes` - An arbitrary-length byte array.
- `String` - An UTF-8 encoded string value.

&nbsp;

- `Address` - A 33-byte Tangle address.
- `AgentID` - A 37-byte ISCP Agent ID.
- `ChainID` - A 33-byte ISCP Chain ID.
- `Color` - A 32-byte token color ID.
- `ContractID` - A 37-byte ISCP smart contract ID.
- `Hash` - A 32-byte hash value.
- `Hname` - A 4-byte unsigned integer hash value derived from a name string.
- `RequestID` - A 34-byte transaction request ID.

The first group consists of the basic value data types that are found in all programming
languages, whereas the second group consists of WasmLib versions of ISCP-specific value
data types. More detailed explanations about their specific uses can be found in the
[documentation of the ISCP](https://github.com/iotaledger/wasp/blob/develop/documentation/docs/misc/coretypes.md)
. WasmLib provides its own implementations for each of the ISCP value data types. They can
all be serialized into and deserialized from a byte array. Each value data type can also
be used as a key in key/value maps.

WasmLib implements value [proxies](proxies.md) for each value type. WasmLib also
implements a set of container proxies. There is a map proxy that allows the value types to
be used as key and/or stored value. There are also array proxies for arrays of each of
these value types and for arrays of maps.

Another thing we need to consider is that some data provided by the host is mutable,
whereas other data may be immutable. To facilitate this distinction, each proxy type comes
in two flavors that reflect this and makes sure the data can only be used as intended. The
rule is that from an immutable container proxy you can only derive immutable container and
value proxies. The referenced data can never be changed through immutable proxies.
Separating these constraints for types into separate proxy types allows us to use
compile-time type-checking to enforce these constraints. The ISCP sandbox will also check
these constraints at runtime on the host to guard against client code that tries to bypass
them.

Here is the full matrix of WasmLib types (excluding array proxies):

| ISCP type  | WasmLib type     | Mutable proxy           | Immutable proxy           |
| ---------- | ---------------- | ----------------------- | ------------------------- |
| Bytes      | *byte array*     | ScMutable**Bytes**      | ScImmutable**Bytes**      |
| Int16      | *16-bit integer* | ScMutable**Int16**      | ScImmutable**Int16**      |
| Int32      | *32-bit integer* | ScMutable**Int32**      | ScImmutable**Int32**      |
| Int64      | *64-bit integer* | ScMutable**Int64**      | ScImmutable**Int64**      |
| String     | *UTF-8 string*   | ScMutable**String**     | ScImmutable**String**     |
|            |                  |                         |                           |
| Address    | Sc**Address**    | ScMutable**Address**    | ScImmutable**Address**    |
| AgentId    | Sc**AgentId**    | ScMutable**AgentId**    | ScImmutable**AgentId**    |
| ChainId    | Sc**ChainId**    | ScMutable**ChainId**    | ScImmutable**ChainId**    |
| Color      | Sc**Color**      | ScMutable**Color**      | ScImmutable**Color**      |
| ContractId | Sc**ContractId** | ScMutable**ContractId** | ScImmutable**ContractId** |
| HName      | Sc**HName**      | ScMutable**HName**      | ScImmutable**HName**      |
| Hash       | Sc**Hash**       | ScMutable**Hash**       | ScImmutable**Hash**       |
| Map        | Sc**Map**        | ScMutable**Map**        | ScImmutable**Map**        |
| RequestId  | Sc**RequestId**  | ScMutable**RequestId**  | ScImmutable**RequestId**  |

Note how consistent naming makes it easy to remember the type names and how Bytes, Int16,
Int32, Int64, and String are the odd ones out. They are implemented in WasmLib by the
closest equivalents in the chosen implementation programming language.

Here is the full matrix of WasmLib types for array proxies:

| ISCP type  | Mutable array proxy          | Immutable array proxy          |
| ---------- | ---------------------------- | ------------------------------ |
| Bytes      | ScMutable**Bytes**Array      | ScImmutable**Bytes**Array      |
| Int16      | ScMutable**Int16**Array      | ScImmutable**Int16**Array      |
| Int32      | ScMutable**Int32**Array      | ScImmutable**Int32**Array      |
| Int64      | ScMutable**Int64**Array      | ScImmutable**Int64**Array      |
| String     | ScMutable**String**Array     | ScImmutable**String**Array     |
|            |                              |                                |
| Address    | ScMutable**Address**Array    | ScImmutable**Address**Array    |
| AgentId    | ScMutable**AgentId**Array    | ScImmutable**AgentId**Array    |
| ChainId    | ScMutable**ChainId**Array    | ScImmutable**ChainId**Array    |
| Color      | ScMutable**Color**Array      | ScImmutable**Color**Array      |
| ContractId | ScMutable**ContractId**Array | ScImmutable**ContractId**Array |
| HName      | ScMutable**HName**Array      | ScImmutable**HName**Array      |
| Hash       | ScMutable**Hash**Array       | ScImmutable**Hash**Array       |
| Map        | ScMutable**Map**Array        | ScImmutable**Map**Array        |
| RequestId  | ScMutable**RequestId**Array  | ScImmutable**RequestId**Array  |

Again, consistency in naming makes them easy to remember.

In the next section we will show how the WasmLib types can be used in smart contract code.

Next: [Function Call Context](context.md)
