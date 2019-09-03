package types

import "github.com/QOSGroup/litewallet/litewallet/slim/base/types"

type CoinsTx struct {
	Address    types.Address
	Cointype   string
	Amount     types.BigInt
	ChangeType string //0 plus  1 minus
}
