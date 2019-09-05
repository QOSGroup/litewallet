package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	tendermint_types "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"github.com/tendermint/go-amino"
	go_amino "github.com/tendermint/go-amino"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	types "github.com/tendermint/tendermint/rpc/lib/types"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	protoHTTP  = "http"
	protoHTTPS = "https"
	protoWSS   = "wss"
	protoWS    = "ws"
	protoTCP   = "tcp"
)

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTx(hashHexStr string) (btypes.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return btypes.TxResponse{}, err
	}

	//resTx, err := txs.RPC.Tx(hash, false)
	resTx, err := Tx(hash, false)

	if err != nil {
		return btypes.TxResponse{}, err
	}

	resBlocks, err := getBlocksForTxResults([]*tendermint_types.ResultTx{resTx})
	if err != nil {
		return btypes.TxResponse{}, err
	}

	out, err := formatTxResult(txs.Cdc, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

func Tx(hash []byte, prove bool) (*tendermint_types.ResultTx, error) {
	result := new(tendermint_types.ResultTx)
	params := map[string]interface{}{
		"hash":  hash,
		"prove": prove,
	}
	_, err := Call("tx", params, result)
	if err != nil {
		return nil, errors.Wrap(err, "Tx")
	}
	return result, nil
}

func Call(method string, params map[string]interface{}, result interface{}) (interface{}, error) {
	cdc := amino.NewCodec()
	ctypes.RegisterAmino(cdc)
	address, client := makeHTTPClient(txs.Shost)

	request, err := types.MapToRequest(cdc, types.JSONRPCStringID("jsonrpc-client"), method, params)
	if err != nil {
		return nil, err
	}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	// log.Info(string(requestBytes))
	requestBuf := bytes.NewBuffer(requestBytes)
	// log.Info(Fmt("RPC request to %v (%v): %v", c.remote, method, string(requestBytes)))
	httpResponse, err := client.Post(address, "text/json", requestBuf)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close() // nolint: errcheck

	responseBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	// 	log.Info(Fmt("RPC response: %v", string(responseBytes)))
	return unmarshalResponseBytes(cdc, responseBytes, result)
}

func makeHTTPClient(remoteAddr string) (string, *http.Client) {
	protocol, address, dialer := makeHTTPDialer(remoteAddr)
	return protocol + "://" + address, &http.Client{
		Transport: &http.Transport{
			// Set to true to prevent GZIP-bomb DoS attacks
			DisableCompression: true,
			Dial:               dialer,
		},
	}
}

// TODO: Deprecate support for IP:PORT or /path/to/socket
func makeHTTPDialer(remoteAddr string) (string, string, func(string, string) (net.Conn, error)) {
	// protocol to use for http operations, to support both http and https
	clientProtocol := protoHTTP

	parts := strings.SplitN(remoteAddr, "://", 2)
	var protocol, address string
	if len(parts) == 1 {
		// default to tcp if nothing specified
		protocol, address = protoTCP, remoteAddr
	} else if len(parts) == 2 {
		protocol, address = parts[0], parts[1]
	} else {
		// return a invalid message
		msg := fmt.Sprintf("Invalid addr: %s", remoteAddr)
		return clientProtocol, msg, func(_ string, _ string) (net.Conn, error) {
			return nil, errors.New(msg)
		}
	}

	// accept http as an alias for tcp and set the client protocol
	switch protocol {
	case protoHTTP, protoHTTPS:
		clientProtocol = protocol
		protocol = protoTCP
	case protoWS, protoWSS:
		clientProtocol = protocol
	}

	// replace / with . for http requests (kvstore domain)
	trimmedAddress := strings.Replace(address, "/", ".", -1)
	return clientProtocol, trimmedAddress, func(proto, addr string) (net.Conn, error) {
		return net.Dial(protocol, address)
	}
}

func unmarshalResponseBytes(cdc *amino.Codec, responseBytes []byte, result interface{}) (interface{}, error) {
	// Read response.  If rpc/core/types is imported, the result will unmarshal
	// into the correct type.
	// log.Notice("response", "response", string(responseBytes))
	var err error
	response := &types.RPCResponse{}
	err = json.Unmarshal(responseBytes, response)
	if err != nil {
		return nil, errors.Errorf("Error unmarshalling rpc response: %v", err)
	}
	if response.Error != nil {
		return nil, errors.Errorf("Response error: %v", response.Error)
	}
	// Unmarshal the RawMessage into the result.
	err = cdc.UnmarshalJSON(response.Result, result)
	if err != nil {
		return nil, errors.Errorf("Error unmarshalling rpc response result: %v", err)
	}
	return result, nil
}

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

func getBlocksForTxResults(resTxs []*tendermint_types.ResultTx) (map[int64]*ctypes.ResultBlock, error) {
	resBlocks := make(map[int64]*ctypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := txs.RPC.Block(&resTx.Height)
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
