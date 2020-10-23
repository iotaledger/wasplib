package client

type ScAddress struct {
	address string
}

func NewScAddress(bytes string) *ScAddress {
	return &ScAddress{address: bytes}
	//if len(bytes) != 33 {
	//	panic("address should be 33 bytes")
	//}
	//a := &ScAddress{}
	//copy(a.address[:], bytes)
	//return a
}

func (a *ScAddress) Bytes() string {
	return a.address
}

func (a *ScAddress) Equals(other *ScAddress) bool {
	return a.address == other.address
}

func (a *ScAddress) String() string {
	return a.address
	//return Encode58(a.address[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	color string
}

var IOTA = &ScColor{}
var MINT = &ScColor{}

func init() {
	IOTA.color = "iota"
	MINT.color = "new"
	//for i := range MINT.color {
	//	MINT.color[i] = 0xff
	//}
}

func NewScColor(bytes string) *ScColor {
	return &ScColor{color: bytes}
	//if len(bytes) != 32 {
	//	panic("color should be 32 bytes")
	//}
	//a := &ScColor{}
	//copy(a.color[:], bytes)
	//return a
}

func (c *ScColor) Bytes() string {
	return c.color
}

func (c *ScColor) Equals(other *ScColor) bool {
	return c.color == other.color
}

func (c *ScColor) String() string {
	return c.color
	//return Encode58(c.color[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScRequestId struct {
	id string
}

func NewScRequestId(bytes string) *ScRequestId {
	return &ScRequestId{id: bytes}
	//if len(bytes) != 34 {
	//	panic("request id should be 34 bytes")
	//}
	//a := &ScRequestId{}
	//copy(a.id[:], bytes)
	//return a
}

func (r *ScRequestId) Bytes() string {
	return r.id
}

func (r *ScRequestId) Equals(other *ScRequestId) bool {
	return r.id == other.id
}

func (r *ScRequestId) String() string {
	return r.id
	//return Encode58(r.id[:])
}

//func (r *ScRequestId) TransactionId() *ScTransactionId {
//	return NewScTransactionId(r.id[:32])
//}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransactionId struct {
	id string
}

func NewScTransactionId(bytes string) *ScTransactionId {
	return &ScTransactionId{id: bytes}
	//if len(bytes) != 32 {
	//	panic("transaction id should be 33 bytes")
	//}
	//a := &ScTransactionId{}
	//copy(a.id[:], bytes)
	//return a
}

func (t *ScTransactionId) Bytes() string {
	return t.id
}

func (t *ScTransactionId) Equals(other *ScTransactionId) bool {
	return t.id == other.id
}

func (t *ScTransactionId) String() string {
	return t.id
	//return Encode58(t.id[:])
}
