## Calling Functions

There are 2 ways of calling other smart contract functions from within a smart
contract function. Each method has certain properties that have their specific
uses. The main distinction is the following:

- Synchronous calls. These are invoked through the `call()` method of the
  function context.
- Asynchronous calls. These are invoked through the `post()` method of the
  function context.

We will now go deeper into the aspects of these invocation methods.

#### Synchronous Function Calls

Synchronous function calls between smart contracts act very similar to how
normal function calls work in any programming language, but with a twist. With
normal function calls you share all the memory that you can access with every
function that you call. However, when calling a smart contract function you can
only access the memory assigned to that specific smart contract. Remember, each
smart contract runs in its own sandbox environment. Therefore, the only way to
share data between smart contracts that call each other is through function
parameters and return values.

Synchronous calls can only be made between contracts that are running on the
same contract chain. The ISCP host knows about all the contracts it is running
on a chain, and therefore is able to dispatch the call to the correct contract
function. The `call()` method of the function context is used to indicate which
function of which contract to invoke, and to specify the call parameters. In
addition, when the function that is called is not a View it is possible to pass
tokens to the function call if necessary.

The calling function will be paused and wait for the called function to complete
successfully and can then access any returned values through the ScImmutableMap
returned by the call() method.

When calling from a View function, it is only possible to call other View
functions. In fact the call() method of the ScViewContext does not even provide
you with the option to pass tokens to the called function since only Funcs can
process incoming tokens.

We also provide a short-cut method `call_self()` that assumes the function you
are calling is part of the same smart contract as the current function.

#### Asynchronous Function Calls

Asynchronous function calls between smart contracts are posted as requests on
the Tangle. They allow you to invoke any smart contract function on any smart
contract chain. The `post()` method of the function context is used indicate
which function of which contract on which chain to invoke, and to specify the
call parameters. In addition, it is possible to pass tokens to the function call
if necessary. You can also provide a delay, which enables timed execution of
contract functions. In contrast to synchronous calls it is not possible to
post() an asynchronous call from a View. The reason is that the posted request
will become part of the state update message on the Tangle, and therefore
modifies the contract state.

The calling function will continue execution after posting the call and does not
wait for its completion. Therefore, it is not possible to return values from
such a call.

As with the call() method we also provide a short-cut method `post_self()` that
assumes the function you are calling is part of the same smart contract as the
current function.

#### Final word

This concludes this tutorial. Please see the documentation provided with the
source code of WasmLib for a complete list of all functionality that WasmLib
offers through the ISCP sandbox API.

Back to top: [ISCP Tutorial](../../Tutorial.md)
