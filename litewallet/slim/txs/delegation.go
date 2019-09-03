package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

type TxCreateDelegation struct {
	Delegator      btypes.Address //委托人
	ValidatorOwner btypes.Address //验证者Owner
	Amount         uint64         //委托QOS数量
	IsCompound     bool           //定期收益是否复投
}

var _ txs.ITx = (*TxCreateDelegation)(nil)

func (tx *TxCreateDelegation) GetSignData() (ret []byte) {
	ret = append(ret, tx.Delegator...)
	ret = append(ret, tx.ValidatorOwner...)
	ret = append(ret, btypes.Int2Byte(int64(tx.Amount))...)
	ret = append(ret, btypes.Bool2Byte(tx.IsCompound)...)
	return
}

type TxUnbondDelegation struct {
	Delegator      btypes.Address //委托人
	ValidatorOwner btypes.Address //验证者Owner
	UnbondAmount   uint64         //unbond数量
	IsUnbondAll    bool           //是否全部解绑, 为true时覆盖UnbondAmount
}

var _ txs.ITx = (*TxUnbondDelegation)(nil)

func (tx *TxUnbondDelegation) GetSignData() (ret []byte) {
	ret = append(ret, tx.Delegator...)
	ret = append(ret, tx.ValidatorOwner...)
	ret = append(ret, btypes.Int2Byte(int64(tx.UnbondAmount))...)
	ret = append(ret, btypes.Bool2Byte(tx.IsUnbondAll)...)
	return
}
