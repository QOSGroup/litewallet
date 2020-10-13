package sdksource

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

const ctxAccStoreName = "acc"

//NewCLIContext is used to init the config context without using Viper, the argues are all from the input of the func
func newCLIContext(rootDir, node, chainID string) context.CLIContext {
	var (
		rpc rpcclient.Client
	)

	//init the rpc instance
	nodeURI := node
	if nodeURI == "" {
		fmt.Printf("The nodeURI can not be nil for the rpc connection!")
	}
	rpc = rpcclient.NewHTTP(nodeURI, "/websocket")



	CliContext := context.CLIContext{
		Client:       rpc,
		Output:       os.Stdout,
		NodeURI:      nodeURI,
		AccountStore: auth.StoreKey,
		//Verifier:    verifier,
		//BroadcastMode: broadcastMode,
	}
	return CliContext

}

