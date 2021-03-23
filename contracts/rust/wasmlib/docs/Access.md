## Limiting Access

There is one important thing related to smart contract security that we briefly
touched on but did not really discuss yet, and that is how to limit access to
certain functions. It would not be very secure if anyone could call the
dividend 'member' function to set or change the factor for an address. We want
this function to be a configuration function that is only accessible to the
entity that created the smart contract. Luckily, it is pretty simple to set up
this kind of security, because we can use the caller() method of the function
context to determine which entity invoked the function. By starting the function
with a simple test we can rule out that anyone other than the intended entity
can run the function.

Here is how to check that the caller is the contract creator:

```rust
ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");
```

As you can see we compare the ScAgentId value returned by the caller() method
with the ScAgentId value returned by the contract_creator() method of the
function context. By using require() we make sure that if these two do not match
we will panic out of the function with an error message.

Another often-used ScAgentId value to compare the caller() to is returned by the
contract_id() method of the function context. By converting the contract id to
an agent id using its as_agent_id() method you can make sure that a function was
invoked by the smart contract itself.

```rust
ctx.require(ctx.caller() == ctx.contract_id().as_agent_id(), "no permission");
```

Yet another ScAgentId value that can be used is returned by the chain_owner_id()
method of the context. This returns the owner of the current chain that the
smart contract is running on.

```rust
ctx.require(ctx.caller() == ctx.chain_owner_id(), "no permission");
```

To securely set other ScAgentId values to compare the caller() to, you want to
either use the "init" function that is called upon deploying the smart contract
on the host chain to provide one or more ScAgentId values as parameters to be
stored in the state storage, where they can subsequently be retrieved by any
function that requires limiting access to these entities only. Alternatively you
can provide a function that can only be called by the contract creator that can
be used to set or update these values in state storage. That decision depends
much on the requirements for the smart contract. Here is an example where an
'owner' is defined and stored in state storage:

```rust
let owner: ScAgentId = ctx.state().get_agent_id("owner").value();
ctx.require(ctx.caller() == owner, "no permission");
```

It is very important to thoroughly think access limitation and other security
measures through before deploying a smart contract on a chain.

For completeness here is the 'setOwner' function of the dividend smart contract:
```rust

// 'setOwner' is used to change the owner of the smart contract.
// It updates the 'owner' state variable with the provided agent id.
// The 'setOwner' function takes a single mandatory parameter:
// - 'owner', which is the agent id of the entity that will own the contract.
// Only the current owner can change the owner.
pub fn func_set_owner(ctx: &ScFuncContext) {

    // Log initiation of the 'setOwner' Func in the host log.
    ctx.log("dividend.setOwner");

    // Get a proxy to the 'owner' variable in state storage.
    let state_owner: ScMutableAgentId = ctx.state().get_agent_id(VAR_OWNER);

    // Require the caller to be the current owner.
    ctx.require(ctx.caller() == state_owner.value(), "no permission");

    // Get a proxy to the 'owner' parameter.
    let param_owner: ScImmutableAgentId = ctx.params().get_agent_id(PARAM_OWNER);

    // Require that the 'owner' parameter is mandatory.
    ctx.require(param_owner.exists(), "missing mandatory owner");

    // Save the new owner parameter value in the 'owner' variable in state storage.
    state_owner.set_value(&param_owner.value());

    // Log successful completion of the 'setOwner' Func in the host log.
    ctx.log("dividend.setOwner ok");
}
```

In the next section we will explore how we can have smart contracts invoke or
call other smart contract functions.

Next: [Calling Functions](Calls.md)