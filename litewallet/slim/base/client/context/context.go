package context

import (
	"fmt"
	cmn "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/libs/common"
	rpcclient "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/client"
	"github.com/pkg/errors"
	go_amino "github.com/tendermint/go-amino"
)

// CLIContext implements a typical CLI context created in SDK modules for
// transaction handling and queries.
type CLIContext struct {
	Codec     *go_amino.Codec
	Client    *rpcclient.HTTP
	Height    int64
	NodeURI   string
	TrustNode bool
}

// NewCLIContext returns a new initialized CLIContext with parameters from the
// command line using Viper.
func NewCLIContext(remote string) CLIContext {
	rpc := rpcclient.NewHTTP(remote, "/websocket")

	return CLIContext{
		Client:    rpc,
		NodeURI:   remote,
		Height:    0,
		TrustNode: true,
	}
}

// WithCodec returns a copy of the context with an updated codec.
func (ctx CLIContext) WithCodec(cdc *go_amino.Codec) CLIContext {
	ctx.Codec = cdc
	return ctx
}

func (ctx CLIContext) GetNode() (*rpcclient.HTTP, error) {
	if ctx.Client == nil {
		return nil, errors.New("no RPC client defined")
	}
	return ctx.Client, nil
}

func (ctx CLIContext) Query(path string, data []byte) (res []byte, err error) {
	return ctx.query(path, cmn.HexBytes(data))
}

// query performs a query from a Tendermint node with the provided store name
// and path.
func (ctx CLIContext) query(path string, key cmn.HexBytes) (res []byte, err error) {
	node, err := ctx.GetNode()
	if err != nil {
		return res, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: ctx.Height,
		Prove:  ctx.TrustNode,
	}

	result, err := node.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return res, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return res, fmt.Errorf(resp.Log)
	}

	return resp.Value, nil
}
