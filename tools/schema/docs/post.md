#### Posting Asynchronous Requests

Asynchronous function calls between smart contracts are posted as requests on the Tangle.
They allow you to invoke any smart contract function that is not a View on any smart
contract chain. You will notice that the behavior is very similar to a normal function
call, but instead of using the call() method of the `func` member in the function
descriptor you will now use the post() or post_to_chain() methods. The former posts the
request within the current chain, and the latter takes the chain ID of the desired chain
as parameter.

In addition to the previously discussed transfer_tokens() and of_contract() methods you
can modify the behavior further by providing a delay() in seconds, which enables delayed
execution of the request. Note that this is of particular interest to smart contracts that
need a delayed action, like betting contracts with a timed betting round, or to create
time-lock functionality in a smart contract. Here's how that works:

```rust
...

let eor = ScFuncs::end_of_round(ctx);
eor.func.delay(3600).transfer_iotas(1).post();

...
```

Note that an asynchronous request always needs to send at least 1 token, because it is
posted as a request on the Tangle, and it is not possible to have a request without a
transfer. So if you post to a function that expects tokens you just specify the amount of
tokens required, but if you post to a function that does not expect any tokens then you
still have to provide 1 iota.

Also note that providing a delay() before a call() will result in a run-time error. We do
not know the intention of the user until the actual call() or post() is encountered, so we
cannot check for this at compile-time unless we are willing to accept a lot of extra
overhead. It should not really be a problem because using delay() is rare, anyway, and
using it with call() cannot have been the intention.

The function that posts the request through the function descriptor will immediately
continue execution and does not wait for its completion. Therefore, it is not possible to
directly retrieve the return values from such a call.

If you need some return values you will have to create a mechanism that can do so, for
example by providing a callback chain/contract/function combination as part of the input
parameters of the requested function, so that upon completion that function can
asynchronously post the results to the indicated function. It will require a certain
degree of cooperation between both smart contracts. In the future we will probably be
looking at providing a generic mechanism for this.

In the next section we will look at how we can use the function descriptors when testing
smart contracts with Solo.

Next: [Testing Smart Contracts](test.md)


