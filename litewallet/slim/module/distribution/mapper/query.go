package mapper

import (
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
)

type DelegatorIncomeInfoQueryResult struct {
	OwnerAddr             btypes.Address `json:"owner_address"`
	ValidatorPubKey       crypto.PubKey  `json:"validator_pub_key"`
	PreviousPeriod        uint64         `json:"previous_validator_period"`
	BondToken             uint64         `json:"bond_token"`
	CurrentStartingHeight uint64         `json:"earns_starting_height"`
	FirstDelegateHeight   uint64         `json:"first_delegate_height"`
	HistoricalRewardFees  btypes.BigInt  `json:"historical_rewards"`
	LastIncomeCalHeight   uint64         `json:"last_income_calHeight"`
	LastIncomeCalFees     btypes.BigInt  `json:"last_income_calFees"`
}
