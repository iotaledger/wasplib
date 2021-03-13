## View-Only Functions

'getFactor' is a simple example of a View function. It will retrieve the factor
associated with the (mandatory) address parameter it was provided with from the
smart contract state storage.

```Rust
pub fn view_get_factor(ctx: &ScViewContext) {
    // get mandatory address parameter
    let p = ctx.params();
    let param_address = p.get_address("address");
    ctx.require(param_address.exists(), "missing mandatory address");
    let address = param_address.value();

    // get factor associated with address from state
    let state = ctx.state();
    let members = state.get_map("members");
    let factor = members.get_int64(&address).value();

    // set result value to be returned
    let results = ctx.results();
    results.get_int64("factor").set_value(factor);
}
```

First the function checks and retrieves the mandatory address parameter in the
same way as before (see [Function Parameters](Params.md)).

Then we retrieve the current factor for the provided address in the same way we
did before (see [Smart Contract State](State.md)). The only difference here is
that the state() method of the function context returns an ScImmutableMap proxy,
as does its get_map() method. Note that these are *immutable* maps, as opposed
to the mutable maps we get when we call the state() method on an ScFuncContext.

The factor associated with the address parameter is retrieved through an
ScImmutableInt64 proxy to the value stored in the 'members' map. Again, this
proxy only allows *immutable* access to the value.

Now that we have the factor value we want to return it to the caller. Return
values from a smart contract can be anything and are therefore stored in a
key/value map as usual, with this map being returned to the caller. So we create
an ScMutableMap proxy to the results map on the host that will store the
key/value pairs that we want to return from this View function.

Next we set the value associated with the 'factor' key of the results map to the
factor retrieved from the members map in state storage. Note that it is possible
to put result values in the results map at any time. You don't have to wait
until right before returning from the function. It is the *results* map that is
returned automatically by the system upon returning from the call.

Next: [Calling Functions](Calls.md)