package client

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	approvetypes "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
)

type operateType int

const (
	createType operateType = iota
	increaseType
	decreaseType
	useType
	cancleType
)

func QueryApprove(cliCtx context.CLIContext,toAddrStr, privKey string) (*approvetypes.Approve, error) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
	err := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())

	queryPath := "store/approve/key"

	fromAddr, err := account.GetAddrFromValue(addrben32)
	if err != nil {
		return nil, err
	}

	toAddr, err := account.GetAddrFromValue(toAddrStr)
	if err != nil {
		return nil, err
	}

	output, err := txs.Query(queryPath, approvetypes.BuildApproveKey(fromAddr, toAddr))
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, errors.New("approve does not exist")
	}

	approve := approvetypes.Approve{}
	txs.Cdc.MustUnmarshalBinaryBare(output, &approve)
	return &approve, nil
}

//func CreateApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
//	return applyApprove(createType, toAddrStr, coinsStr, privKey, chainId)
//}
//
//func IncreaseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
//	return applyApprove(increaseType, toAddrStr, coinsStr, privKey, chainId)
//}
//
//func DecreaseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
//	return applyApprove(decreaseType, toAddrStr, coinsStr, privKey, chainId)
//}
//
//func UseApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
//	return applyApprove(useType, toAddrStr, coinsStr, privKey, chainId)
//}
//
//func CancelApprove(toAddrStr, coinsStr, privKey, chainId string) ([]byte, error) {
//	return applyApprove(cancleType, toAddrStr, coinsStr, privKey, chainId)
//}
//
//func applyApprove(operType operateType, toAddrStr, coinstr, privKey, chainId string) ([]byte, error) {
//	return txs.BuildAndSignTx(privKey, chainId, func() (btxs.ITx, error) {
//		var key ed25519local.PrivKeyEd25519
//		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
//		err := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
//		if err != nil {
//			fmt.Println(err)
//		}
//		addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())
//
//		fromAddr, err := btypes.GetAddrFromBech32(addrben32)
//		if err != nil {
//			return nil, err
//		}
//
//		toAddr, err := btypes.GetAddrFromBech32(toAddrStr)
//		if err != nil {
//			return nil, err
//		}
//
//		if operType == cancleType {
//			return txs2.TxCancelApprove{
//				From: fromAddr,
//				To:   toAddr,
//			}, nil
//		}
//
//		qos, qscs, err := client.NewParseCoins(coinstr)
//		if err != nil {
//			return nil, err
//		}
//		appr := types.NewApprove(fromAddr, toAddr, qos, qscs)
//
//		switch operType {
//		case createType:
//			return txs2.TxCreateApprove{Approve: appr}, nil
//		case increaseType:
//			return txs2.TxIncreaseApprove{Approve: appr}, nil
//		case decreaseType:
//			return txs2.TxDecreaseApprove{Approve: appr}, nil
//		case useType:
//			return txs2.TxUseApprove{Approve: appr}, nil
//		default:
//			return nil, errors.New("operType invalid")
//		}
//	})
//}
