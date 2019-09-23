package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

type TxCreateDelegation struct {
	Delegator     btypes.AccAddress //委托人
	ValidatorAddr btypes.ValAddress // 验证人
	Amount        btypes.BigInt     //委托QOS数量
	IsCompound    bool              //定期收益是否复投
}

var _ txs.ITx = (*TxCreateDelegation)(nil)

func (tx *TxCreateDelegation) GetSignData() (ret []byte) {
	ret = Cdc.MustMarshalBinaryBare(tx)
	return
}

type TxUnbondDelegation struct {
	Delegator     btypes.AccAddress //委托人
	ValidatorAddr btypes.ValAddress //验证者
	UnbondAmount  btypes.BigInt     //unbond数量
	IsUnbondAll   bool              //是否全部解绑, 为true时覆盖UnbondAmount
}

var _ txs.ITx = (*TxUnbondDelegation)(nil)


func (tx *TxUnbondDelegation) GetSignData() (ret []byte) {
	ret = Cdc.MustMarshalBinaryBare(tx)
	return
}

type TxCreateReDelegation struct {
	Delegator         btypes.AccAddress //委托人
	FromValidatorAddr btypes.ValAddress //原委托验证人
	ToValidatorAddr   btypes.ValAddress //现委托验证人
	Amount            btypes.BigInt     //委托数量
	IsRedelegateAll   bool              //
	IsCompound        bool              //
}

var _ txs.ITx = (*TxCreateReDelegation)(nil)

func (tx *TxCreateReDelegation) GetSignData() (ret []byte) {
	ret = Cdc.MustMarshalBinaryBare(tx)
	return
}
