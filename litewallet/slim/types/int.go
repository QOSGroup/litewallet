package types

import "math/big"

type Int struct {
	i *big.Int
}

func (i Int) IsZero() bool {
	return i.i.Sign() == 0
}

func (i Int) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

//genStdSendTx for the Tx send operation
// NewInt constructs BigInt from int64
func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}
