package mapper

import (
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

/*

custom path:
/custom/stake/$query path

 query path:
	/delegation/:delegatorAddr/:ownerAddr : 根据delegator和owner查询委托信息(first: delegator)
	/delegations/owner/:ownerAddr : 查询owner下的所有委托信息
	/delegations/delegator/:delegatorAddr : 查询delegator的所有委托信息

return:
  json字节数组
*/

type DelegationQueryResult struct {
	DelegatorAddr            btypes.AccAddress `json:"delegator_address"`
	ValidatorAddr            btypes.ValAddress `json:"validator_address"`
	ValidatorConsensusPubKey string            `json:"validator_cons_pub_key"`
	Amount                   btypes.BigInt     `json:"delegate_amount"`
	IsCompound               bool              `json:"is_compound"`
}
