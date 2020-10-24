package client

type ScAddress struct {
	address [33]byte
}

func NewScAddress(bytes []byte) *ScAddress {
	if len(bytes) != 33 {
		panic("address should be 33 bytes")
	}
	a := &ScAddress{}
	copy(a.address[:], bytes)
	return a
}

func (a *ScAddress) Bytes() []byte {
	return a.address[:]
}

func (a *ScAddress) Equals(other *ScAddress) bool {
	return a.address == other.address
}

func (a *ScAddress) String() string {
	return base58Encode(a.address[:])
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScColor struct {
	color [32]byte
}

var IOTA = &ScColor{}
var MINT = &ScColor{}

func init() {
	for i := range MINT.color {
		MINT.color[i] = 0xff
	}
}

func NewScColor(bytes []byte) *ScColor {
	if len(bytes) != 32 {
		panic("color should be 32 bytes")
	}
	a := &ScColor{}
	copy(a.color[:], bytes)
	return a
}

func (c *ScColor) Bytes() []byte {
	return c.color[:]
}

func (c *ScColor) Equals(other *ScColor) bool {
	return c.color == other.color
}

func (c *ScColor) String() string {
	return base58Encode(c.color[:])
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

func base58Encode(bytes []byte) string {
	return NewScContext().Utility().Base58Encode(bytes)
}
