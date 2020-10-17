## Java Smart Contracts

Sample smart contracts:

- donatewithfeedback

    Allows for donations and registers feedback associated with the donation.
    The contract owner can at any point decide to withdraw donated funds
    from the contract.

- fairauction

    Allows an auctioneer to auction a number of tokens.
    The contract owner takes a small fee.
    The contract guarantees that the tokens will be sent to the highest bidder,
    and that the losing bidders will be completely refunded. 
    Everyone involved stakes their tokens, so there is no possibility for anyone
    to cheat.
    
- fairroulette

    A simple betting contract. Betters can bet on a random color and after
    a predefined time period the contract will automatically pay the total
    bet amount proportionally to the bet size of the winners. 
    
- helloworld

    ISCP version of the ubiquitous "Hello, world!" program.

- increment

    A simple test contract. All it does is increment a counter value.
    It was used to test basic ISCP capabilities, like persisting state,
    batching requests, and sending (time-locked) request from a contract.

- tokenregistry

    Mints and registers colored tokens in a token registry.

### How to create your own smart contracts

Building a Java smart contract is very simple when using an IntelliJ based
Java development environment. Open the _java_ sub folder in your IntelliJ,
which then provides you with the Java workspace.
 
The easiest way to create a new contract is to copy the _HelloWorld.java_ file
in the _contracts_ package to a properly named new java file within
the _contracts_ package.

Building the project will compile all smart contracts.

`NOTE: currently there is no way to compile the Java contracts into Wasm.
Once this becomes available the intention is that the contracts will simply work
unchanged by implementing the Host calls in the client package.`