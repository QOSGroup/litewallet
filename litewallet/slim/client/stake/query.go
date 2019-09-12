package stake

import (
	"encoding/hex"
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/client/stake/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	flagActive   = "active"
	activeDesc   = "active"
	inactiveDesc = "inactive"

	inactiveRevokeDesc        = "Revoked"
	inactiveMissVoteBlockDesc = "Kicked"
	inactiveMaxValidatorDesc  = "Replaced"
	inactiveDoubleDesc        = "DoubleSign"
)

type validatorDisplayInfo struct {
	Owner           btypes.Address    `json:"owner"`
	ValidatorAddr   string            `json:"validatorAddress"`
	ValidatorPubKey crypto.PubKey     `json:"validatorPubkey"`
	BondTokens      uint64            `json:"bondTokens"` //不能超过int64最大值
	Description     types.Description `json:"description"`
	Commission      types.Commission  `json:"commission"`

	Status         string    `json:"status"`
	InactiveDesc   string    `json:"InactiveDesc"`
	InactiveTime   time.Time `json:"inactiveTime"`
	InactiveHeight uint64    `json:"inactiveHeight"`

	BondHeight uint64 `json:"bondHeight"`
}

func toValidatorDisplayInfo(validator types.Validator) validatorDisplayInfo {
	info := validatorDisplayInfo{
		Owner:           validator.Owner,
		ValidatorPubKey: validator.ValidatorPubKey,
		BondTokens:      validator.BondTokens,
		Description:     validator.Description,
		InactiveTime:    validator.InactiveTime,
		InactiveHeight:  validator.InactiveHeight,
		BondHeight:      validator.BondHeight,
		Commission:      validator.Commission,
	}

	if validator.Status == types.Active {
		info.Status = activeDesc
	} else {
		info.Status = inactiveDesc
	}

	if validator.InactiveCode == types.Revoke {
		info.InactiveDesc = inactiveRevokeDesc
	} else if validator.InactiveCode == types.MissVoteBlock {
		info.InactiveDesc = inactiveMissVoteBlockDesc
	} else if validator.InactiveCode == types.MaxValidator {
		info.InactiveDesc = inactiveMaxValidatorDesc
	} else if validator.InactiveCode == types.DoubleSign {
		info.InactiveDesc = inactiveDoubleDesc
	}

	info.ValidatorAddr = strings.ToUpper(hex.EncodeToString(validator.ValidatorPubKey.Address()))

	return info
}

func QueryValidatorInfo(remote, address string) (types.Validator, error) {
	cliCtx := context.NewCLIContext(remote)
	ownerAddress, err := qcliacc.GetAddrFromValue(address)
	if err != nil {
		return types.Validator{}, err
	}
	validator, err := getValidator(cliCtx, ownerAddress)
	return validator, nil
}

func getValidator(ctx context.CLIContext, ownerAddress btypes.Address) (types.Validator, error) {
	result, err := ctx.Client.ABCIQueryWithOptions(string(types.BuildStakeStoreQueryPath()), types.BuildOwnerWithValidatorKey(ownerAddress), buildQueryOptions())
	if err != nil {
		return types.Validator{}, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return types.Validator{}, errors.New("owner does't have validator")
	}

	var address btypes.Address
	txs.Cdc.UnmarshalBinaryBare(valueBz, &address)

	key := types.BuildValidatorKey(address)
	result, err = ctx.Client.ABCIQueryWithOptions(string(types.BuildStakeStoreQueryPath()), key, buildQueryOptions())
	if err != nil {
		return types.Validator{}, err
	}

	valueBz = result.Response.GetValue()
	if len(valueBz) == 0 {
		return types.Validator{}, errors.New("response empty value")
	}

	var validator types.Validator
	txs.Cdc.UnmarshalBinaryBare(valueBz, &validator)
	return validator, nil
}

func buildQueryOptions() client.ABCIQueryOptions {
	//height := viper.GetInt64(bctypes.FlagHeight)
	//if height <= 0 {
	//	height = 0
	//}

	//trust := viper.GetBool(bctypes.FlagTrustNode)

	return client.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	}
}

func QueryValidators(remote string) ([]byte, error) {
	cliCtx := context.NewCLIContext(remote)

	opts := client.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	}

	subspace := "/store/validator/subspace"
	result, err := cliCtx.Client.ABCIQueryWithOptions(subspace, []byte{0x01}, opts)

	if err != nil {
		return nil, err
	}

	valueBz := result.Response.GetValue()
	if len(valueBz) == 0 {
		return nil, errors.New("response empty value")
	}

	var validators []validatorDisplayInfo

	var vKVPair []store.KVPair
	err = txs.Cdc.UnmarshalBinaryLengthPrefixed(valueBz, &vKVPair)
	if err != nil {
		return nil, err
	}
	for _, kv := range vKVPair {
		var validator types.Validator
		fmt.Println(kv.Value)
		txs.Cdc.UnmarshalBinaryBare(kv.Value, &validator)
		validators = append(validators, toValidatorDisplayInfo(validator))
	}
	return txs.Cdc.MarshalJSON(validators)
}
