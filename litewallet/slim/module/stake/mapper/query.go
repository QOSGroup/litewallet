package mapper

import (
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
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
	DelegatorAddr   btypes.Address `json:"delegator_address"`
	OwnerAddr       btypes.Address `json:"owner_address"`
	ValidatorPubKey crypto.PubKey  `json:"validator_pub_key"`
	Amount          uint64         `json:"delegate_amount"`
	IsCompound      bool           `json:"is_compound"`
}
