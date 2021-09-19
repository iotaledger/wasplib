## Using the Schema Tool

We tried to make the creation of smart contracts as simple as possible. The `schema`
tool will assist you along the way as unobtrusively as possible. We will now walk you
through the steps to create a new smart contract from scratch.

First you decide on a central folder where you want to keep all your smart contracts. Each
smart contract you create will be maintained in a separate subfolder of this folder. We
will use certain naming conventions that the schema tool expects throughout this usage
section. First we will select a camel case name for our smart contract. For our example
here we will use the name `MySmartContract`.

Once you know what your smart contract will be named it is time to set up your subfolder.
Simply navigate to the central smart contract folder, and run the schema tool's
initialization function there like this:

`schema -init MySmartContract`

This command will create a subfolder named `mysmartcontract` and generate an initial
schema definition file `schema.json` inside this subfolder. Note that the generated
subfolder name is all lower case. This is due to best practices for package names both in
Rust and in Go. The generated schema.json looks like this:

```json
{
  "name": "MySmartContract",
  "description": "MySmartContract description",
  "structs": {},
  "typedefs": {},
  "state": {
    "owner": "AgentID // current owner of this smart contract"
  },
  "funcs": {
    "init": {
      "params": {
        "owner": "?AgentID // optional owner of this smart contract"
      }
    },
    "setOwner": {
      "access": "owner // current owner of this smart contract",
      "params": {
        "owner": "AgentID // new owner of this smart contract"
      }
    }
  },
  "views": {
    "getOwner": {
      "results": {
        "owner": "AgentID // current owner of this smart contract"
      }
    }
  }
}
```

Schema.json has been pre-populated with all sections that you could need, and some
functions that allow you to maintain the ownership of the smart contract. Now that
schema.json exists it is up to you to modify it further to reflect the requirements of
your smart contract.

We use camel case naming convention throughout schema.json when naming items. Function and
variable names always start with a lower case character, and type names always start with
an upper case character.

The first thing you may want to do before you do anything else is to modify the
`description` field to something more sensible. And if you already know how to use the
schema tool then now is the moment to fill out some sections with the definitions you know
you will need.

The next step is to go into the new `mysmartcontract` subfolder and run the schema tool
there to generate the initial code. If you just want to generate Rust code you run the
schema tool with the `-rust` option like this:

`schema -rust`

If you just want to generate Go code you run the schema tool with the `-go` option like
this:

`schema -go`

And if you want to generate both Rust and Go code you need to specify both options like
this:

`schema -rust -go`

The schema tool will generate a complete set of source files for the desired language(s).
After generating the Rust code for the first time you should modify the Cargo.toml file to
your likings, and potentially add the new project to a Rust workspace. Cargo.toml will not
be regenerated once it already exists. The generated files together readily compile into a
Wasm file by using the appropriate command:

* For Rust: `wasm-pack build`. This will use the `src` subfolder that contains all Rust
  source files. The only file in this folder that you should edit manually is
  `mysmartcontract.rs`. All other files will be regenerated and overwritten whenever the
  schema tool is run again.
* For Go: `tinygo build -target wasm wasmmain/main.go`. This will use the go source files
  in the current folder. The only file in this folder that you should edit manually is
  `mysmartcontract.go`. All other files will be regenerated and overwritten whenever the
  schema tool is run again.

For now, we will focus on the Rust code that is generated, but the Go code is essentially
identical, barring some language idiosyncrasy differences. Just view .rs files next to .go
files with the same name, and you will see what we mean.

Anyway, to show you an example of the initially generated Rust code, `mysmartcontract.rs`
looks like this before you even start modifying it:

```rust
// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) {
    if f.params.owner().exists() {
        f.state.owner().set_value(f.params.owner().value());
        return;
    }
    f.state.owner().set_value(ctx.contract_creator());
}

pub fn func_set_owner(_ctx: &ScFuncContext, f: &SetOwnerContext) {
    f.state.owner().set_value(f.params.owner().value());
}

pub fn view_get_owner(_ctx: &ScViewContext, f: &GetOwnerContext) {
    f.results.owner().set_value(f.state.owner().value());
}
```

As you can see the schema tool generated an initial working version of the functions that
are used to maintain the smart contract owner that will suffice for most cases.

For a smooth building experience it is a good idea to set up a build rule in your build
environment that runs the schema tool with the required parameters whenever the
schema.json file changes. That way regeneration of files is automatic and you no longer
have to start the schema tool manually each time after changing schema.json. Note that the
schema tool will only regenerate the code when it finds that schema.json has been modified
since the last time it generated the code. You can force the schema tool to regenerate all
code by adding the `-force` flag to its command line parameter.

In the next section we will look at how a smart contract uses state storage.

Next:[Smart Contract State](state.md)
