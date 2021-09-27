This is the actual client to interact with the fairroulette smart contract.

It uses the wasp_client library for the low level interaction. In the future, this won't be nessecary,
as we will provide an external library at some point.   


It is responsible for: 
* de/encoding smart contract returned data
* invoking smart contract functions
* listening and handling emitted events
* wrapping the invokactions into easy to use functions (See: placeBetOnLedger)

//TODO: Go into more detail
