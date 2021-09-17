## Type Definitions

Since we allow nesting of container types it is a bit difficult to create proper
declarations for such nested types. Especially because in a field definition you can only
indicate either a single type, or an array of single type, or a map of single type.

We devised a simple solution to this problem. You can add a `typedefs` section to
schema.json where you can define a single type name for a container type. That way you can
easily create containers that contain such container types. The schema tool will
automatically generate the in-between proxy types necessary to make all of this work.

To keep it at the `betting` smart contract from before, imagine we would want to keep
track of all betting rounds. Since a betting round contains an array of all bets in a
round you could not easily define it if it weren't for typedefs.

Instead, now you add the following to your schema.json:

```json
{
  "typedefs": {
    "BettingRound": "[]Bet // one round of bets"
  },
  "state": {
    "rounds": "[]BettingRound // keep track of all betting rounds"
  }
}
```

The schema tool will generate the following proxies in `typedefs.rs`:

```rust
// @formatter:off

#![allow(dead_code)]

use wasmlib::*;
use wasmlib::host::*;

use crate::types::*;

pub type ImmutableBettingRound = ArrayOfImmutableBet;

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

pub type MutableBettingRound = ArrayOfMutableBet;

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

// @formatter:on
```

Note how ImmutableBettingRound and MutableBettingRound type aliases are created for the
types ArrayOfImmutableBet and ArrayOfMutableBet. These are subsequently used in the state
definition in `state.rs`:

```rust
#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::subtypes::*;
use crate::types::*;

pub struct ArrayOfImmutableBettingRound {
    pub(crate) obj_id: i32,
}

impl ArrayOfImmutableBettingRound {
    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_betting_round(&self, index: i32) -> ImmutableBettingRound {
        let sub_id = get_object_id(self.obj_id, Key32(index), TYPE_ARRAY | TYPE_BYTES)
        ImmutableBettingRound { obj_id: sub_id }
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableBettingState {
    pub(crate) id: i32,
}

impl ImmutableBettingState {
    pub fn owner(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }

    pub fn rounds(&self) -> ArrayOfImmutableBettingRound {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_ROUNDS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBettingRound { obj_id: arr_id }
    }
}

pub struct ArrayOfMutableBettingRound {
    pub(crate) obj_id: i32,
}

impl ArrayOfMutableBettingRound {
    pub fn clear(&self) {
        clear(self.obj_id);
    }

    pub fn length(&self) -> i32 {
        get_length(self.obj_id)
    }

    pub fn get_betting_round(&self, index: i32) -> MutableBettingRound {
        let sub_id = get_object_id(self.obj_id, Key32(index), TYPE_ARRAY | TYPE_BYTES)
        MutableBettingRound { obj_id: sub_id }
    }
}

#[derive(Clone, Copy)]
pub struct MutableBettingState {
    pub(crate) id: i32,
}

impl MutableBettingState {
    pub fn owner(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_STATE_OWNER))
    }

    pub fn rounds(&self) -> ArrayOfMutableBettingRound {
        let arr_id = get_object_id(self.id, idx_map(IDX_STATE_ROUNDS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBettingRound { obj_id: arr_id }
    }
}
```

Notice how the rounds() member function returns a proxy to an array of BettingRound. Which
in turn is an array of Bet. So the desired result has been achieved. And every access step
along the way only allows you to take the path laid out which is checked at compile-time.

In the next section we will explore how the schema tool helps to simplify function
definitions.

Next: [Function Definitions](funcs.md)
