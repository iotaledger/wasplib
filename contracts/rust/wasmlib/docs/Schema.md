## Smart Contract Schema

Smart contracts need to be very robust. Preferably it would be very hard to make
mistakes when writing smart contract code. The generic nature of WasmLib allows
for a lot of flexibility, but it also provides you with a lot of opportunities
to make mistakes. In addition, there is a lot of repetitive coding involved. The
setup code that is needed for every smart contract must follow strict rules. You
also want to assure that certain functions can only be called by specific
entities, and that function parameters and return values have been properly
checked before their usage.

The best way to increase robustness is by using a code generator that will take
care of most repetitive coding tasks. A code generator only needs to be debugged
once, after which the generated code is 100% accurate and trustworthy. Another
advantage of code generation is that you can regenerate code to correctly
reflect any changes. A code generator can also help you by generating wrapper
code that limits what you can do to mirror the intent behind it. This enables
compile-time enforcing of the defined smart contract rules. A code generator 
can also support multiple different programming languages. 

We call the code generator that we developed the `schema` tool, for obvious
reasons. The tool can take a JSON schema definition file for a smart contract
and generate the corresponding wrappers that ensure all the above and more. In
fact, it will generate a complete set of compilable smart contract code that
only needs the developer to fill in the provided skeleton functions.

We are currently using a JSON schema definition file because it was easier to
use an existing JSON parser than to create a specific schema definition language
with corresponding parser. In the future we may change this but for now we will
focus on the core concepts of the schema definition.

There are other uses we envision for the schema definition file in the future,
like being able to automatically generate a web API for accessing the smart
contract, or a client-side interface for web applications, or providing a
starting point for a tool that automatically audits a smart contract.

Here is the schema.json file for the `dividend` example. It has purposely been
kept simple and does not show all features of the schema definition file yet. We
will explore the missing features in later articles.

```json
{
  "name": "Dividend",
  "description": "Simple dividend smart contract",
  "state": {
    "memberList": "[]Address // array with all the recipients of this dividend",
    "members": "[Address]Int64 // map with all the recipient factors of this dividend",
    "owner": "AgentID // owner of contract, the only one who can call 'member' func",
    "totalFactor": "Int64 // sum of all recipient factors"
  },
  "funcs": {
    "init": {
      "params": {
        "owner": "?AgentID // optional owner of contract, defaults to contract creator"
      }
    },
    "member": {
      "access": "owner // only defined owner of contract can add members",
      "params": {
        "address": "Address // address of dividend recipient",
        "factor": "Int64 // relative division factor"
      }
    },
    "divide": {
    },
    "setOwner": {
      "access": "owner // only defined owner of contract can change owner",
      "params": {
        "owner": "AgentID // new owner of smart contract"
      }
    }
  },
  "views": {
    "getFactor": {
      "params": {
        "address": "Address // address of dividend recipient"
      },
      "results": {
        "factor": "Int64 // relative division factor"
      }
    }
  }
}
```

Note the main sections in this schema.json:

* The header defines a `name` and `description` for the smart contract. In the
  future we may define additional global properties.
* The `state` section defines what the state storage looks like.
* The `funcs` section defines the properties of each func.
* The `views` section defines the properties of each view.

In the next section we will examine the `state` section in more detail.

Next: [Smart Contract State](State.md)
