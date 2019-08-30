package types

import (
	"encoding/json"
	"math/big"
)

// NewInt constructs BigInt from int64
//func NewInt(n int64) BigInt {
//	return BigInt{big.NewInt(n)}
//}

type BigInt struct {
	i *big.Int
}

// NewInt constructs BigInt from int64
func NewInt(n int64) BigInt {
	return BigInt{big.NewInt(n)}
}

func (i BigInt) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

// Add adds BigInt from another
func (i BigInt) Add(i2 BigInt) (res BigInt) {
	res = BigInt{add(i.i, i2.i)}
	// Check overflow
	if res.i.BitLen() > 255 {
		panic("BigInt overflow")
	}
	return
}

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func (bi BigInt) IsNil() bool {
	return bi.i == nil
}

func (i BigInt) NilToZero() BigInt {
	if i.IsNil() {
		return ZeroInt()
	}
	return i
}

// ZeroInt returns BigInt value with zero
func ZeroInt() BigInt { return BigInt{big.NewInt(0)} }

func (i BigInt) String() string {
	return i.i.String()
}

// MarshalAmino defines custom encoding scheme
func (i BigInt) MarshalAmino() (string, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalAmino(i.i)
}

// UnmarshalAmino defines custom decoding scheme
func (i *BigInt) UnmarshalAmino(text string) error {
	if i.i == nil { // Necessary since default BigInt initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalAmino(i.i, text)
}

// MarshalJSON defines custom encoding scheme
func (i BigInt) MarshalJSON() ([]byte, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *BigInt) UnmarshalJSON(bz []byte) error {
	if i.i == nil { // Necessary since default BigInt initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
}

// MarshalAmino for custom encoding scheme
func marshalAmino(i *big.Int) (string, error) {
	bz, err := i.MarshalText()
	return string(bz), err
}

// UnmarshalAmino for custom decoding scheme
func unmarshalAmino(i *big.Int, text string) (err error) {
	return i.UnmarshalText([]byte(text))
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i *big.Int) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}
