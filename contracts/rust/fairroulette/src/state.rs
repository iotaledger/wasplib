// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;
use crate::types::*;

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

pub struct FairRouletteFuncState {
    pub(crate) state_id: i32,
}

impl FairRouletteFuncState {
    pub fn bets(&self) -> ArrayOfMutableBet {
        let arr_id = get_object_id(self.state_id, idx_map(IDX_VAR_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBet { obj_id: arr_id }
    }

    pub fn last_winning_number(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.state_id, idx_map(IDX_VAR_LAST_WINNING_NUMBER))
    }

    pub fn locked_bets(&self) -> ArrayOfMutableBet {
        let arr_id = get_object_id(self.state_id, idx_map(IDX_VAR_LOCKED_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfMutableBet { obj_id: arr_id }
    }

    pub fn play_period(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.state_id, idx_map(IDX_VAR_PLAY_PERIOD))
    }
}

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

pub struct FairRouletteViewState {
    pub(crate) state_id: i32,
}

impl FairRouletteViewState {
    pub fn bets(&self) -> ArrayOfImmutableBet {
        let arr_id = get_object_id(self.state_id, idx_map(IDX_VAR_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBet { obj_id: arr_id }
    }

    pub fn last_winning_number(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.state_id, idx_map(IDX_VAR_LAST_WINNING_NUMBER))
    }

    pub fn locked_bets(&self) -> ArrayOfImmutableBet {
        let arr_id = get_object_id(self.state_id, idx_map(IDX_VAR_LOCKED_BETS), TYPE_ARRAY | TYPE_BYTES);
        ArrayOfImmutableBet { obj_id: arr_id }
    }

    pub fn play_period(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.state_id, idx_map(IDX_VAR_PLAY_PERIOD))
    }
}