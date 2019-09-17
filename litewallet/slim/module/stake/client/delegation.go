package client

import (
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	stake_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
)

func CreateSignedDelegation(validatorAddress string, coins int64, privKey, chainId string) ([]byte, error) {
	return ctxs.BuildAndSignTx(privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		owner, err := types.GetAddrFromBech32(validatorAddress)
		if err != nil {
			return nil, err
		}

		delegator, err := types.GetAddrFromBech32(addrben32)
		if err != nil {
			return nil, err
		}

		tokens := coins
		if tokens <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}

		return &stake_txs.TxCreateDelegation{
			Delegator:      delegator,
			ValidatorOwner: owner,
			Amount:         uint64(tokens),
			IsCompound:     false,
		}, nil
	})
}

func CreateSignedUnbondDelegation(validatorAddress string, coins int64, privKey, chainId string) ([]byte, error) {
	return ctxs.BuildAndSignTx(privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		owner, err := types.GetAddrFromBech32(validatorAddress)
		if err != nil {
			return nil, err
		}

		delegator, err := types.GetAddrFromBech32(addrben32)
		if err != nil {
			return nil, err
		}

		tokens := coins
		if tokens <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}

		return &stake_txs.TxUnbondDelegation{
			Delegator:      delegator,
			ValidatorOwner: owner,
			UnbondAmount:   uint64(tokens),
			IsUnbondAll:    false,
		}, nil
	})
}

func CreateReDelegationCommand(fromValidatorAddr, toValidatorAddr string, tokens int64, privKey, chainId string) ([]byte, error) {
	return ctxs.BuildAndSignTx(privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		if tokens <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}

		delegator, err := qcliacc.GetAddrFromValue(addrben32)
		if err != nil {
			return nil, err
		}

		fromValidatorOwner, err := qcliacc.GetAddrFromValue(fromValidatorAddr)
		if err != nil {
			return nil, err
		}

		toValidatorOwner, err := qcliacc.GetAddrFromValue(toValidatorAddr)
		if err != nil {
			return nil, err
		}

		return &stake_txs.TxCreateReDelegation{
			Delegator:          delegator,
			FromValidatorOwner: fromValidatorOwner,
			ToValidatorOwner:   toValidatorOwner,
			Amount:             uint64(tokens),
			IsCompound:         true,
			IsRedelegateAll:    false,
		}, nil
	})
}
