## Example Tests

We saw in the previous section how you can call() or post() function requests. We also
created a few wrapper functions to simplify calling these functions even further. Now we
will look at how to use the SoloContext to create full-blown tests for the
`dividend` example smart contract.

Let's start with a simple test. We're going to use the `member` function to add a valid
new member/factor combination to the member group.

```golang
func TestAddMemberOk(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member1 := ctx.NewSoloAgent()
    dividendMember(ctx, member1, 100)
    require.NoError(t, ctx.Err)
}
```

The above test first deploys the `dividend` smart contract to a new chain and returns a
SoloContext `ctx`. Then it uses ctx to create a new SoloAgent. A SoloAgent is an actor
with its own Tangle address, which contains solo.Saldo tokens. The SoloAgent can be used
whenever an address or agent ID needs to be provided, it can be used to sign a token
transfer from its address, and can be used to inspect the balance of tokens on the
address.

In this case we simply create `member`, and pass it to the `member` function, which will
receive the addres of member1 and a dispersal factor of 100. Finally we check if ctx has
received an error during the execution of the call. Remember that the only way to pass an
error from a WasmLib function call is through a panic() call. The code that handles the
actual call will intercept any panic() that was raised and return an error in that case.
The SoloContext saves this error for later inspection, because the function descriptor
used in the call itself has no way of passing back this error.

The next two example tests each call the same `member` function in the exact same way, but
in both cases one required parameter is omitted. The idea is to test that the function
properly panics by checking the ctx.Err value is not nil after the call. Finally, the
error message text is checked to contain the correct error message.

```golang
func TestAddMemberFailMissingAddress(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member := dividend.ScFuncs.Member(ctx)
    member.Params.Factor().SetValue(100)
    member.Func.TransferIotas(1).Post()
    require.Error(t, ctx.Err)
    require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory address"))
}

func TestAddMemberFailMissingFactor(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member1 := ctx.NewSoloAgent()
    member := dividend.ScFuncs.Member(ctx)
    member.Params.Address().SetValue(member1.ScAddress())
    member.Func.TransferIotas(1).Post()
    require.Error(t, ctx.Err)
    require.True(t, strings.HasSuffix(ctx.Err.Error(), "missing mandatory factor"))
}
```

Notice how each test has to set up the chain/contract/context from scratch. We will often
use a specific setupTest() function to do all setup work that is shared by many tests.

Also notice how we cannot use the `dividendMember` wrapper function in these two tests
because of the missing required function parameters. So we have copy/pasted the code and
removed the Params initialization we wanted to be missing.

Now let's see a more complex example:

```golang
func TestDivide1Member(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member1 := ctx.NewSoloAgent()
    dividendMember(ctx, member1, 100)
    require.NoError(t, ctx.Err)
    
    require.EqualValues(t, 1, ctx.Balance(nil))
    
    dividendDivide(ctx, 99)
    require.NoError(t, ctx.Err)
    
    // 99 from divide() + 1 from the member() call
    require.EqualValues(t, solo.Saldo+100, member1.Balance())
    require.EqualValues(t, 0, ctx.Balance(nil))
}
```

The first half of the code is identical to our first test above. We set up the test,
create an agent, set the factor for that agent to 100, and verify that no error occurred.
Then in the next line we verify that the smart contract associated with ctx now holds a
balance of 1 iota. This is the token that was transferred as part of the Post() request
inside the dividendMember() function.

Next we transfer 99 iotas as part of thePost() request to the `divide` function. We
subsequently check that no error has occurred. Finally, we expect the balance of member1
address to have increased by the total of 100 tokens that were stored in the
`dividend` smart contract account, as 100/100th of the tokens should have been sent to
member1. And the contract account should end up empty.

Now let's skip to the most complex test of all:

```golang
func TestDivide3Members(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member1 := ctx.NewSoloAgent()
    dividendMember(ctx, member1, 25)
    require.NoError(t, ctx.Err)
    
    member2 := ctx.NewSoloAgent()
    dividendMember(ctx, member2, 50)
    require.NoError(t, ctx.Err)
    
    member3 := ctx.NewSoloAgent()
    dividendMember(ctx, member3, 75)
    require.NoError(t, ctx.Err)
    
    require.EqualValues(t, 3, ctx.Balance(nil))
    
    dividendDivide(ctx, 97)
    require.NoError(t, ctx.Err)
    
    // 97 from divide() + 3 from the member() calls
    require.EqualValues(t, solo.Saldo+16, member1.Balance())
    require.EqualValues(t, solo.Saldo+33, member2.Balance())
    require.EqualValues(t, solo.Saldo+50, member3.Balance())
    // 1 remaining due to fractions
    require.EqualValues(t, 1, ctx.Balance(nil))
}
```

This function creates 3 agents, and associates factors of 25, 50, and 75, respectively, to
them. Since that required 3 `member` requests, the contract account should now contain 3
iotas. Next the `divide` function is called, with 97 iotas passed to it, for a total of
100 into the contract account.

After this we verify that each agent has been distributed tokens according to its relative
factor. Those factors are 25/150th, 50/150th, and 75/150th, respectively. Note that we
cannot transfer fractional tokens, so we truncate to the nearest integer and ultimately
are left with 1 iota in the contract account. This 1 iota will be part of the dispersal
amount when the next `divide` call request is executed.

We can test this behavior by adding extra calls to `divide` at the end of this test like
this:

```golang
dividendDivide(ctx, 100)
require.NoError(t, ctx.Err)

// 100 from divide() + 1 remaining
require.EqualValues(t, solo.Saldo+16+16, member1.Balance())
require.EqualValues(t, solo.Saldo+33+33, member2.Balance())
require.EqualValues(t, solo.Saldo+50+50, member3.Balance())
// now we have 2 remaining due to fractions
require.EqualValues(t, 2, ctx.Balance(nil))

dividendDivide(ctx, 100)
require.NoError(t, ctx.Err)

// 100 from divide() + 2 remaining
require.EqualValues(t, solo.Saldo+16+16+17, member1.Balance())
require.EqualValues(t, solo.Saldo+33+33+34, member2.Balance())
require.EqualValues(t, solo.Saldo+50+50+51, member3.Balance())
// managed to give every one an exact integer amount, so no remainder
require.EqualValues(t, 0, ctx.Balance(nil))
```

Note how after the final `divide` call we ended up with the exact amounts to disperse, so
no remainder iotas were left in the contract account.

Also note how each divide is cumulative to the balances of the members. We have
highlighted this by indicating the separate increases after every `divide` call.

Finally, we will show how to test Views and/or Funcs that return a result. Since solo
executes post() requests synchronously it is possible to have a Func return a result and
test for certain result values

```golang
func TestGetFactor(t *testing.T) {
    ctx := wasmsolo.NewSoloContext(t, dividend.ScName, dividend.OnLoad)
    
    member1 := ctx.NewSoloAgent()
    dividendMember(ctx, member1, 25)
    require.NoError(t, ctx.Err)
    
    member2 := ctx.NewSoloAgent()
    dividendMember(ctx, member2, 50)
    require.NoError(t, ctx.Err)
    
    member3 := ctx.NewSoloAgent()
    dividendMember(ctx, member3, 75)
    require.NoError(t, ctx.Err)
    
    require.EqualValues(t, 3, ctx.Balance(nil))
    
    value := dividendGetFactor(ctx, member3)
    require.NoError(t, ctx.Err)
    require.EqualValues(t, 75, value)
    
    value = dividendGetFactor(ctx, member2)
    require.NoError(t, ctx.Err)
    require.EqualValues(t, 50, value)
    
    value = dividendGetFactor(ctx, member1)
    require.NoError(t, ctx.Err)
    require.EqualValues(t, 25, value)
}
```

Here we first set up the same 3 dispersion factors, and then we retrieve the dispersion
factors for each member in reverse order and verify its value.

In the next section we will describe a few more helper member functions of the
SoloContext.

Next: [Colored Tokens and Time Locks](timelock.md)

