## Proxy Objects

As stated before, ISCP data is held in key/value stores. In their essence these
are simple, straight-forward key/value dictionaries. Because the data that needs
to be stored by smart contracts can be more complex and hierarchical we provide
a way to build these hierarchical data structures on top of these key/value
stores. To that end we provide a way to nest data structures within a single
key/value store, and hide the complexity of building the more complex keys this
requires. The nested structure is much similar to that used in JSON. We
distinguish two specific kind of container objects: maps and arrays.

A map is a key/value store where the key is one of our supported value objects.
The value associated with a key can again be one of our supported value objects,
but it can also be another container object (map or array), or a serializable
structured object that consists of a number of value objects.

An array can be seen as a special kind of key/value store, where the key is an
integer value with the property that keys form a sequence from 0 to N-1 for an
array with N elements. Arrays always store elements of the same data type, which
can be one of the value objects, or a map object. We do not support arrays of
arrays at this moment.

Maps and arrays are proxy container objects that refer to corresponding
container objects in a hierarchical tree of objects that are created as needed
and live in storage on the host. These proxy container objects are used to get
to the specific proxy value objects that in turn represent the corresponding
data values stored on the host. The way that they are used makes sure that they
have implied (im)mutability and value types. This allows us to leverage the
compiler's type checking to make sure that the constraints on these values are
not unwittingly being broken. These constraints will also be checked at runtime
on the host to prevent malicious client code from bypassing these constraints.

The proxy value objects can be used to determine the existence of a specific
value on the host, and to get and/or set the data value on the host.

To facilitate the distinction between mutability and immutability the proxy
objects come in two flavors. Mutable objects will provide mutable access to
everything in their object tree, while immutable objects only allow immutable
access to everything in their object tree.

WasmLib provides a full set of all possible permutations of (im)mutability,
value type + map, and array. The proxy objects provided are named accordingly.
We decided on a simple, consistent naming scheme. All proxy object names start
with Sc (as in Smart Contract), followed by Immutable or Mutable, followed by
one of the value type names above or Map. In the case where it is an array the
name is subsequently followed by Array.

Examples:

- `ScMutableInt` - proxy to mutable int value
- `ScImmutableString` - proxy to immutable string value
- `ScImmutableColorArray` - proxy to immutable array of immutable color values
- `ScMutableMap` - proxy to mutable map
- `ScImmutableMapArray` - proxy to immutable array of immutable map

Next: [Function Call Context](Context.md)
