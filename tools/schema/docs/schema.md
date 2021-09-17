## Smart Contract Schema Tool

Smart contracts need to be very robust. Preferably it would be very hard to make mistakes
when writing smart contract code. The generic nature of WasmLib allows for a lot of
flexibility, but it also provides you with a lot of opportunities to make mistakes. In
addition, there is a lot of repetitive coding involved. The setup code that is needed for
every smart contract must follow strict rules. You also want to assure that certain
functions can only be called by specific entities, and that function parameters values
have been properly checked before their usage.

The best way to increase robustness is by using a code generator that will take care of
most repetitive coding tasks. A code generator only needs to be debugged once, after which
the generated code is 100% accurate and trustworthy. Another advantage of code generation
is that you can regenerate code to correctly reflect any changes to the smart contract
interface. A code generator can also help you by generating wrapper code that limits what
you can do to mirror the intent behind it. This enables compile-time enforcing of some
aspects of the defined smart contract behavior. A code generator can also support multiple
different programming languages.

During our initial experiences with creating demo smart contracts for WasmLib we quickly
identified a number of areas where there was a lot of repetitive coding going on. Examples
of such repetition were:

* setting up the `on_load` function and keeping it up to date
* checking function access rights
* verifying function parameter types
* verifying the presence of mandatory function parameters
* setting up access to state, params, and results maps
* defining common strings as constants

To facilitate the code generation we decided to use a _schema definition file_ for smart
contracts. In such a schema definition file all aspects of a smart contract that should be
known by someone who wants to use the contract are clearly defined in a single place. This
schema definition then becomes the source of truth for how the smart contract works.

The schema file defines things like the state variables that the smart contract uses, the
Funcs and Views that the contract implements, the access rights for each function, the
input parameters and output results for each function, and additional data structures that
the contract uses.

With detailed schema information readily available in a single location it suddenly
becomes possible to do a lot more than just generating repetitive code fragments. We can
use the schema information to generate interfaces for functions, parameters, results, and
state that use strict compile-time type-checking. This reduces the likelihood that errors
are introduced significantly.

Another advantage of knowing everything about important smart contract aspects is that it
is possible to generate constants to prevent repeating of typo-prone key strings and
precalculate necessary values like Hnames and encode them as constants instead of having
the code recalculate them every time they are needed.

Similarly, since we know all static keys that are going to be used by the smart contract
in advance, we can now generate code that will negotiate the corresponding key IDs with
the host only once in the `on_load` function and cache those values for use in future
function calls.

The previous two optimizations mean that the code becomes both simpler and more 
efficient. Note
that all the improvements described above are independent of the programming language 
used.

Future additions to the schema tool that we envision are the automatic generation of smart
contract interface classes for use with client side Javascript, and automatic generation
of a web API for smart contracts. The schema file can also provide a starting point for
other tooling, for example a tool that automatically audits a smart contract

In the next section we will look at how the schema tool works.

Next: [Using the Schema Tool](usage.md)
