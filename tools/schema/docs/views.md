## View-Only Functions

View-only functions, or Views for short, are smart contract functions that only allow you
to *retrieve* state information about the smart contract. They have a special, limited
ISCP function context that does not allow access to functionality that could result in
changes to the smart contract state. That means that all access to the state storage will
be through immutable proxies. It also means that they cannot receive or transfer tokens,
because changes to the smart contract account are by definition state changes as well.

Views are allowed to call other views on the same chain, but they cannot call any non-view
smart contract function, nor can they post cross-chain requests.

View functions will always return some data to their caller. It would be silly not to
return data from a View because by definition it cannot have any other side effects that
show up elsewhere.

For demonstration purposes we provided a View function with the `dividend` smart contract,
called 'getFactor':

```Rust
// 'getFactor' is a simple View function. It will retrieve the factor
// associated with the (mandatory) address parameter it was provided with.
pub fn view_get_factor(_ctx: &ScViewContext, f: &GetFactorContext) {

    // Since we are sure that the 'address' parameter actually exists we can
    // retrieve its actual value into an ScAddress value type.
    let address: ScAddress = f.params.address().value();

    // Create an immutable map proxy to the 'members' map in the state storage.
    // Note that for views this is an *immutable* map as opposed to the *mutable*
    // map we can access from the *mutable* state that gets passed to funcs.
    let members: MapAddressToImmutableInt64 = f.state.members();

    // Retrieve the factor associated with the address parameter.
    let factor: i64 = members.get_int64(&address).value();

    // Set the factor in the results map of the function context.
    // The contents of this results map is returned to the caller of the function.
    f.results.factor().set_value(factor);
}
```

Return values are passed to the caller through the predefined results() map associated
with the ISCP function context. Again, this works the same way for normal functions,
although normal functions do not necessarily return values to the caller. The schema tool
will set up a function-specific results structure with proxies to the result fields in
this map.

In the next section we will look at smart contract initialization.

Next: [Smart Contract Initialization](init.md)

