## Function Call Context

With the proxy objects providing the generic access capability to the data on
the host it is now time to introduce the gateway to the host that allows us to
access everything that the host sandbox interface provides. We call this gateway
the function call context, and it is provided as a predefined parameter to the
smart contract function. In fact, we distinguish two separate kinds of smart
contract functions in the ISCP:

- Func, which allows full mutable access to the smart contract state, and always
  results in a state update.
- View, which allows only limited, immutable access to the smart contract state,
  and therefore does not result in a state update. Views are used to query the
  current state of the smart contract.

To support this function type distinction, Func and View each receive a
separate, different function call context, and only the functionality that is
necessary for their implementation is provided through their respective contexts
`ScFuncContext` and `ScViewContext`. `ScViewContext` only provides a limited,
immutable subset of the functionality provided by `ScFuncContext`. Again, the
compiler will prevent the programmer from using the wrong functionality in the
wrong context.

An important part of setting up a smart contract is defining exactly which Funcs
and Views are available and informing the host about them, because it's the host
that will dispatch the function calls to the smart contract code. To that end
the smart contract Wasm code will expose an externally callable function named
_on\_load_ that will be called by the host upon initialization of the smart
contract code. The on_load function must provide the host with the list of Funcs
and Views and specific identifiers that can be used to invoke them. It does this
by creating a special context named _ScExports_ and use that to provide the host
with a set of function, type, name, and identifier for each Func and View that
can be called in the smart contract.

Once the host is asked to call a smart contract function it will do so by
invoking a second externally callable function named _on\_call\_entrypoint_
and passing it the identifier for the smart contract Func or View that need to
be invoked. The client code will then use this identifier to set up the correct
context and call the function. Note that there are no other parameters necessary
because the function can access any other function call parameters through its
context object.

Here is a (simplified) example from the _dividend_ smart contract, written in
Rust, that we will use to showcase the features of WasmLib:

```rust
fn on_load() {
    let exports = ScExports::new();
    exports.add_func("divide", func_divide);
    exports.add_func("member", func_member);
    exports.add_view("getFactor", view_get_factor);
}
```

As you can see this on_load() function first creates the required `ScExports`
context and then proceeds to define two Funcs named "divide" and "member" by
calling the add_func() method of the `ScExports` object and then one View named
"getFactor" by calling its add_view() method. The second parameter to these
methods is the smart contract function associated with the name specified. These
methods will also automatically assign unique identifiers and then send it all
to the host.

In its simplest form this is all that is necessary to initialize the smart
contract. However, it could be that the smart contract needs further
configuration upon initialization with parameters specified by whichever actor
deployed the smart contract. To facilitate such further configuration the smart
contract programmer can optionally provide a Func called "init". The host will
automatically finalize the initialization process by calling the
"init" function when present and passing it any configuration parameters that
were specified.

To finalize this example, here is what the skeleton function implementations for
the above smart contract definition would look like:

```rust
fn func_divide(ctx: &ScFuncContext) {
    ctx.log("Calling divide");
}

fn func_member(ctx: &ScFuncContext) {
    ctx.log("Calling member");
}

fn view_get_factor(ctx: &ScViewContext) {
    ctx.log("Calling getFactor");
}
```

As you can see the functions are each provided with a context object, which is
conventionally named _ctx_. Notice that the two Funcs are passed an
ScFuncContext, whereas the View is passed an ScViewContext. We're also already
showcasing an important feature of the contexts: the log() method. This can be
used to log human-readable text to the host's log output. Logging text is the
only way to add tracing to a smart contract, because it does not have any I/O
capabilities other than what the host provides. There is a second logging
method, called trace(), that can be used to provide extra debug information to
the host's log output, and which can be selectively turned on and off at the
host.

Next: [Function Parameters](Params.md)