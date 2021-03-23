## Function Parameters

Smart contract functions can be invoked in two ways:

- Externally, by sending a smart contract request message via the Tangle.
- Internally, by using the call() or post() methods of the function context.

In both cases it is possible to pass parameters to the smart contract function
that is being invoked. These parameters are presented as a key/value map through
the params() method of the function context. Keys can be any byte array, but as
a convention we will use human-readable string names, which greatly simplifies
debugging.

We will use the 'member' function of the dividend smart contract to highlight
how we can properly interact with these parameters:

```rust
// 'member' is a function that can be used only by the entity that owns the
// 'dividend' smart contract. It can be used to define the group of member
// addresses and dispersal factors one by one prior to sending tokens to the
// smart contract's 'divide' function. The 'member' function takes 2 parameters,
// which are both required:
// - 'address', which is an Address to use as member in the group, and
// - 'factor',  which is an Int64 relative dispersal factor associated with
//              that address
// The 'member' function will save the address/factor combination in its state
// storage and also calculate and store a running sum of all factors so that the
// 'divide' function can simply start using these precalculated values
pub fn func_member(ctx: &ScFuncContext) {

    // Log initiation of the 'member' Func in the host log.
    ctx.log("dividend.member");

    // The 'init' func previously determined which agent is the owner of this
    // contract and stored that value in the 'owner' variable in state storage.
    // So we start out by accessing state storage by creating an ScMutableMap
    // proxy that refers to the state storage map on the host.
    let state: ScMutableMap = ctx.state();

    // Next we create an ScMutableAgentId proxy to the 'owner' variable in state
    // storage.
    let owner: ScMutableAgentId = state.get_agent_id(VAR_OWNER);

    // Only the defined smart contract owner can add members, so we require
    // that the caller's agent id is equal to the stored owner's agent id.
    // Otherwise we panic out with an error message.
    ctx.require(ctx.caller() == owner.value(), "no permission");
```

Note how we use the require() method of the function context to verify that a
condition is true and to specify the error message to be logged in the host log
before panic()-ing out of this function call. You can use the require() method
anywhere you need to assert a condition.

Panicking out of the function call is the only way to signal to the host that an
error has occurred. It will cause the host to roll back any changes that were
made to the state and return any tokens that were transferred as part of the
function call (minus any fees, if fees are required). Smart contracts are not
equipped to deal with unexpected errors and should therefore always abort and
roll back when an error condition occurs to prevent the error to propagate into
the state as invalid data. You will notice that whenever a non-fatal error
condition can occur that *can* be handled by the smart contract correctly we
will provide a method to explicitly test for that condition.

```rust
// Now it is time to check the parameters that were provided to the function.
// First we create an ScImmutableMap proxy to the params map on the host.
let p: ScImmutableMap = ctx.params();

// Create an ScImmutableAddress proxy to the 'address' parameter that is still
// stored in the params map on the host. Note that we use constants defined in
// consts.rs to prevent typos in name strings. This is good practice and will
// save time in the long run.
let param_address: ScImmutableAddress = p.get_address(PARAM_ADDRESS);

// Require that the mandatory 'address' parameter actually exists in the map on
// the host. If it doesn't we panic out with an error message.
ctx.require(param_address.exists(), "missing mandatory address");

// Now that we are sure that the 'address' parameter actually exists we can
// retrieve its actual value into an ScAddress value type.
let address: ScAddress = param_address.value();
```

The first thing we do is create the ScImmutableMap proxy to the map containing
the raw data bytes of the parameters. The host system only stores the raw
key/value bytes that were passed to the call. It is up to us to interpret these
bytes correctly as proper value types. But before we can do that we need to
verify that there actually wre value bytes provided under the specific key we
expect. So we create an ScImmutableAddress value proxy for the key "address"
from the params map proxy. This will allow us to do both.

We use the exists() method of ScImmutableAddress in a require() condition that
verifies that the key "address" actually exists in the key/value map on the
host. If it doesn't exist then we panic() out of the function, because this was
supposed to be a mandatory parameter.

When the key exists we use the value() method of the proxy to retrieve the data
bytes associated with that key from the map on the host and because this is an
ScImmutableAddress proxy it will try to interpret the data bytes as an ScAddress
value type. When that fails, it will automatically panic() out of the function.

Now let's do the same for the mandatory "factor" parameter.

```rust
// Create an ScImmutableInt64 proxy to the 'factor' parameter that is still
// stored in the map on the host. Note how the get_xxx() method defines what
// type of parameter we expect. In this case it's an Int64 parameter.
let param_factor: ScImmutableInt64 = p.get_int64(PARAM_FACTOR);

// Require that the mandatory 'factor' parameter actually exists in the map on
// the host. If it doesn't we panic out with an error message.
ctx.require(param_factor.exists(), "missing mandatory factor");

// Now that we are sure that the 'factor' parameter actually exists we can
// retrieve its actual value into an i64. Note that we use Rust's built-in
// data types when manipulating Int64, String, or Bytes value objects.
let factor: i64 = param_factor.value();

// As an extra requirement we check that the 'factor' parameter value is not
// negative. If it is, we panic out with an error message.
// Note how we use an if expression here. We could have achieved the same in a
// single line by using the require() method instead:
// ctx.require(factor >= 0, "negative factor");
// Using the require() method reduces typing and enhances readability.
if factor < 0 {
    ctx.panic("negative factor");
}
```

This concludes the checking of the parameters and retrieving their values. In
the next section we will show how to access the state storage of the smart
contract.

Next: [Smart Contract State](State.md)