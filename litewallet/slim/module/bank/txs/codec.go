package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/baseabci"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
	"github.com/tendermint/go-amino"
)

var Cdc = baseabci.MakeQBaseCodec()

func init() {
	types.RegisterCodec(Cdc)
	RegisterCodec(Cdc)
}

func RegisterCodec(cdc *amino.Codec) {

	cdc.RegisterConcrete(&TxTransfer{}, "bank/txs/TxTransfer", nil)
	//cdc.RegisterConcrete(&TxInvariantCheck{}, "transfer/txs/TxInvariantCheck", nil)
}
