#### Calling Functions

Synchronous function calls between smart contracts act very similar to how normal function
calls work in any programming language, but with a slight twist. With normal function
calls you share all the memory that you can access with every function that you call.
However, when calling a smart contract function you can only access the memory assigned to
that specific smart contract. Remember, each smart contract runs in its own sandbox
environment. Therefore, the only way to share data between smart contracts that call each
other is through function parameters and return values.

Synchronous calls can only be made between contracts that are running on the same contract
chain. The ISCP host knows about all the contracts it is running on a chain, and therefore
is able to dispatch the call to the correct contract function. The function descriptor is
used to specify the call parameters (if any) through its `params` proxy, and to invoke the
function through its `func` interface.

In addition, when the function that is called is not a View it is possible to pass tokens
to the function call through this interface. Note that the only way to call a function and
properly pass tokens to it _within the same contract_ is through the function descriptor,
because otherwise the incoming()
function will not register any incoming tokens.

Then the call is made the calling function will be paused and wait for the called function
to complete. After completion, it may access returned values (if any) through the results
proxy of the function descriptor.

When calling from a View function, it is only possible to call other View functions. The
ScFuncs interface enforces this at compile-time through the ISCP function context that
needs to be passed to the member function that creates the function descriptor.

Here's how a smart contract would tell a `dividend` contract on the same chain to divide
the 1000 tokens it passes to the function:

```rust
...

let div = dividend::ScFuncs::divide(ctx);
div.func.transfer_iotas(1000).call();

...
```

And here is how a smart contract would ask a `dividend` contract on the same chain to
return the dispersion factor for a specific address:

```rust
...

let gf = dividend::ScFuncs::get_factor(ctx);
gf.params.address().set_value( & address);
gf.func.call();
let factor = gf.results.factor().value();

...
```

You see how we first create a function descriptor for the desired function, then use
the `params` proxy in the function descriptor to set any required parameters, then direct
the `func` member of the function descriptor to call the associated function, and finally
we use the `results` proxy in the function descriptor to retrieve any results we are
interested in.

Note that the function descriptors assume that the function to be called is associated
with the default Hname of the contract, in this case ScHname::new("dividend"). If you
deployed the contract that contains the function you want to call under a different name
then you would have to provide its associated Hname to the `func` member through the
of_contract() member function like this:

```rust
...

let altContract = ScHname::new("alternateName");
let div = dividend::ScFuncs::divide(ctx);
div.func.of_contract(altContract).transfer_iotas(1000).call();

...
```

In the next section we will look at how we can request smart contract functions in a
different chain to be executed in a similar way.

Next: [Posting Asynchronous Requests](post.md)

