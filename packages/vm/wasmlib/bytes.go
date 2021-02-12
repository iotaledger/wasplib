// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

type BytesDecoder struct {
	data []byte
}

func NewBytesDecoder(data []byte) *BytesDecoder {
	return &BytesDecoder{data: data}
}

func (d *BytesDecoder) Address() *ScAddress {
	return NewScAddressFromBytes(d.Bytes())
}

func (d *BytesDecoder) AgentId() *ScAgentId {
	return NewScAgentIdFromBytes(d.Bytes())
}

func (d *BytesDecoder) Bytes() []byte {
	size := d.Int()
	if len(d.data) < int(size) {
		panic("Cannot decode bytes")
	}
	value := d.data[:size]
	d.data = d.data[size:]
	return value
}

func (d *BytesDecoder) ChainId() *ScChainId {
	return NewScChainIdFromBytes(d.Bytes())
}

func (d *BytesDecoder) Color() *ScColor {
	return NewScColorFromBytes(d.Bytes())
}

func (d *BytesDecoder) ContractId() *ScContractId {
	return NewScContractIdFromBytes(d.Bytes())
}

func (d *BytesDecoder) Hash() *ScHash {
	return NewScHashFromBytes(d.Bytes())
}

func (d *BytesDecoder) Hname() ScHname {
	return NewScHnameFromBytes(d.Bytes())
}

func (d *BytesDecoder) Int() int64 {
	// leb128 decoder
	val := int64(0)
	s := 0
	for {
		b := int8(d.data[0])
		d.data = d.data[1:]
		val |= int64(b&0x7f) << s
		if b >= 0 {
			if int8(val>>s)&0x7f != b&0x7f {
				panic("Integer too large")
			}
			// extend int7 sign to int8
			if (b & 0x40) != 0 {
				b |= -0x80
			}
			// extend int8 sign to int64
			return val | (int64(b) << s)
		}
		s += 7
		if s >= 64 {
			panic("integer representation too long")
		}
	}
}

func (d *BytesDecoder) String() string {
	return string(d.Bytes())
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type BytesEncoder struct {
	data []byte
}

func NewBytesEncoder() *BytesEncoder {
	return &BytesEncoder{data: make([]byte, 0, 128)}
}

func (e *BytesEncoder) Address(value *ScAddress) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) AgentId(value *ScAgentId) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) Bytes(value []byte) *BytesEncoder {
	e.Int(int64(len(value)))
	e.data = append(e.data, value...)
	return e
}

func (e *BytesEncoder) ChainId(value *ScChainId) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) Color(value *ScColor) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) ContractId(value *ScContractId) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) Data() []byte {
	return e.data
}

func (e *BytesEncoder) Hash(value *ScHash) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) Hname(value ScHname) *BytesEncoder {
	return e.Bytes(value.Bytes())
}

func (e *BytesEncoder) Int(value int64) *BytesEncoder {
	// leb128 encoder
	for {
		b := byte(value)
		s := b & 0x40
		value >>= 7
		if (value == 0 && s == 0) || (value == -1 && s != 0) {
			e.data = append(e.data, b&0x7f)
			return e
		}
		e.data = append(e.data, b|0x80)
	}
}

func (e *BytesEncoder) String(value string) *BytesEncoder {
	return e.Bytes([]byte(value))
}
