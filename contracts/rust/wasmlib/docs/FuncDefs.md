## Function Definitions

Here is once more the schema.json file for the `dividend` example. We will now
focus on its `funcs` and `views` sections. Since they are structured identically
we will only need to explain this once.

```json
{
  "name": "Dividend",
  "description": "Simple dividend smart contract",
  "state": {
    "memberList": "[]Address // array with all the recipients of this dividend",
    "members": "[Address]Int64 // map with all the recipient factors of this dividend",
    "owner": "AgentID // owner of contract, the only one who can call 'member' func",
    "totalFactor": "Int64 // sum of all recipient factors"
  },
  "funcs": {
    "init": {
      "params": {
        "owner": "?AgentID // optional owner of contract, defaults to contract creator"
      }
    },
    "member": {
      "access": "owner // only defined owner of contract can add members",
      "params": {
        "address": "Address // address of dividend recipient",
        "factor": "Int64 // relative division factor"
      }
    },
    "divide": {
    },
    "setOwner": {
      "access": "owner // only defined owner of contract can change owner",
      "params": {
        "owner": "AgentID // new owner of smart contract"
      }
    }
  },
  "views": {
    "getFactor": {
      "params": {
        "address": "Address // address of dividend recipient"
      },
      "results": {
        "factor": "Int64 // relative division factor"
      }
    }
  }
}
```

As you can see each of the `funcs` and `views` sections defines their functions
in the same way. The only resulting difference is in the way the schema tool
generates code for them. Funcs will be able to modify the smart contract state,
while views can only retrieve it.

Functions are defined as named subsections, with the name tag being the name of
the function. Under each function subsection in turn there can be 3 optional
subsections.

* `access` indicates who is allowed to access the function.
* `params` holds the field definitions that describe the function parameters.
* `results` holds the field definitions that describe the function results.

We will now examine each subsection in more detail.

### access

The optional `access` subsection is made of a single definition. It indicates
the agent who is allowed to access the function. When this definition is omitted
anyone is allowed to call the function. When the definition is present it should
be an access identifier, optionally followed by an explanatory comment. Access
identifiers can be one of the following:

* `self`: only the smart contract itself can call this function
* `chain`: only the chain owner can call this function
* `creator`: only the contract creator can call this function
* anything else: the name of an AgentID or []AgentID variable in state storage.
  Only the agent(s) defined there can call this function. When this option is
  used you should also provide functionality that can initialize and/or modify
  this variable. As long as this state variable has not been set, nobody is
  allowed to call this function.

The schema tool will automatically generate code to properly check the access
rights of the agent that called the function before the actual function is
called.

You can see usage examples of the access identifier in the schema.json above,
where the state variable `owner` is used as an access identifier. The `init`
function initializes this state variable, and the `setOwner` function can be
used only by the current owner to set a new owner. Finally, the `member`
function can also only be called by the current owner.

Note that there can be different access identifiers for different functions 
as needed. You can set up as many access identifiers as you like.

### params

The optional `params` subsection contains field definitions for each of the
parameters that a function takes. The layout of the field definitions is
identical to that of the [state](State.md) field definitions, with one addition.
The field type can be prefixed with a question mark which indicates that that
parameter is optional.

The schema tool will automatically generate an immutable structure with member
variables for proxies to each parameter variable in the params map. It will also
generate code to check the presence of each non-optional parameter, and it will
also verify the parameter's data type. This checking is done before the function
is called. The user will be able to immediately start using the parameter proxy
through the structure that is passed to the function.

When this subsection is empty or completely omitted, no structure will be
generated or passed to the function.

For example, here is the structure generated for the immutable params for the
`member` function:

```rust
#[derive(Clone, Copy)]
pub struct ImmutableMemberParams {
    pub(crate) id: i32,
}

impl ImmutableMemberParams {
    pub fn address(&self) -> ScImmutableAddress {
        ScImmutableAddress::new(self.id, idx_map(IDX_PARAM_ADDRESS))
    }

    pub fn factor(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_PARAM_FACTOR))
    }
}
```

Note that the schema tool will also generate a mutable version of the structure,
suitable for providing the parameters when calling this smart contract function.

### results

The optional `results` subsection contains field definitions for each of the
results a function produces. The layout of the field definitions is identical to
that of the [state](State.md) field definitions.

The schema tool will automatically generate a mutable structure with member
variables for proxies to each result variable in the results map. The user will
be able to set the result variables through this structure that is passed to the
function.

When this subsection is empty or completely omitted, no structure will be
generated or passed to the function.

For example, here is the structure generated for the mutable results for the
`getFactor` function:

```rust
#[derive(Clone, Copy)]
pub struct MutableGetFactorResults {
    pub(crate) id: i32,
}

impl MutableGetFactorResults {
    pub fn factor(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_FACTOR))
    }
}
```

Note that the schema tool will also generate an immutable version of the
structure, suitable for accessing the results after calling this smart contract
function.

In the next section we will go deeper into the details of how the schema 
tool generates code to make all this work.

Next: [Function Wrappers](Wrappers.md)
