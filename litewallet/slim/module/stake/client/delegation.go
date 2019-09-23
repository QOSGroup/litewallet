package client

import (
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/tx"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	stake_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
)

func CreateDelegation(ctx context.CLIContext, validatorAddress string, coins int64, privKey, chainId string) ([]byte, error) {
	return tx.BuildAndSignTx(ctx, privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		validatorAddr, err := qcliacc.GetValidatorAddrFromValue(validatorAddress)
		if err != nil {
			return nil, err
		}

		delegator, err := qcliacc.GetAddrFromValue(addrben32)
		if err != nil {
			return nil, err
		}

		if coins <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}
		tokens := types.NewInt(coins)

		return &stake_txs.TxCreateDelegation{
			Delegator:     delegator,
			ValidatorAddr: validatorAddr,
			Amount:        tokens,
			IsCompound:    false,
		}, nil
	})
}

func CreateUnbondDelegation(ctx context.CLIContext, validatorAddress string, coins int64, privKey, chainId string) ([]byte, error) {
	return tx.BuildAndSignTx(ctx, privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		validatorAddr, err := qcliacc.GetValidatorAddrFromValue(validatorAddress)
		if err != nil {
			return nil, err
		}

		delegator, err := qcliacc.GetAddrFromValue(addrben32)
		if err != nil {
			return nil, err
		}

		tokens := types.NewInt(coins)
		if !tokens.GT(types.ZeroInt()) {
			return nil, errors.New("unbond QOS amount must gt 0")
		}

		return &stake_txs.TxUnbondDelegation{
			Delegator:     delegator,
			ValidatorAddr: validatorAddr,
			UnbondAmount:  tokens,
			IsUnbondAll:   false,
		}, nil
	})
}

func CreateReDelegationCommand(ctx context.CLIContext, fromValidatorAddr, toValidatorAddr string, coins int64, privKey, chainId string) ([]byte, error) {
	return tx.BuildAndSignTx(ctx, privKey, chainId, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())

		tokens :=types.NewInt(coins)
		if !tokens.GT(types.ZeroInt())  {
			return nil, errors.New("redelegate QOS amount must gt 0")
		}

		delegator, err := qcliacc.GetAddrFromValue(addrben32)
		if err != nil {
			return nil, err
		}

		fromValidatorAddr, err := qcliacc.GetValidatorAddrFromValue(fromValidatorAddr)
		if err != nil {
			return nil, err
		}

		toValidatorAddr, err := qcliacc.GetValidatorAddrFromValue(toValidatorAddr)
		if err != nil {
			return nil, err
		}

		return &stake_txs.TxCreateReDelegation{
			Delegator:         delegator,
			FromValidatorAddr: fromValidatorAddr,
			ToValidatorAddr:   toValidatorAddr,
			Amount:            tokens,
			IsCompound:         true,
			IsRedelegateAll:    false,
		}, nil
	})
}
