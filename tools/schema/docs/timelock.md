## Colored Tokens and Time Locks

Let's examine some less commonly used member functions of the SoloContext. We will 
switch to the `fairauction` example to show their usage. Here is the startAuction() 
function of the fairauction test suite:

```golang
var (
    auctioneer *wasmsolo.SoloAgent
    tokenColor wasmlib.ScColor
)

func startAuction(t *testing.T) *wasmsolo.SoloContext {
    ctx := wasmsolo.NewSoloContract(t, fairauction.ScName, fairauction.OnLoad)
    
    // set up auctioneer account and mint some tokens to auction off
    auctioneer = ctx.NewSoloAgent()
    tokenColor, ctx.Err = auctioneer.Mint(10)
    require.NoError(t, ctx.Err)
    require.EqualValues(t, solo.Saldo-10, auctioneer.Balance())
    require.EqualValues(t, 10, auctioneer.Balance(tokenColor))
    
    // start the auction
    sa := fairauction.ScFuncs.StartAuction(ctx.Sign(auctioneer))
    sa.Params.Color().SetValue(tokenColor)
    sa.Params.MinimumBid().SetValue(500)
    sa.Params.Description().SetValue("Cool tokens for sale!")
    transfer := ctx.Transfer()
    transfer.Set(wasmlib.IOTA, 25) // deposit, must be >=minimum*margin
    transfer.Set(tokenColor, 10)   // the tokens to auction
    sa.Func.Transfer(transfer).Post()
    require.NoError(t, ctx.Err)
    return ctx
}
```

The function first sets up the SoloContext as usual, and then it performs quite a bit of
extra work. This is because we want the startAuction() function to start an auction, so
that the tests that subsequently use startAuction() can then focus on testing all kinds of
bidding and auction results.

First, we're going to need an agent that functions as the `auctioneer`. This auctioneer
will auction off some colored tokens. To provide the auctioneer with colored tokens we use
the `Mint()` method to convert 10 of his plain iota tokens into colored tokens. The mint
process will assign the color value, which is equal to the hash of the Tangle transaction
that minted them. We save the resulting ScColor value in `tokenColor`. Note that both 
`auctioneer` and `tokenColor` are global variables that are accessible by any test 
that needs them.

Next we check that no error occurred during the minting process, and then we verify that
the auctioneer now has 10 less plain iota and also has a balance of 10 tokens with the
saved token color in its address. Notice how we use the same Balance() method to retrieve
both balances. When the token color parameter is omitted, the Balance() method defaults to
returning the balance of plain iotas in the address.

Now we are going to start the auction by calling the `startAuction` function of the
fairauction contract. We get the function descriptor in the usual way, but we also call
the `Sign()` method of the SoloContext to make sure that the transaction we're about to
post takes its tokens from the auctioneer address and signs the transaction with the
corresponding private key. Very often you don't care who posts a request, and we have set
it up for you in such a way that by default tokens come from the chain originator address,
which has been seeded with tokens just for this occasion. But whenever it is important
where the tokens come from, or who invokes the request, you need to specify the agent
involved by using the Sign() method.

Next we set up the function parameters as usual. Note how we pass the saved tokenColor for
example. After the parameters have been set up we see something new happening. We create
a `Transfer` proxy and initialize it with the 25 iota that we need to deposit plus the 10
tokens of the saved tokenColor that we are auctioning. Next we use the `Transfer()` method
to pass this proxy before posting the request. This is exactly how we would do it from
within the smart contract code. We have a shorthand function called TransferIotas() that
can be used instead when all you need to transfer is plain iotas and which encapsulates
the creation of the Transfer proxy and the initialization with the required amount of
iotas.

Finally, we make sure there was no error while posting the request and return the 
SoloContext. That concludes the startAuction() function.

Here is the first test function that uses our startAuction() function:

```golang
func TestFaStartAuction(t *testing.T) {
    ctx := startAuction(t)
    
    // note 1 iota should be stuck in the delayed finalize_auction
    require.EqualValues(t, 25-1, ctx.Balance(nil))
    require.EqualValues(t, 10, ctx.Balance(nil, tokenColor))
    
    // auctioneer sent 25 deposit + 10 tokenColor
    require.EqualValues(t, solo.Saldo-25-10, auctioneer.Balance())
    require.EqualValues(t, 0, auctioneer.Balance(tokenColor))
    require.EqualValues(t, 0, ctx.Balance(auctioneer))
    
    // remove pending finalize_auction from backlog
    ctx.AdvanceClockBy(61 * time.Minute)
    require.True(t, ctx.WaitForPendingRequests(1))
}
```

The `startAuction` function of the smart contract will have posted a time-locked request
to the `finalizeAuction` function by using the Delay() method. This request needed 1 iota
for the request, but the request is still 'in transit' until it is unlocked. We can verify
the contract balance after the transfer of 25 iota plus 10 colorToken, minus the 1 iota
still locked. Note how we again have an account Balance() method where the color parameter
can be omitted, in which case it defaults to the account balance of plain iotas.

We also verify the address balance of the auctioneer after sending the startAuction 
request. And double-check that no tokens ended up in his contract account.

The final 2 lines of the code are used to remove the pending `finalizeAuction` request 
from the backlog. First we move the logical clock forward to a point when that request 
is supposed to have triggered. Then we wait for this request to actually be processed. 
Note that this will happen in a separate goroutine in the background, so we explicitly 
wait for the request counters to catch up with the one request that is pending.

The WaitForPendingRequests() method can also be used whenever a smart contract 
function is known to Post() a request to itself. Such requests are not immediately 
executed, but added to the backlog. So you need to wait for these pending requests to 
actually be processed. The advantage here is that you can inspect the in-between state,
which means that you can test even a function that posts a request in isolation.