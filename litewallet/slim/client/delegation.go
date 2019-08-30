package client

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
)

func CreateSignedDelegation(validatorAddress string, coins int64, privkey, chainid string) (string, error) {
	return ctxs.BuildAndSignTx(privkey, chainid, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(ctxs.PREF_ADD, key.PubKey().Address().Bytes())
		owner, err := types.GetAddrFromBech32(addrben32)
		if err != nil {
			return nil, err
		}

		delegator, err := types.GetAddrFromBech32(validatorAddress)
		if err != nil {
			return nil, err
		}

		tokens := coins
		if tokens <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}

		return &ctxs.TxCreateDelegation{
			Delegator:      delegator,
			ValidatorOwner: owner,
			Amount:         uint64(tokens),
			IsCompound:     true,
		}, nil
	})
}

func CreateSignedUnbondDelegation(validatorAddress string, coins int64, privkey, chainid string) (string, error) {
	return ctxs.BuildAndSignTx(privkey, chainid, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(ctxs.PREF_ADD, key.PubKey().Address().Bytes())
		owner, err := types.GetAddrFromBech32(addrben32)
		if err != nil {
			return nil, err
		}

		delegator, err := types.GetAddrFromBech32(validatorAddress)
		if err != nil {
			return nil, err
		}

		tokens := coins
		if tokens <= 0 {
			return nil, errors.New("delegate QOS amount must gt 0")
		}

		return &ctxs.TxUnbondDelegation{
			Delegator:      delegator,
			ValidatorOwner: owner,
			UnbondAmount:   uint64(tokens),
			IsUnbondAll:    false,
		}, nil
	})
}
