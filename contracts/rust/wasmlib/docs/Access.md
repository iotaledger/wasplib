## Limiting Access

There is one important thing we did not discuss yet, and that is how to limit
access to certain functions. It would not be very secure if anyone could call
the "member" function to set the factor for an address. We want this to be a
configuration function that is only accessible to the entity that created the
smart contract. Luckily, it is pretty simple to set up this kind of security,
because we can use the caller() method of the function context to determine what
entity invoked the function. By starting the function with a simple test we can
rule out that anyone other than the intended entity can run the function. In the
case of the "member" function this is achieved with a single line of code:

```rust
ctx.require(ctx.caller() == ctx.contract_creator(), "no permission");
```

As you can see we compare the ScAgentId value returned by the caller() method
with the ScAgentId value returned by the contract_creator() method of the
functions context. By using require() we make sure that if these two do not
match we will panic out of the function with an error message.

Another often-used ScAgentId value to compare the caller() to is returned by the
contract_id() method of the function context. This value can be used to make
sure that a function was invoked by the smart contract itself.

Yet another ScAgentId value that can be used is returned by the chain_owner_id()
method of the context. This returns the owner of the current chain the smart
contract is running on.

To securely set other ScAgentId values to compare the caller() to, you want to
either use the "init" function that is called upon loading the smart contract on
the host chain to provide one or more ScAgentId values as parameters to be
stored in the state storage, where they can subsequently be retrieved by any
function that requires limiting access to these entities only. Alternatively you
can provide a function that can only be called by the contract creator that can
be used to set or update these values in state storage. That decision depends
much on the requirements for the smart contract.

It is important to think these security measures through before deploying a
smart contract on a chain.

Next: [View-Only Functions](Views.md)