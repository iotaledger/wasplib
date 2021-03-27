## View-Only Functions

View-only functions, or Views for short are smart contract functions that only
allow you to *retrieve* state information about the smart contract. They have a
special, limited function context that does not allow access to functionality
that could result in changes to the smart contract state. That means that all
access to the state storage will be through immutable proxies. It also means
that they cannot receive or transfer tokens, because changes to the smart
contract account are by definition state changes as well.

Views are allowed to call other views on the same chain, but they cannot call
any non-view smart contract function, nor can they post cross-chain requests.

View functions will always return some data to their caller. It would be silly
not to return data from a View because by definition it cannot have any other
side effects that show up elsewhere.

For demonstration purposes we provided a View function with the dividend smart
contract, called 'getFactor':

```Rust
// 'getFactor' is a simple View function. It will retrieve the factor
// associated with the (mandatory) address parameter it was provided with.
pub fn view_get_factor(ctx: &ScViewContext) {

    // Log initiation of the 'getFactor' Func in the host log.
    ctx.log("dividend.getFactor");

    // Now it is time to check the mandatory parameter.
    // First we create an ScImmutableMap proxy to the params map on the host.
    let p: ScImmutableMap = ctx.params();

    // Create an ScImmutableAddress proxy to the 'address' parameter that is
    // still stored in the map on the host.
    let param_address: ScImmutableAddress = p.get_address(PARAM_ADDRESS);

    // Require that the mandatory 'address' parameter actually exists in the map
    // on the host. If it doesn't we panic out with an error message.
    ctx.require(param_address.exists(), "missing mandatory address");

    // Now that we are sure that the 'address' parameter actually exists we can
    // retrieve its actual value into an ScAddress value type.
    let address: ScAddress = param_address.value();
```

As you can see View function parameters are extracted in the exact same way as
with normal functions (see [Function Parameters](Params.md)).

```rust
    // Now that we have sorted out the parameter we will access the state
    // storage on the host. First we create an ScImmutableMap proxy to the state
    // storage map on the host. Note that this is an *immutable* map, as opposed
    // to the *mutable* map we get when we call the state() method on an
    // ScFuncContext.
    let state: ScImmutableMap = ctx.state();
    
    // Create an ScImmutableMap proxy to the 'members' map in the state storage.
    // Note that again, this is an *immutable* map as opposed to the *mutable*
    // map we get from the *mutable* state map we get through ScFuncContext.
    let members: ScImmutableMap = state.get_map(VAR_MEMBERS);
    
    // Retrieve the factor associated with the address parameter through
    // an ScImmutableInt64 proxy to the value stored in the 'members' map.
    let factor: i64 = members.get_int64(&address).value();
```

Accessing the smart contract state also works pretty much the same as with
normal functions (see [Smart Contract State](State.md)). The only difference is
that it is not possible to modify the state in any way.

```rust
    // Create an ScMutableMap proxy to the map on the host that will store
    // the key/value pairs that we want to return from this View function.
    let results: ScMutableMap = ctx.results();
    
    // Set the value associated with the 'factor' key to the factor we got from
    // the members map through an ScMutableInt64 proxy to the results map.
    results.get_int64(VAR_FACTOR).set_value(factor);

    // Log successful completion of the 'getFactor' Func in the host log.
    ctx.log("dividend.getFactor ok");
}
```

Return values are passed to the caller through the predefined results() map
associated with the function context. Again, this works the same way as with
normal functions, although normal functions do not necessarily always return
values to the caller.

In the next section we will go deeper into how to limit access to smart contract
functions.

Next: [Limiting Access](Access.md)
