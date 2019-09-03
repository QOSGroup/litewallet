package types

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

// 授权 Common 结构
type Approve struct {
	From types.Address `json:"from"` // 授权账号
	To   types.Address `json:"to"`   // 被授权账号
	QOS  types.BigInt  `json:"qos"`  // QOS
	QSCs types.QSCs    `json:"qscs"` // QSCs
}

func NewApprove(from types.Address, to types.Address, qos types.BigInt, qscs types.QSCs) Approve {
	return Approve{
		From: from,
		To:   to,
		QOS:  qos.NilToZero(),
		QSCs: qscs,
	}
}

// 签名字节
func (approve Approve) GetSignData() (ret []byte) {
	approve.QOS = approve.QOS.NilToZero()

	ret = append(ret, approve.From...)
	ret = append(ret, approve.To...)
	ret = append(ret, approve.QOS.String()...)
	for _, coin := range approve.QSCs {
		ret = append(ret, []byte(coin.Name)...)
		ret = append(ret, []byte(coin.Amount.String())...)
	}

	return ret
}
