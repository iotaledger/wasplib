## Function Parameters

Smart contract functions can be invoked in two ways:

- Externally, by sending a smart contract request message via the Tangle.
- Internally, by using the call() or post() methods of the function context.

In both cases it is possible to pass parameters to the smart contract function
that is being invoked. These parameters are presented as a key/value map through
the params() method. Keys can be any byte array, but as a convention we will use
human-readable string names, which greatly simplifies debugging. To show how
this all works we will slowly start fleshing out the smart contract functions of
the dividend example.

First, we have the "member" function, which takes two parameters that are both
required to be present. The first parameter is an Address value called
"address", and the second is an Int64 value called "factor". The purpose of
the "member" function is to store this combination in the smart contract state
and keep track of the factors for each unique address, as well as precalculate
a "totalFactor" which is the sum of all separate factors.

Let's start with extracting and checking the parameters:

```rust
let params = ctx.params();

let param_address = params.get_address("address");,
ctx.require(params.address.exists(), "missing mandatory address");
let address = param_address.value();

let param_factor = params.get_int64("factor");
ctx.require(params.factor.exists(), "missing mandatory factor");
let factor = param_factor.value();
ctx.require(factor > = 0, "invalid factor");
```

The first thing we do is access the ScImmutableMap proxy object to the
parameters in raw form through the params() method of the function context.
Next, we create an ScImmutableAddress proxy object by asking the params proxy
map to interpret the value for key "address" as an Address. Then we use the
exists() method of ScImmutableAddress to verify that params_address actually
exists in the key/value map on the host. Next, we use the require() method of
the context to verify that this is so and otherwise log an error message in the
host log and panic() out of this call. You can use the require() method anywhere
you need to be sure of a condition. And finally we retrieve the actual value of
the "address" parameter as an ScAddress by calling the value() method of the
proxy object.

Note that panicking out of the function call is the only way to signal that an
error has occurred. It will cause the host to roll back any changes that were
made to the state and return any tokens that were transferred as part of the
function call (minus any fees, if fees are required). Smart contracts are not
equipped to deal with unexpected errors and should therefore always abort and
roll back when an error condition occurs to prevent the error to propagate into
the state as invalid data. You will notice that whenever a special non-fatal
error condition can occur that could be handled by the smart contract correctly
we will provide a method to explicitly test for that condition.

In the second part of the above code excerpt we do the same thing for the
"factor" parameter that we did for the "address" parameter, except that in this
case we expect it to be an Int64 value instead of an Address. And we also
double-check that the value is not negative in another call to the require()
method of the context.

Next: [Smart Contract State](State.md)