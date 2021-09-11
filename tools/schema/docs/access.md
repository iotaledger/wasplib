## Limiting Access

The optional `access` subsection is made of a single definition. It indicates the agent
who is allowed to access the function. When this definition is omitted anyone is allowed
to call the function. When the definition is present it should be an access identifier,
optionally followed by an explanatory comment. Access identifiers can be one of the
following:

* `self`: only the smart contract itself can call this function
* `chain`: only the chain owner can call this function
* `creator`: only the contract creator can call this function
* anything else: the name of an AgentID or []AgentID variable in state storage. Only the
  agent(s) defined there can call this function. When this option is used you should also
  provide functionality that can initialize and/or modify this variable. As long as this
  state variable has not been set, nobody is allowed to call this function.

The schema tool will automatically generate code to properly check the access rights of
the agent that called the function before the actual function is called.

You can see usage examples of the access identifier in the schema.json above, where the
state variable `owner` is used as an access identifier. The `init` function initializes
this state variable, and the `setOwner` function can be used only by the current owner to
set a new owner. Finally, the `member` function can also only be called by the current
owner.

Note that there can be different access identifiers for different functions as needed. You
can set up as many access identifiers as you like.

In the next section we will look at the `params` subsection.

Next: [Function Parameters](params.md)
