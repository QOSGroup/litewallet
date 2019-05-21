package javasdk

import (
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	cskeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"os"
	"time"
)

var cdc = app.MakeCodec()

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

	////create the verifier for the LCD verification
	//var trustNode bool
	//trustNode = false
	//if trustNode {
	//	fmt.Printf("The default value for the trustNode is false!")
	//}
	////chainID := ChainID
	////home := rootDir
	//
	//cacheSize := 10 // TODO: determine appropriate cache size
	//verifier, err := tmliteProxy.NewVerifier(
	//	chainID, filepath.Join(rootDir, ".gaiacli", ".gaialite"),
	//	rpc, log.NewNopLogger(), cacheSize,
	//)
	//
	//
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
		//Verifier:      verifier,

	}
	return CliContext

}
//Make the transfer with Async mode
func TransferAsync(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr string) string {
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}
	//SetKeyBase(rootDir)
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	var info cskeys.Info
	var err error
	info, err = kb.Get(fromName)
	if err != nil {
		fmt.Printf("could not find key %s\n", fromName)
		os.Exit(1)
	}

	fromAddr := info.GetAddress()
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(fromAddr); err != nil {
		return err.Error()
	}

	to, err := sdk.AccAddressFromBech32(toStr)
	if err != nil {
		return err.Error()
	}

	// parse coins trying to be sent
	coins, err := sdk.ParseCoins(coinStr)
	if err != nil {
		return err.Error()
	}

	account, err := cliCtx.GetAccount(fromAddr)
	if err != nil {
		return err.Error()
	}

	// ensure account has enough coins
	if !account.GetCoins().IsAllGTE(coins) {
		return fmt.Sprintf("Address %s doesn't have enough coins to pay for this transaction.", fromAddr)
	}

	// build and sign the transaction, then broadcast to Tendermint
	msg := bank.NewMsgSend(fromAddr, to, coins)

	//init a txBuilder for the transaction with fee
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(fromAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(fromAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)


	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(fromName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}

	// broadcast to a Tendermint node with Aysc!!!
	res, err := cliCtx.BroadcastTxAsync(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}

func QueryTx(rootDir,Node,chainID,Txhash string) string {
	cliCtx := newCLIContext(rootDir,Node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	hash, err := hex.DecodeString(Txhash)
	if err != nil {
		return err.Error()
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return err.Error()
	}

	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return err.Error()
	}

	//get the resBlocks
	resTxs:= []*ctypes.ResultTx{resTx}
	resBlocks := make(map[int64]*ctypes.ResultBlock)
	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(&resTx.Height)
			if err != nil {
				return err.Error()
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	//parse Tx
	var tx auth.StdTx
	errz := cdc.UnmarshalBinaryLengthPrefixed(resTx.Tx, &tx)
	if errz != nil {
		return errz.Error()
	}


	//format Tx result
	info := sdk.NewResponseResultTx(resTx, tx, resBlocks[resTx.Height].Block.Time.Format(time.RFC3339))

	//json output the result
	resp, _ := cdc.MarshalJSON(info)
	return string(resp)

}