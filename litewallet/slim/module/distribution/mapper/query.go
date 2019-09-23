package mapper

import (
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

type DelegatorIncomeInfoQueryResult struct {
	OperatorAddress       btypes.ValAddress `json:"validator_address"`
	ConsPubKey            string            `json:"consensus_pubkey"`
	PreviousPeriod        int64             `json:"previous_validator_period"`
	BondToken             btypes.BigInt     `json:"bond_token"`
	CurrentStartingHeight int64             `json:"earns_starting_height"`
	FirstDelegateHeight   int64             `json:"first_delegate_height"`
	HistoricalRewardFees  btypes.BigInt     `json:"historical_rewards"`
	LastIncomeCalHeight   int64             `json:"last_income_calHeight"`
	LastIncomeCalFees     btypes.BigInt     `json:"last_income_calFees"`
}