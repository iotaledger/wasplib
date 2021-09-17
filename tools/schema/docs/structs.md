## Structured Data Types

The schema tool allows you to define your own structured data types that are composed of
the predefined WasmLib value data types. The tool will generate a struct with named fields
according to the definition in schema.json, and also generates code to serialize and
deserialize the structure to a byte array, so that it can be saved as a single unit of
data, for example in state storage.

You can use such structs directly as a type in state storage definitions and the schema
tool will automatically generate the proxy code to access it properly.

For example, let's say you are creating a `betting` smart contract. Then you would want to
store information for each bet. The Bet structure could consist of the bet amount and time
of the bet, the number of the item that was bet on, and the agent ID of the one who placed
the bet. And you would keep track of all bets in state storage in an array of Bet structs.
You would insert the following into schema.json:

```json
{
  "structs": {
    "Bet": {
      "amount": "Int64 // bet amount",
      "time": "Int64 // timestamp of this bet",
      "number": "Int32 // number of item we bet on",
      "better": "AgentID // Who placed this bet"
    }
  },
  "state": {
    "bets": "[]Bet // all bets made in this round"
  }
}
```

The schema tool will generate `types.rs` which contains the following code for the Bet
struct:

```rust
// @formatter:off

#![allow(dead_code)]

use wasmlib::*;
use wasmlib::host::*;

pub struct Bet {
    pub amount: i64,       // bet amount
    pub better: ScAgentID, // Who placed this bet
    pub number: i32,       // number of item we bet on
    pub time:   i64,       // timestamp of this bet
}

impl Bet {
    pub fn from_bytes(bytes: &[u8]) -> Bet {
        let mut decode = BytesDecoder::new(bytes);
        Bet {
            amount: decode.int64(),
            better: decode.agent_id(),
            number: decode.int32(),
            time: decode.int64(),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut encode = BytesEncoder::new();
        encode.int64(self.amount);
        encode.agent_id(&self.better);
        encode.int32(self.number);
        encode.int64(self.time);
        return encode.data();
    }
}

pub struct ImmutableBet {
    pub(crate) obj_id: i32,
    pub(crate) key_id: Key32,
}

impl ImmutableBet {
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    pub fn value(&self) -> Bet {
        Bet::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_BYTES))
    }
}

pub struct MutableBet {
    pub(crate) obj_id: i32,
    pub(crate) key_id: Key32,
}

impl MutableBet {
    pub fn exists(&self) -> bool {
        exists(self.obj_id, self.key_id, TYPE_BYTES)
    }

    pub fn set_value(&self, value: &Bet) {
        set_bytes(self.obj_id, self.key_id, TYPE_BYTES, &value.to_bytes());
    }

    pub fn value(&self) -> Bet {
        Bet::from_bytes(&get_bytes(self.obj_id, self.key_id, TYPE_BYTES))
    }
}

// @formatter:on
```

Notice how the generated ImmutableBet and MutableBet proxies use the from_bytes() and
to_bytes() (de)serialization code to automatically transform byte arrays into Bet structs.

The generated code in `state.rs` that implements the state interface is shown here:

```rust
#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::types::*;

pub struct ArrayOfImmutableBet {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableBet {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bet(&self, index: i32) -> ImmutableBet {
        ImmutableBet { obj_id: self.obj_id, key_id: Key32(index) }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableBettingState {
    pub(crate) id: i32,
}

impl ImmutableBettingState {
    pub fn bets(&self) -> ArrayOfImmutableBet {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBet { obj_id: arr_id }
    }

    pub fn owner(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }
}

pub struct ArrayOfMutableBet {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableBet {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_bet(&self, index: i32) -> MutableBet {
        MutableBet { obj_id: self.obj_id, key_id: Key32(index) }
    }
}

#[derive(Clone, Copy)]
pub struct MutableBettingState {
    pub(crate) id: i32,
}

impl MutableBettingState {
    pub fn bets(&self) -> ArrayOfMutableBet {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBet { obj_id: arr_id }
    }

    pub fn owner(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }
}
```

The end result is an ImmutableBettingState and MutableBettingState structure that can
directly interface to the state of the betting contract.

In the next section we will look at how to make even more complex type definitions.

Next: [Type Definitions](typedefs.md)
