package client

import (
	"fmt"
	btxs "github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
	"github.com/pkg/errors"
)

type operateType int

const (
	createType operateType = iota
	increaseType
	decreaseType
	useType
	cancleType
)

func QueryApprove(toAddrStr, privKey string) (*types.Approve, error) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
	err := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())

	queryPath := "store/approve/key"

	fromAddr, err := btypes.GetAddrFromBech32(addrben32)
	if err != nil {
		return nil, err
	}

	toAddr, err := btypes.GetAddrFromBech32(toAddrStr)
	if err != nil {
		return nil, err
	}

	output, err := txs.Query(queryPath, types.BuildApproveKey(fromAddr, toAddr))
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, errors.New("approve does not exist")
	}

	approve := types.Approve{}
	txs.Cdc.MustUnmarshalBinaryBare(output, &approve)
	return &approve, nil
}

func CreateApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
	return applyApprove(createType, toAddrStr, coinsStr, privKey, chainId)
}

func IncreaseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
	return applyApprove(increaseType, toAddrStr, coinsStr, privKey, chainId)
}

func DecreaseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
	return applyApprove(decreaseType, toAddrStr, coinsStr, privKey, chainId)
}

func UseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
	return applyApprove(useType, toAddrStr, coinsStr, privKey, chainId)
}

func CancelApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
	return applyApprove(cancleType, toAddrStr, coinsStr, privKey, chainId)
}

func applyApprove(operType operateType, toAddrStr, coinstr, privKey, chainId string) ([]byte, error) {
	return txs.BuildAndSignTx(privKey, chainId, func() (btxs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
		err := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())

		fromAddr, err := btypes.GetAddrFromBech32(addrben32)
		if err != nil {
			return nil, err
		}

		toAddr, err := btypes.GetAddrFromBech32(toAddrStr)
		if err != nil {
			return nil, err
		}

		if operType == cancleType {
			return txs.TxCancelApprove{
				From: fromAddr,
				To:   toAddr,
			}, nil
		}

		qos, qscs, err := NewParseCoins(coinstr)
		if err != nil {
			return nil, err
		}
		appr := types.NewApprove(fromAddr, toAddr, qos, qscs)

		switch operType {
		case createType:
			return txs.TxCreateApprove{Approve: appr}, nil
		case increaseType:
			return txs.TxIncreaseApprove{Approve: appr}, nil
		case decreaseType:
			return txs.TxDecreaseApprove{Approve: appr}, nil
		case useType:
			return txs.TxUseApprove{Approve: appr}, nil
		default:
			return nil, errors.New("operType invalid")
		}
	})
}
