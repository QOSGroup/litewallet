package sdksource

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"os"
)


const ctxAccStoreName = "acc"

//NewCLIContext is used to init the config context without using Viper, the argues are all from the input of the func
func newCLIContext(rootDir,node,chainID string) context.CLIContext {
	var (
		rpc rpcclient.Client

	)

	//init the rpc instance
	nodeURI := node
	if nodeURI == "" {
		fmt.Printf("The nodeURI can not be nil for the rpc connection!")
	}
	rpc = rpcclient.NewHTTP(nodeURI, "/websocket")

	//create the verifier for the LCD verification
	//var trustNode bool
	//trustNode = false
	//if trustNode {
	//	fmt.Printf("The default value for the trustNode is false!")
	//}
	//chainID := ChainID
	//home := rootDir

	//cacheSize := 10 // TODO: determine appropriate cache size
	//verifier, err := tmliteProxy.NewVerifier(
	//	chainID, filepath.Join(rootDir, ".gaiacli", ".gaialite"),
	//	rpc, log.NewNopLogger(), cacheSize,
	//)


	//if err != nil {
	//	fmt.Printf("Create verifier failed: %s\n", err.Error())
	//	fmt.Printf("Please check network connection and verify the address of the node to connect to\n")
	//	os.Exit(1)
	//}

	CliContext := context.CLIContext{
		Client:        rpc,
		Output:        os.Stdout,
		NodeURI:       nodeURI,
		AccountStore:  auth.StoreKey,
		//Verifier:    verifier,
		//BroadcastMode: broadcastMode,
	}
	return CliContext

}

// TxBuilder implements a transaction context created in SDK modules.
//type TxBuilder struct {
//	txEncoder          sdk.TxEncoder
//	keybase            crkeys.Keybase
//	accountNumber      uint64
//	sequence           uint64
//	gas                uint64
//	gasAdjustment      float64
//	simulateAndExecute bool
//	chainID            string
//	memo               string
//	fees               sdk.Coins
//	gasPrices          sdk.DecCoins
//}

// NewTxBuilderFromCLI returns a new initialized TxBuilder with parameters input
//func newTxBuilderFromCLI(ChainID string) authtxb.TxBuilder {
//	txBldr := authtxb.TxBuilder{
//		chainID:            ChainID,
//	}
//	var txBldr authtxb.TxBuilder
//	return txBldr
//}

//node discovery with round-robin algorithm
//var i = 0
//var servers = []string{"127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8003"}
//
//// Balance returns one of the servers based using round-robin algorithm
//func Balance() string {
//	server := servers[i]
//	i++
//
//	// it means that we reached the end of servers
//	// and we need to reset the counter and start
//	// from the beginning
//	if i >= len(servers) {
//		i = 0
//	}
//	return server
//}
//
//
//func main() {
//	// requests loop
//	for i := 0; i < 20; i++ {
//		fmt.Println(Balance())
//	}
//}