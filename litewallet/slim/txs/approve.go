package txs

import (
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
)

// 创建授权
type TxCreateApprove struct {
	types.Approve
}

// 增加授权
type TxIncreaseApprove struct {
	types.Approve
}

// 减少授权
type TxDecreaseApprove struct {
	types.Approve
}

// 使用授权
type TxUseApprove struct {
	types.Approve
}

// 取消授权 Tx
type TxCancelApprove struct {
	From btypes.Address `json:"from"` // 授权账号
	To   btypes.Address `json:"to"`   // 被授权账号
}

// 签名字节
func (tx TxCancelApprove) GetSignData() (ret []byte) {
	ret = append(ret, tx.From...)
	ret = append(ret, tx.To...)

	return ret
}
