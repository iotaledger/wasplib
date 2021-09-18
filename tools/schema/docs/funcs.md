## Function Definitions

Here is the full schema.json file for the `dividend` example. We will now focus on
its `funcs` and `views` sections. Since they are structured identically we will only need
to explain the layout of these sections once.

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

As you can see each of the `funcs` and `views` sections defines their functions in the
same way. The only resulting difference is in the way the schema tool generates code for
them. The code generated for Funcs will be able to inspect and modify the smart contract
state, whereas the code generated for Views will only be able to inspect the state.

Functions are defined as named subsections in schema.json. The name of the subsection will
become the name of the function. Under each function subsection in turn there can be 3
optional subsections.

* `access` indicates who is allowed to access the function.
* `params` holds the field definitions that describe the function parameters.
* `results` holds the field definitions that describe the function results.

We will now examine each subsection in more detail. In the next section we will first look
at the `access` subsection.

Next: [Limiting Access](access.md)
