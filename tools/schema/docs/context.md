## Function Call Context

We saw that proxy objects provide generic access capabilities to the data on the host. It
is now time to introduce the gateway to the host that allows us to access the
functionality that the host sandbox interface provides. We call this gateway the _function
call context_
, and it is provided as a predefined parameter to the smart contract function. In fact, we
distinguish two separate flavors of smart contract functions in the ISCP:

- **Func**, which allows full mutable access to the smart contract state, and always
  results in a state update.
- **View**, which allows only limited, immutable access to the smart contract state, and
  therefore does not result in a state update. Views can be used to efficiently query the
  current state of the smart contract.

To support this function type distinction, Func and View functions each receive a
separate, different function call context. Only the functionality that is necessary for
their implementation can be accessed through their respective contexts,
named `ScFuncContext`
and `ScViewContext`. ScViewContext only provides a limited, immutable subset of the full
functionality provided by ScFuncContext. By having separate context types compile-time
type-checking can be used to enforce their usage constraints.

An important part of setting up a smart contract is defining exactly which Funcs and Views
are available and informing the host about them. It is the host that will have to dispatch
the function calls to the corresponding smart contract code. To that end the smart
contract Wasm code will expose an externally callable function named `on_load` that will
be called by the host upon initial loading of the smart contract code. The `on_load`
function must provide the host with the list of Funcs and Views and specific identifiers
that can be used to invoke them. It uses a special temporary function context named
`ScExports`. That context can be used to provide the host with a function, type, name, and
identifier for each Func and View that can be called in the smart contract.

When the host must call a smart contract function it will do so by invoking a second
externally callable function named `on_call`. The host passes the identifier for the smart
contract Func or View that need to be invoked. The client Wasm code will then use this
identifier to set up the corresponding function context and call the function. Note that
there are no other parameters necessary because the function can subsequently access any
other function-related data through its context object.

Here is a (simplified) example from the `dividend` example smart contract that we will use
to showcase some features of WasmLib:

```rust
fn on_load() {
    let exports = ScExports::new();
    exports.add_func("divide", func_divide);
    exports.add_func("init", func_init);
    exports.add_func("member", func_member);
    exports.add_func("setOwner", func_set_owner);
    exports.add_view("getFactor", view_get_factor);
    exports.add_view("getOwner", view_get_owner);
}
```

As you can see this on_load() function first creates the required ScExports context and
then proceeds to define four Funcs named `divide`, `init`, `member`, and `setOwner`. It
does this by calling the add_func() method of the ScExports context. Next it defines two
Views named `getFactor` and `getOwner` by calling the add_view() method of the ScExports
context. The second parameter to these methods is the actual smart contract function
associated with the name specified. These methods will also automatically assign a unique
identifier to the function and then send everything to the host.

In its simplest form this is all that is necessary to initialize a smart contract. To
finalize this example, here is what the skeleton function implementations for the above
smart contract definition would look like:

```rust
fn func_divide(ctx: &ScFuncContext) {
    ctx.log("Calling dividend.divide");
}

fn func_init(ctx: &ScFuncContext) {
    ctx.log("Calling dividend.init");
}

fn func_member(ctx: &ScFuncContext) {
    ctx.log("Calling dividend.member");
}

fn func_set_owner(ctx: &ScFuncContext) {
    ctx.log("Calling dividend.setOwner");
}

fn view_get_factor(ctx: &ScViewContext) {
    ctx.log("Calling dividend.getFactor");
}

fn view_get_owner(ctx: &ScViewContext) {
    ctx.log("Calling dividend.getOwner");
}
```

As you can see the functions are each provided with a context parameter, which is
conventionally named _ctx_. Notice that the four Funcs are passed an ScFuncContext,
whereas the two Views receive an ScViewContext. We're also already showcasing an
important feature of the contexts: the log() method. This can be used to log
human-readable text to the host's log output. Logging text is the only way to add tracing
to a smart contract, because it does not have any I/O capabilities other than what the
host provides. There is a second logging method, called trace(), that can be used to
provide extra debug information to the host's log output. This output can be selectively
turned on and off at the host.

In the next section we will introduce the `schema` tool that simplifies smart contract
programming a lot.

Next: [Smart Contract Schema Tool](schema.md)
