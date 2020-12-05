## Go Smart Contracts

Sample smart contracts:

- donatewithfeedback

  Allows for donations and registers feedback associated with the donation. The contract owner can at any point decide
  to withdraw donated funds from the contract.

- fairauction

  Allows an auctioneer to auction a number of tokens. The contract owner takes a small fee. The contract guarantees that
  the tokens will be sent to the highest bidder, and that the losing bidders will be completely refunded. Everyone
  involved stakes their tokens, so there is no possibility for anyone to cheat.

- fairroulette

  A simple betting contract. Betters can bet on a random color and after a predefined time period the contract will
  automatically pay the total bet amount proportionally to the bet size of the winners.

- helloworld

  ISCP version of the ubiquitous "Hello, world!" program.

- increment

  A simple test contract. All it does is increment a counter value. It was used to test basic ISCP capabilities, like
  persisting state, batching requests, and sending (time-locked) request from a contract.

- tokenregistry

  Mints and registers colored tokens in a token registry.

### How to create your own smart contracts

Building a Go smart contract is very simple when using the GoLand IntelliJ based development environment. Open the _
wasplib_ folder in your Goland, which then provides you with the Go workspace.

The easiest way to create a new contract is to copy the _helloworld.go_ file in the _contracts_ sub folder to a properly
named new Go file within the _contracts_ sub folder.

To build the new smart contract select _Run->Edit Configurations_. Add a new configuration based on the _Shell Script_
template, type the _name_ of the new configuration, select _tinygo_build.bat_ in the
_wasplib_ root as the _Script Path_, and enter the name of the new contract as the _script options_, and select the _
wasplib_ root as the _Working Directory_. the new folder as the _working directory_. You can now run this configuration
to compile the smart contract directly to Wasm. Once compilation is successful you will find the resulting Wasm file in
the
_wasplib/wasm_ folder.

