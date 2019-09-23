package client

import (
	cmn "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/libs/common"
	tendermint_types "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/core/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/lib/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/types"
	"github.com/pkg/errors"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

type HTTP struct {
	remote string
	rpc    *client.JSONRPCClient

	*baseRPCClient
}

type baseRPCClient struct {
	caller client.JSONRPCCaller
}

func NewHTTP(remote, wsEndpoint string) *HTTP {
	rc := client.NewJSONRPCClient(remote)
	cdc := rc.Codec()
	ctypes.RegisterAmino(cdc)
	rc.SetCodec(cdc)

	return &HTTP{
		rpc:           rc,
		remote:        remote,
		baseRPCClient: &baseRPCClient{caller: rc},
	}
}

func (c *baseRPCClient) ABCIQuery(path string, data cmn.HexBytes) (*ctypes.ResultABCIQuery, error) {
	return c.ABCIQueryWithOptions(path, data, DefaultABCIQueryOptions)
}

func (c *baseRPCClient) ABCIQueryWithOptions(path string, data cmn.HexBytes, opts ABCIQueryOptions) (*ctypes.ResultABCIQuery, error) {
	result := new(ctypes.ResultABCIQuery)
	_, err := c.caller.Call("abci_query",
		map[string]interface{}{"path": path, "data": data, "height": opts.Height, "prove": opts.Prove},
		result)
	if err != nil {
		return nil, errors.Wrap(err, "ABCIQuery")
	}
	return result, nil
}

//func (c *baseRPCClient) BroadcastTxCommit(tx types.Tx) (*ctypes.ResultBroadcastTxCommit, error) {
//	result := new(ctypes.ResultBroadcastTxCommit)
//	_, err := c.caller.Call("broadcast_tx_commit", map[string]interface{}{"tx": tx}, result)
//	if err != nil {
//		return nil, errors.Wrap(err, "broadcast_tx_commit")
//	}
//	return result, nil
//}

func (c *baseRPCClient) BroadcastTxAsync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return c.broadcastTX("broadcast_tx_async", tx)
}

func (c *baseRPCClient) BroadcastTxSync(tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	return c.broadcastTX("broadcast_tx_sync", tx)
}

func (c *baseRPCClient) broadcastTX(route string, tx types.Tx) (*ctypes.ResultBroadcastTx, error) {
	result := new(ctypes.ResultBroadcastTx)
	_, err := c.caller.Call(route, map[string]interface{}{"tx": tx}, result)
	if err != nil {
		return nil, errors.Wrap(err, route)
	}
	return result, nil
}

func (c *baseRPCClient) Block(height *int64) (*tendermint_types.ResultBlock, error) {
	result := new(tendermint_types.ResultBlock)
	_, err := c.caller.Call("block", map[string]interface{}{"height": height}, result)
	if err != nil {
		return nil, errors.Wrap(err, "Block")
	}
	return result, nil
}

func (c *baseRPCClient) Tx(hash []byte, prove bool) (*types.ResultTx, error) {
	result := new(types.ResultTx)
	params := map[string]interface{}{
		"hash":  hash,
		"prove": prove,
	}
	_, err := c.caller.Call("tx", params, result)
	if err != nil {
		return nil, errors.Wrap(err, "Tx")
	}
	return result, nil
}
