package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
)

type TxTransfer struct {
	Senders   types.TransItems `json:"senders"`   // 发送集合
	Receivers types.TransItems `json:"receivers"` // 接收集合
}

//type TxTransfer struct {
//	Senders   []TransItem `json:"senders"`   // 发送集合
//	Receivers []TransItem `json:"receivers"` // 接收集合
//}

var _ txs.ITx = (*TxTransfer)(nil)

// 签名字节
func (tx TxTransfer) GetSignData() (ret []byte) {
	for _, sender := range tx.Senders {
		ret = append(ret, sender.Address...)
		ret = append(ret, (sender.QOS.NilToZero()).String()...)
		ret = append(ret, sender.QSCs.String()...)
	}
	for _, receiver := range tx.Receivers {
		ret = append(ret, receiver.Address...)
		ret = append(ret, (receiver.QOS.NilToZero()).String()...)
		ret = append(ret, receiver.QSCs.String()...)
	}

	return ret
}
