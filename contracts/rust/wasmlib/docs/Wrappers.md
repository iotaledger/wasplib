## Function Wrappers

In the previous sections we explained some of what can be achieved through the
definitions in the schema,json file. In this section we will explain how the
schema tool uses the schema.json file to generate all kinds of repetitive
boilerplate functionality.

First, the schema tool generates a file `consts.rs` that contains symbolic
constants for all identifiers used in the code. This prevents typos in strings
and such and allows the development environment to use autocompletion as it
already knows all possible valid identifiers. Here are the constants that are
generated for the `divide` example:

```rust
// name and description of this smart contract
pub const SC_NAME: &str = "dividend";
pub const SC_DESCRIPTION: &str = "Simple dividend smart contract";

// precalculated Hname value for the smart contract name
pub const HSC_NAME: ScHname = ScHname(0xcce2e239);

// all parameter identifier names that are used
pub const PARAM_ADDRESS: &str = "address";
pub const PARAM_FACTOR: &str = "factor";
pub const PARAM_OWNER: &str = "owner";

// all result identifier names that are used
pub const RESULT_FACTOR: &str = "factor";

// all state variable names that are used
pub const STATE_MEMBER_LIST: &str = "memberList";
pub const STATE_MEMBERS: &str = "members";
pub const STATE_OWNER: &str = "owner";
pub const STATE_TOTAL_FACTOR: &str = "totalFactor";

// names of all functions in this smart contract
pub const FUNC_DIVIDE: &str = "divide";
pub const FUNC_INIT: &str = "init";
pub const FUNC_MEMBER: &str = "member";
pub const FUNC_SET_OWNER: &str = "setOwner";
pub const VIEW_GET_FACTOR: &str = "getFactor";

// precalculated Hnames for all functions in this smart contract
pub const HFUNC_DIVIDE: ScHname = ScHname(0xc7878107);
pub const HFUNC_INIT: ScHname = ScHname(0x1f44d644);
pub const HFUNC_MEMBER: ScHname = ScHname(0xc07da2cb);
pub const HFUNC_SET_OWNER: ScHname = ScHname(0x2a15fe7b);
pub const HVIEW_GET_FACTOR: ScHname = ScHname(0x0ee668fe);
```

In the section about [Function Call Context](Context.md) we saw that it is
necessary for the creator of a smart contract to correctly set up an `on_load`
function that refers to each function in the smart contract:

```rust
fn on_load() {
    let exports = ScExports::new();
    exports.add_func("divide", func_divide);
    exports.add_func("init", func_init);
    exports.add_func("member", func_member);
    exports.add_func("setOwner", func_set_owner);
    exports.add_view("getFactor", view_get_factor);
}
```

Well, we have some good news for you: the schema tool will handle this
automatically for you from now on. It can find all the information it needs to
do this in the `funcs` and `views` subsections of the schema.json file. Since
the schema tool will regenerate all code after every change in the schema.json
file, it is no longer possible to forget to add a definition, or worse, leave an
outdated definition in the `on_load` function. Here is what the schema tool will
generate for you:

```rust
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DIVIDE, func_divide_thunk);
    exports.add_func(FUNC_INIT, func_init_thunk);
    exports.add_func(FUNC_MEMBER, func_member_thunk);
    exports.add_func(FUNC_SET_OWNER, func_set_owner_thunk);
    exports.add_view(VIEW_GET_FACTOR, view_get_factor_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}
```

### Thunk Functions

The first thing you will notice is that it uses the symbolic constants from
consts.rs instead of strings for the names of the functions. The second thing
you will notice is that it added `_thunk` to each function. In computer
programming, a thunk is a function used to inject a code into another function.
We insert the thunk function instead of the original at the point where it will
be called, and the thunk functions as a wrapper function around the original
function. We will use the thunk function to preprocess the parameters and set up
the structures necessary before calling the actual user function. Note that the
schema tool can regenerate the thunk function without having any impact on the
code in the user function. This means that most changes made to schema.json will
translate into seamless changes to the structures and preprocessing. Once you
start using this you will notice that it is very unusual for a change in
schema.json to force a change in the user code. It will mostly be addition of
new parameters or functions, and the user will of course have to adapt his code
to those changes.

For example, here is the thunk function generated for the `member` function of
the `dividend` example:

```rust
pub struct MemberContext {
    params: ImmutableMemberParams,
    state: MutableDividendState,
}

fn func_member_thunk(ctx: &ScFuncContext) {
    ctx.log("dividend.funcMember");
    // only defined owner of contract can add members
    let access = ctx.state().get_agent_id("owner");
    ctx.require(access.exists(), "access not set: owner");
    ctx.require(ctx.caller() == access.value(), "no permission");

    let f = MemberContext {
        params: ImmutableMemberParams {
            id: OBJ_ID_PARAMS,
        },
        state: MutableDividendState {
            id: OBJ_ID_STATE,
        },
    };
    ctx.require(f.params.address().exists(), "missing mandatory address");
    ctx.require(f.params.factor().exists(), "missing mandatory factor");
    func_member(ctx, &f);
    ctx.log("dividend.funcMember ok");
}
```

The thunk function will automatically insert logging of when the function is
called and when it completed successfully. You can also see the code
(including comment) that checks the access rights for the function based on the
owner state variable. It gets a proxy to the state variable, checks for its
existence, and then requires the agent ID of the caller to be equal to the agent
ID found in the owner variable. Notice the usage of the `require`
function instead of a more verbose if-statement.

Next it creates a function-specific context structure MemberContext, which
contains proxies to the params map and the state map. This allows the user to
immediately start using those without having to set them up every time.

The thunk function then requires the mandatory `address` and `factor`
parameters to actually exist, and finally it calls the actual user function,
which now has a slightly modified signature:

```rust
pub fn func_member(ctx: &ScFuncContext, f: &MemberContext) {}
```

Notice that user functions will receive two parameters now. First the original
ISCP function context (ScFuncContext or ScViewContext), and second the
function-specific context that contains proxies to access the state map and
optional params map. When a function produces results this structure will also
contain a proxy to the results map.

### Static Key Caching

We haven't yet discussed the other piece of code that the schema tool inserts in
the `on_load` function. The idea behind the `on_load` function is that it is
called only once upon initialization of the smart contract. This means this is
an ideal location to put other code that only needs to run once.

Since the schema tool knows every static key that is being used for the state,
params, and result maps, it is possible to add some code that caches the key IDs
for each of these static keys. Normally, it would be necessary with each usage
of a key name to have a small exchange with the ISCP to ask it to look up the
corresponding key ID. By doing this only once upon initialization for the known
static keys, we can remove this execution overhead from subsequent function
calls.

The schema tool keeps track of all the static key IDs in a static array. Since
it knows the index of each static key in this array, it can generate code that
directly uses the key ID by indexing this array instead of needing the overhead
of asking the ISCP for the key ID.

The schema tool automatically generates keys.rs, which contains the necessary
code for this mechanism:

```rust
// indexes for each of the static params/results/state keys
pub const IDX_PARAM_ADDRESS: usize = 0;
pub const IDX_PARAM_FACTOR: usize = 1;
pub const IDX_PARAM_OWNER: usize = 2;
pub const IDX_RESULT_FACTOR: usize = 3;
pub const IDX_STATE_MEMBER_LIST: usize = 4;
pub const IDX_STATE_MEMBERS: usize = 5;
pub const IDX_STATE_OWNER: usize = 6;
pub const IDX_STATE_TOTAL_FACTOR: usize = 7;

pub const KEY_MAP_LEN: usize = 8;

// array of all static keys, which each key at proper index
// this array is used in the on_load initialization code
pub const KEY_MAP: [&str; KEY_MAP_LEN] = [
    PARAM_ADDRESS,
    PARAM_FACTOR,
    PARAM_OWNER,
    RESULT_FACTOR,
    STATE_MEMBER_LIST,
    STATE_MEMBERS,
    STATE_OWNER,
    STATE_TOTAL_FACTOR,
];

// array that will receive all key IDs received from the ISCP
pub static mut IDX_MAP: [Key32; KEY_MAP_LEN] = [Key32(0); KEY_MAP_LEN];

// wrapper function to simplify access to the unsafe static array
pub fn idx_map(idx: usize) -> Key32 {
    unsafe {
        IDX_MAP[idx]
    }
}
```

You will find that the IDX_ constants found in the keys.rs file will be used to
directly access static key IDs in the proxy structures that are automatically 
generated by the schema tool in params.rs, results.rs, and state.rs.

In the next section we will start fleshing out the `dividend` code now that 
the schema tool provided us with the necessary parts.

Next: [Dividend Code](Dividend.md)
