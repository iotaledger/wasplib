## Data Access Proxies

To recap why we need a library to access the sandbox functionality: we cannot call the
ISCP sandbox functions directly. There is no way for the Wasm code to access any memory
outside its own memory space. Therefore, any data that is governed by the ISCP sandbox has
to be copied in and out of that memory space through well-defined protected channels in
the Wasm runtime system.

To make this whole process as seamless as possible the WasmLib interface provides a number
of so-called `proxies`. Proxies are objects that can perform the underlying data transfers
between the separate systems. Proxies are a bit like data references in that regard.
Proxies refer to the actual objects or values stored on the ISCP host and know how to
manipulate them. They provide a consistent interface to the smart contract to achieve
that.

The most basic proxies are value proxies. They refer to a single value instance stored on
the ISCP host. All values are stored as key/value pairs in container objects on the ISCP
host. Value proxies refer to their values through an object id and key id combination that
uniquely defines the value's location in the container object.

Another type of proxies are container proxies. To keep things as simple and understandable
as possible these are limited to only two different kinds. Array proxies and map proxies.
These are enough to be able to define quite complex data structures, because we allow
nesting of container objects. The underlying ISCP sandbox provides access to its data in
the form of simple key/value stores that can have arbitrary byte data for both key and
value. WasmLib provides an abstraction on top of this one-dimensional storage system that
supports nesting of maps and arrays, very similar to the way objects in JSON can be
nested.

A map is a key/value store where the key is one of our supported value types. Maps always
store elements of the same data type, which can be one of our supported value types, a
user-defined data type, or another container object
(map or array).

An array can be seen as a special kind of map. Its key is an integer value with the
property that keys always form a sequence from 0 to N-1 for an array with N elements.
Arrays always store elements of the same data type, which can be one of our supported
value types, a user-defined type, or a map. We do not support arrays of arrays at this
moment.

Here is an example that shows the use of proxies in WasmLib:

![Proxies image](img/Proxies.png)

In this example we have a single map in ISCP state storage that contains a number of
key/value combinations (Key 1 through Key 4). One of them (Key 4)
refers to an array, which in turn contains indexed values stored at indexes 0 through N.
Notice how the WasmLib proxies mirror these exactly. There is a container proxy for every
container, and a value proxy for each value stored. Each container proxy can uniquely
identify the container it references through the container's id. Each value proxy uniquely
identifies the value it references through the container id of the container it is in, and
the key id (or index)
that correlates to its position in the container.

Note that despite the one-to-one correspondence in the example it is not necessary for a
smart contract function to define a proxy for every value or container in ISCP state
storage. In practice a function will only define proxies to data that it actually needs to
access.

In the next section we will go into more detail about the supported data types.

Next: [WasmLib Data Types](types.md)
