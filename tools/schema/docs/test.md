## Testing Smart Contracts

Testing of smart contracts initially happens in the Solo testing environment. This enables
synchronous, deterministic testing of smart contract functionality without the overhead of
having to start nodes, set up a committee, and send transactions over the Tangle. Instead,
we can use Go's built-in test environment in combination with Solo to deploy chains and
smart contracts and simulate transactions.

Solo directly interacts with the ISCP code and uses all the data types that are 
[defined in the ISCP code](https://github.com/iotaledger/wasp/blob/develop/documentation/docs/misc/coretypes.md) directly. But our Wasm smart contracts cannot access these types directly
because they run in a sandboxed environment. Therefore, WasmLib implements its
[own versions](types.md) of these data types and the VM layer acts as a data type
translator between both systems.

The impact of this type transformation used to be that to be able to write tests in the
solo environment the user also needed to know about the ISCP-specific data types and type
conversion functions, and exactly how to properly pass such data in and out of smart
contract function calls. This burdened the user with an unnecessary increased learning
curve.

With the introduction of the schema tool we were able to remove this impedance mismatch
and allow the user to test smart contract functionality in terms of the WasmLib data types
and functions that he is already familiar with. To that end we introduced a new ISCP
function context that specifically works with the Solo testing environment. We aimed to
simplify the testing of smart contracts as much as possible, keeping the full Solo
interface hidden as much as possible, yet available when necessary.

The only concession we still have to make is to the language used. Because Solo only works
in the Go language environment, we have to use the Go language version of the interface
code that the schema tool generates when testing our smart contracts. We feel that this is
not unsurmountable, because WasmLib programming for Rust and Go are practically identical.
They only differ where language idiosyncrasies force differences in syntax or naming
conventions. This hurdle used to be a lot bigger before, when direct programming of Solo
had to be used, and type conversions had to be done manually. Instead, we now get to use
the generated compile-time type-checked interface to our smart contract functions that we
are already familiar with.

Let's look at the simplest way of initializing a smart contract by using the new
`SoloContext` in a test function:

```golang
    func TestDeploy(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    require.NoError(t, ctx.ContractExists(dividend.ScName))
}
```

The first line will automatically create a new chain and upload and deploy the provided
example `dividend` contract to this chain. It returns a `SoloContext` for further use. The
second line verifies the existence of the deployed contract on the chain associated with
the context.

Here is another part of the `dividend` test code, where you can see how we wrap repetitive
calls to smart contract functions that are used in multiple tests:

```golang
func dividendMember(ctx *wasmsolo.SoloContext, agent *wasmsolo.SoloAgent, factor int64) {
    member := dividend.ScFuncs.Member(ctx)
    member.Params.Address().SetValue(agent.ScAddress())
    member.Params.Factor().SetValue(factor)
    member.Func.TransferIotas(1).Post()
}

func dividendDivide(ctx *wasmsolo.SoloContext, amount int64) {
    divide := dividend.ScFuncs.Divide(ctx)
    divide.Func.TransferIotas(amount).Post()
}

func dividendGetFactor(ctx *wasmsolo.SoloContext, member3 *wasmsolo.SoloAgent) int64 {
    getFactor := dividend.ScFuncs.GetFactor(ctx)
    getFactor.Params.Address().SetValue(member3.ScAddress())
    getFactor.Func.Call()
    value := getFactor.Results.Factor().Value()
    return value
}
```

As you can see we pass in the SoloContext and the parameters to the wrapper functions,
then use the context to create a function descriptor for the wrapped function, pass any
parameters through the `Params` proxy, and then either post the function request or call
the function. Any results returned are extracted through the `Results` proxy and returned
by the wrapper.

As you can see there is literally no difference in the way the function interface is used
with the ISCP function context in WasmLib and with the SoloContext. This makes for
seamless testing of smart contracts

In the next section we will go deeper into how the helper member functions of the 
SoloContext are used to simplify tests.

[Example Tests](examples.md)
