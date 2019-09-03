package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
)

type AdvertisersTx struct {
	Tx *ctypes.CoinsTx
}

func (tx AdvertisersTx) GetSignData() (ret []byte) {
	ret = append(ret, tx.Tx.Address.Bytes()...)
	ret = append(ret, types.Int2Byte(tx.Tx.Amount.Int64())...)
	ret = append(ret, []byte(tx.Tx.Cointype)...)
	ret = append(ret, []byte(tx.Tx.ChangeType)...)
	return
}
