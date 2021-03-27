## Smart Contract Initialization

Smart contracts start out with a completely blank state. Sometimes you want to
be able to define initial state, for example if your contract is configurable.
You want to be able to pass this configuration to the contract upon deployment,
so that its state reflects that configuration once the first request comes in.
To support this initialization we allow the smart contract creator to provide an
optional function named `init`.

This 'init' function, when provided, will automatically be called immediately
after the first time the contract has been deployed to the VM. Note that this is
a one-time initialization call, meant to be performed by the contract deployment
mechanism. ISCP will prevent anyone else from calling this function ever again.
So if you need to be able to reconfigure the contract later on, then you will
need to provide a separate configuration function, and guard it from being
accessed by anyone else than properly authorized entities.

To show how creating a smart contract with WasmLib works we will slowly start
fleshing out the smart contract functions of the `dividend` example in this
tutorial. Here is the first part of the Rust code that implements it, which
contains the 'init' function:

```rust
// This example implements 'dividend', a simple smart contract that will
// automatically disperse iota tokens which are sent to the contract to a group
// of member addresses according to predefined division factors. The intent is
// to showcase basic functionality of WasmLib through a minimal implementation
// and not to come up with a complete robust real-world solution.
// Note that we have drawn out constructs that could have been done in a single
// line over multiple statements to be able to properly document step by step
// what is happening in the code.

use wasmlib::*;

use crate::*;

// 'init' is used as a way to initialize a smart contract. It is an optional
// function that will automatically be called upon contract deployment. In this
// case we use it to initialize the 'owner' state variable so that we can later
// use this information to prevent non-owners from calling certain functions.
// The 'init' function takes a single optional parameter:
// - 'owner', which is the agent id of the entity owning the contract.
// When this parameter is omitted the owner will default to the contract creator.
pub fn func_init(ctx: &ScFuncContext) {

    // Log initiation of the 'init' Func in the host log.
    ctx.log("dividend.init");

    // First we set up a default value for the owner in case the optional
    // 'owner' parameter was omitted.
    let mut owner: ScAgentId = ctx.contract_creator();

    // Now it is time to check if parameters were provided to the function.
    // We create an ScImmutableMap proxy to the params map on the host.
    let p: ScImmutableMap = ctx.params();

    // Then we create an ScImmutableAgentId proxy to the 'owner' parameter.
    let param_owner: ScImmutableAgentId = p.get_agent_id(PARAM_OWNER);

    // Now we check if the 'owner' parameter is present in the params map.
    if param_owner.exists() {
        // Yes, it was present, so now we overwrite the default owner with
        // the one specified by the 'owner' parameter.
        owner = param_owner.value();
    }
```

We define an owner variable and allow it to be something other than the default
value of contract creator. It is always a good idea to be flexible enough to be
able to transfer ownership to another entity if necessary. Remember that once a
smart contract is deployed it is no longer possible to change it. Therefore, it
is good practice thinking through those situations that could require change in
advance and allow the contract itself to handle such changes through its state
by providing a proper function interface.

We only define a single owner here. Proper fall-back could require multiple
owners in case the owner entity could disappear, which would allow others to
take over instead of the contract becoming immutable w.r.t. owner functionality.
Again, we cannot stress enough how important it is to think through every aspect
of a smart contract before deployment.

```rust
    // Now that we have sorted out which agent will be the owner of this contract
    // we will save this value in the state storage on the host. First we create
    // an ScMutableMap proxy that refers to the state storage map on the host.
    let state: ScMutableMap = ctx.state();

    // Then we create an ScMutableAgentId proxy to an 'owner' variable in state
    // storage on the host.
    let state_owner: ScMutableAgentId = state.get_agent_id(VAR_OWNER);

    // And then we save the owner value in the 'owner' variable in state storage.
    state_owner.set_value(&owner);

    // Log successful completion of the 'init' Func in the host log.
    ctx.log("dividend.init ok");
}
```

In the next sections we will go deeper into explaining how function parameters
are passed to a smart contract function and how to interact with the state
storage.

Next: [Function Parameters](Params.md)
