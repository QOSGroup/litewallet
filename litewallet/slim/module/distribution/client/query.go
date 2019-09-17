package client

import (
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/distribution/mapper"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/distribution/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
)

func QueryDelegatorIncomeInfo(remote, privKey, ownerAddr string) (mapper.DelegatorIncomeInfoQueryResult, error) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privKey + "\"}"
	err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())

	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	var owner btypes.Address
	var delegator btypes.Address

	if o, err := qcliacc.GetAddrFromValue(ownerAddr); err == nil {
		owner = o
	}

	if d, err := qcliacc.GetAddrFromValue(addrben32); err == nil {
		delegator = d
	}

	path := types.BuildQueryDelegatorIncomeInfoCustomQueryPath(delegator, owner)
	res, err := cliCtx.Query(path, []byte(""))
	if err != nil {
		return mapper.DelegatorIncomeInfoQueryResult{}, err
	}

	var result mapper.DelegatorIncomeInfoQueryResult
	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return result, err
}

func QueryCommunityFeePool(remote string) (btypes.BigInt, error) {
	cliCtx := context.NewCLIContext(remote).WithCodec(txs.Cdc)

	res, err := cliCtx.Query(fmt.Sprintf("/store/%s/key", types.MapperName), types.BuildCommunityFeePoolKey())
	if err != nil {
		return btypes.BigInt{}, err
	}

	var result btypes.BigInt
	cliCtx.Codec.MustUnmarshalBinaryBare(res, &result)
	return result, err
}
