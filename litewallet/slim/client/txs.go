package client

import (
	"encoding/hex"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/core/types"
	tendermint_types "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	go_amino "github.com/tendermint/go-amino"
	"time"
)

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTx(remote, hashHexStr string) (btypes.TxResponse, error) {
	cliCtx := context.NewCLIContext(remote)

	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	resTx, err := cliCtx.Client.Tx(hash, false)
	//resTx, err := tx(hash, false)

	if err != nil {
		return btypes.TxResponse{}, err
	}

	resBlocks, err := getBlocksForTxResults(cliCtx, []*tendermint_types.ResultTx{resTx})
	if err != nil {
		return btypes.TxResponse{}, err
	}

	out, err := formatTxResult(txs.Cdc, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

//func tx(hash []byte, prove bool) (*tendermint_types.ResultTx, error) {
//	result := new(tendermint_types.ResultTx)
//	params := map[string]interface{}{
//		"hash":  hash,
//		"prove": prove,
//	}
//	_, err := rpc.Call(txs.Shost, "tx", params, result)
//	if err != nil {
//		return nil, errors.Wrap(err, "Tx")
//	}
//	return result, nil
//}

//// formatTxResults parses the indexed txs into a slice of TxResponse objects.
//func formatTxResults(cdc *go_amino.Codec, resTxs []*ctypes.ResultTx, resBlocks map[int64]*ctypes.ResultBlock) ([]btypes.TxResponse, error) {
//	var err error
//	out := make([]btypes.TxResponse, len(resTxs))
//	for i := range resTxs {
//		out[i], err = formatTxResult(cdc, resTxs[i], resBlocks[resTxs[i].Height])
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	return out, nil
//}

func getBlocksForTxResults(ctx context.CLIContext, resTxs []*tendermint_types.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := ctx.Client.Block(&resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func formatTxResult(cdc *go_amino.Codec, resTx *tendermint_types.ResultTx, resBlock *ctypes.ResultBlock) (btypes.TxResponse, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	return btypes.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), nil
}

func parseTx(cdc *go_amino.Codec, txBytes []byte) (btypes.Tx, error) {
	var tx btypes.Tx

	err := cdc.UnmarshalBinaryBare(txBytes, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
