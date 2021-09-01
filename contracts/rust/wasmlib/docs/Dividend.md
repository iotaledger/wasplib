## Dividend Code

We have looked at the most important parts of what the schema tool can 
generate from the schema.json file. So now we will start fleshing out the 
user parts of the code of the `dividend` smart contract. We will start with 
the `init` function:

```rust
// 'init' is used as a way to initialize a smart contract. It is an optional
// function that will automatically be called upon contract deployment. In this
// case we use it to initialize the 'owner' state variable so that we can later
// use this information to prevent non-owners from calling certain functions.
// The 'init' function takes a single optional parameter:
// - 'owner', which is the agent id of the entity owning the contract.
// When this parameter is omitted the owner will default to the contract creator.
pub fn func_init(ctx: &ScFuncContext, f: &InitContext) {
    // The schema tool has already created a proper InitContext for this function that
    // allows us to access call parameters and state storage in a type-safe manner.

    // First we set up a default value for the owner in case the optional
    // 'owner' parameter was omitted.
    let mut owner: ScAgentID = ctx.contract_creator();

    // Now we check if the optional 'owner' parameter is present in the params map.
    if f.params.owner().exists() {
        // Yes, it was present, so now we overwrite the default owner with
        // the one specified by the 'owner' parameter.
        owner = f.params.owner().value();
    }

    // Now that we have sorted out which agent will be the owner of this contract
    // we will save this value in the 'owner' variable in state storage on the host.
    f.state.owner().set_value(&owner);
}
```

Note how every parameter and state variable already exists in the structures 
passed in through the function-specific context `f`. This means there is no 
need for the user at this level to create his own proxy objects by using 
identifier constants and forced usage of specific data types. The schema 
tool has already done all that for us, and now it has become very hard to 
make a mistake because everything is checked at compile time. You will also 
notice that at every turn the code completion functionality of your 
development environment will give you only valid choices. Once you get used 
to this you will find that you need to do very little typing to compose your 
smart contract function code.

Here is the implementation of the `setOwner` function:

```rust
// 'setOwner' is used to change the owner of the smart contract.
// It updates the 'owner' state variable with the provided agent id.
// The 'setOwner' function takes a single mandatory parameter:
// - 'owner', which is the agent id of the entity that will own the contract.
// Only the current owner can change the owner.
pub fn func_set_owner(_ctx: &ScFuncContext, f: &SetOwnerContext) {
    // Note that the schema tool has already dealt with making sure that this function
    // can only be called by the owner and that the required parameter is present.
    // So once we get to this point in the code we can take that as a given.

    // Save the new owner parameter value in the 'owner' variable in state storage.
    f.state.owner().set_value(&f.params.owner().value());
}
```

See how the automatic checking of access rights and presence of required
parameters in the thunk function leaves only a single line of code for the user
to write?

Next, we will start fleshing out the `member` function:

```rust
// 'member' is a function that can only be used by the entity that owns the
// 'dividend' smart contract. It can be used to define the group of member
// addresses and dispersal factors one by one prior to sending tokens to the
// smart contract's 'divide' function. The 'member' function takes 2 parameters,
// which are both required:
// - 'address', which is an Address to use as member in the group, and
// - 'factor',  which is an Int64 relative dispersal factor associated with
//              that address
// The 'member' function will save the address/factor combination in state storage
// and also calculate and store a running sum of all factors so that the 'divide'
// function can simply start using these precalculated values when called.
pub fn func_member(ctx: &ScFuncContext, f: &MemberContext) {
    // Note that the schema tool has already dealt with making sure that this function
    // can only be called by the owner and that the required parameters are present.
    // So once we get to this point in the code we can take that as a given.

    // Since we are sure that the 'factor' parameter actually exists we can
    // retrieve its actual value into an i64. Note that we use Rust's built-in
    // data types when manipulating Int64, String, or Bytes value objects.
    let factor: i64 = f.params.factor().value();

    // As an extra requirement we check that the 'factor' parameter value is not
    // negative. If it is, we panic out with an error message.
    // Note how we avoid an if expression like this one here:
    // if factor < 0 {
    //     ctx.panic("negative factor");
    // }
    // Using the require() method instead reduces typing and enhances readability.
    ctx.require(factor >= 0, "negative factor");

    // Since we are sure that the 'address' parameter actually exists we can
    // retrieve its actual value into an ScAddress value type.
    let address: ScAddress = f.params.address().value();

    // We will store the address/factor combinations in a key/value sub-map of the
    // state storage named 'members'. The schema tool has generated an appropriately
    // type-checked proxy map for us from the schema.json state storage definition.
    // If there is no 'members' map present yet in state storage an empty map will
    // automatically be created on the host.
    let members: MapAddressToMutableInt64 = f.state.members();

    // Now we create an ScMutableInt64 proxy for the value stored in the 'members'
    // map under the key defined by the 'address' parameter we retrieved earlier.
    // Note how the only acceptable key type to this map is an ScAddress.
    let current_factor: ScMutableInt64 = members.get_int64(&address);

    // We check to see if this key/value combination exists in the 'members' map.
    if !current_factor.exists() {
        // If it does not exist yet then we have to add this new address to the
        // 'memberList' array that keeps track of all address keys used in the
        // 'members' map. The schema tool has again created the appropriate type
        // for us already. Here too, if the address array was not present yet it
        // will automatically be created on the host.
        let member_list: ArrayOfMutableAddress = f.state.member_list();

        // Now we will append the new address to the memberList array.
        // First we determine the current length of the array.
        let length: i32 = member_list.length();

        // Next we create an ScMutableAddress proxy to the address value that lives
        // at that index in the memberList array (no value yet, since we're appending).
        let new_address: ScMutableAddress = member_list.get_address(length);

        // And finally we append the new address to the array by telling the proxy
        // to update the value it refers to with the 'address' parameter.
        new_address.set_value(&address);

        // Note that we could have achieved the last 3 lines of code in a single line:
        // member_list.get_address(member_list.length()).set_value(&address);
    }
```

Note how we simply define a nested structure of containers within the state map
by using them as if they already existed. The same thing goes for values in the
containers. You can immediately start using them, and they will default to
all-zero values for fixed size value types, and to zero-length values for
variable sized value types. You will see default values in action in the
fragment below.

```rust
    // Create an ScMutableInt64 proxy named 'totalFactor' for an Int64 value in
    // state storage. Note that we don't care whether this value exists or not,
    // because WasmLib will treat it as if it has the default value of zero.
    let total_factor: ScMutableInt64 = f.state.total_factor();

    // Now we calculate the new running total sum of factors by first getting the
    // current value of 'totalFactor' from the state storage, then subtracting the
    // current value of the factor associated with the 'address' parameter, if any
    // exists. Again, if the associated value doesn't exist, WasmLib will assume it
    // to be zero. Finally we add the factor retrieved from the parameters,
    // resulting in the new totalFactor.
    let new_total_factor: i64 = total_factor.value() - current_factor.value() + factor;

    // Now we store the new totalFactor in the state storage.
    total_factor.set_value(new_total_factor);

    // And we also store the factor from the parameters under the address from the
    // parameters in the state storage that the proxy refers to.
    current_factor.set_value(factor);
}
```

This completes the logic for the `member` function. In the next section we will
look at how to detect incoming token transfers and how to send tokens to Tangle
addresses.

Next: [Token Transfers](Transfers.md)
