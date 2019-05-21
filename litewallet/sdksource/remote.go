package sdksource

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"github.com/cosmos/cosmos-sdk/codec"
	cskeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distritypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/bech32"
	"github.com/tendermint/tendermint/libs/cli"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"os"
	"time"
)

var cdc = app.MakeCodec()
const (
	storeStake = "staking"
	//storeDistri = "distr"
	)
//get account from /auth/accounts/{address}
func GetAccount(rootDir,node,chainID,addr string) string {
	key, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	//cliCtx := context.NewCLIContext().
	//	WithCodec(cdc).WithAccountDecoder(cdc)

	if err = cliCtx.EnsureAccountExistsFromAddr(key); err != nil {
		return err.Error()
	}

	acc, err := cliCtx.GetAccount(key)
	if err != nil {
		return err.Error()
	}

	var output []byte
	if cliCtx.Indent {
		output, err = cdc.MarshalJSONIndent(acc, "", "  ")
	} else {
		output, err = cdc.MarshalJSON(acc)
	}
	if err != nil {
		return err.Error()
	}

	return string(output)

}

//complete the whole process with following sequence {Send coins (build -> sign -> send)}
func Transfer(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr, broadcastMode string) string {
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
		WithAccountDecoder(cdc).WithTrustNode(true).WithBroadcastMode(broadcastMode)
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
	//txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithGasPrices(feeStr).WithChainID(chainID)

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
	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	//tmhash to fetch the txhash before broadcast to chain
	//txhash := tmhash.Sum(txBytes)
	//fmt.Println(strings.ToUpper(hex.EncodeToString(txhash)))
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}

//do Delegate operation
func Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, broadcastMode string) string  {
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}
	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	//init a context for this delegate tx
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true).WithBroadcastMode(broadcastMode)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	// parse coin from the delegation
	Delegation, err := sdk.ParseCoin(delegationCoinStr)
	if err != nil {
		return err.Error()
	}

	//check out the account enough money for the delegation
	account, err := cliCtx.GetAccount(DelegatorAddr)
	if err != nil {
		return err.Error()
	}

	DelegationToS := sdk.Coins{Delegation}
	if !account.GetCoins().IsAllGTE(DelegationToS) {
		return fmt.Sprintf("Delegator address %s doesn't have enough coins to perform this transaction.", delegatorAddr)
	}

	//build the stake message
	msg := staking.NewMsgDelegate(DelegatorAddr, ValidatorAddr, Delegation)
	err = msg.ValidateBasic()
	if err != nil {
		return err.Error()
	}

	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}
//get the delegation share under a specific validator
func GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//convert the validator string address to sdk form
	ValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	key := staking.GetDelegationKey(DelAddr, ValAddr)
	res, err := cliCtx.QueryStore(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegation
	delegation, err := types.UnmarshalDelegation(cdc, res)
	if err != nil {
		return err.Error()
	}

	//json output the result
	output, err := codec.MarshalJSONIndent(cdc, delegation)
	if err != nil {
		return err.Error()
	}

	return string(output)

}


//for unbond some of delegation shares from specific validator
func UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, Ubdshares, feeStr, broadcastMode string) string {
	//build procedure
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}

	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	////to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true).WithBroadcastMode(broadcastMode)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	// make a query to get the existing delegation shares
	//key := staking.GetDelegationKey(DelegatorAddr, ValidatorAddr)
	//res, err := cliCtx.QueryStore(key, storeStake)
	//if err != nil {
	//	return err.Error()
	//}
	//
	//// parse out the delegation
	//delegation, err := types.UnmarshalDelegation(cdc, res)
	//if err != nil {
	//	return err.Error()
	//}

	//create the unbond message
	//sharesAmount := delegation.Shares
	sharesAmount, err := sdk.ParseCoin(Ubdshares)
	if err != nil {
		return err.Error()
	}
	msg := staking.NewMsgUndelegate(DelegatorAddr, ValidatorAddr, sharesAmount)

	//build-->sign-->broadcast
	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, []sdk.Msg{msg})
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	resb, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(resb)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}

//get all unbonding delegations from a specific delegator
func GetAllUnbondingDelegations (rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}


	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(staking.GetUBDsKey(DelAddr), storeStake)
	if err != nil {
		return err.Error()
	}

	var ubds staking.UnbondingDelegations
	for _, kv := range resKVs {
		ubds = append(ubds, types.MustUnmarshalUBD(cdc, kv.Value))
	}

	//json output the result
	output, err := codec.MarshalJSONIndent(cdc, ubds)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

//Get bonded validators
func GetBondValidators(rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//generate paras for next query
	params := staking.NewQueryDelegatorParams(DelAddr)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	//query with data
	valids, err := cliCtx.QueryWithData("custom/staking/delegatorValidators", bz)
	//return specific info if there is no delegation between them
	//fmt.Println(valids)
	if len(valids) <= 2 {
		return fmt.Sprintf("None of validators delegated!")
	}
	if err != nil {
		return err.Error()
	}

	var validators []staking.Validator
	if err := cdc.UnmarshalJSON(valids, &validators); err != nil {
		return err.Error()
	}

	var validplus []ValidPlus
	for _, valid := range validators{
		valAddr := valid.GetOperator()
		bz := valAddr.Bytes()
		//var accAddr sdk.AccAddress
		accAddr := sdk.AccAddress(bz)
		//cdc.MustUnmarshalJSON(bz,&accAddr)
		// make a query to get the existing delegation shares
		key := staking.GetDelegationKey(accAddr, valAddr)
		res, err := cliCtx.QueryStore(key, storeStake)
		if err != nil {
			return err.Error()
		}

		var validp ValidPlus
		// parse out the delegation
		delegation, err := types.UnmarshalDelegation(cdc, res)
		if err != nil {
			sharesAmount := "0"
			validp = ValidPlus{
				valid,
				sharesAmount,
			}
		} else {
			sharesAmount := delegation.Shares.String()
			validp = ValidPlus{
				valid,
				sharesAmount,
			}

		}
		validplus = append(validplus,validp)
	}

	output, err := cdc.MarshalJSON(validplus)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

type ValidPlus struct {
	Validator       staking.Validator  `json:"validator"`
	SelfBondShares  string			   `json:"selfbond_shares"`
}

//get all the validators
func GetAllValidators(rootDir, node, chainID string) string {
	key := staking.ValidatorsKey
	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(key, storeStake)
	if err != nil {
		return err.Error()
	}

	var validplus []ValidPlus
	for _, kv := range resKVs {
		//validators = append(validators, types.MustUnmarshalValidator(cdc, kv.Value))

		//fetch the validator info from the key

		valid := types.MustUnmarshalValidator(cdc, kv.Value)
		valAddr := valid.OperatorAddress
		bz := valAddr.Bytes()
		//var accAddr sdk.AccAddress
		accAddr := sdk.AccAddress(bz)
		//cdc.MustUnmarshalJSON(bz,&accAddr)
		// make a query to get the existing delegation shares
		key := staking.GetDelegationKey(accAddr, valAddr)
		res, err := cliCtx.QueryStore(key, storeStake)
		if err != nil {
			return err.Error()
		}

		var validp ValidPlus
		// parse out the delegation
		delegation, err := types.UnmarshalDelegation(cdc, res)
		if err != nil {
			sharesAmount := "0"
			validp = ValidPlus{
				valid,
				sharesAmount,
			}
		} else {
			sharesAmount := delegation.Shares.String()
			validp = ValidPlus{
				valid,
				sharesAmount,
			}

		}
		//add the checkout for tendermint power more than 1
		if validp.Validator.Tokens.GTE(sdk.NewInt(int64(1000000))) {
			validplus = append(validplus,validp)
		}

	}

	output, err := cdc.MarshalJSON(validplus)
	if err != nil {
		return err.Error()
	}
	return string(output)
}

//get all delegations from the delegator
func GetAllDelegations(rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	key := staking.GetDelegationsKey(DelAddr)
	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).WithTrustNode(true)

	resKVs, err := cliCtx.QuerySubspace(key, storeStake)
	if err != nil {
		return err.Error()
	}

	// parse out the delegations
	var delegations staking.Delegations
	for _, kv := range resKVs {
		delegations = append(delegations, types.MustUnmarshalDelegation(cdc, kv.Value))
	}

	output, err := codec.MarshalJSONIndent(cdc, delegations)
	if err != nil {
		return err.Error()
	}

	return string(output)
}

//Withdraw rewards from a specific validator
func WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr, broadcastMode string) string {
	//build procedure
	//get the Keybase
	viper.Set(cli.HomeFlag, rootDir)
	kb, err1 := keys.NewKeyBaseFromHomeFlag()
	if err1 != nil {
		fmt.Println(err1)
	}

	//delegatorName generated from keyspace locally
	if delegatorName == "" {
		fmt.Println("no delegatorName input!")
	}
	info, err := kb.Get(delegatorName)
	if err != nil {
		return err.Error()
	}
	//checkout with rule of own deligation
	DelegatorAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}
	if !bytes.Equal(info.GetPubKey().Address(), DelegatorAddr) {
		return fmt.Sprintf("Must use own delegator address")
	}

	////to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true).WithBroadcastMode(broadcastMode)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelegatorAddr); err != nil {
		return err.Error()
	}

	//validator to address type []byte
	ValidatorAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//generate messages betweeb delegator and validator
	msgs := []sdk.Msg{distritypes.NewMsgWithdrawDelegatorReward(DelegatorAddr, ValidatorAddr)}

	//build-->sign-->broadcast
	//sign the stake message
	//init the txbldr
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithFees(feeStr).WithChainID(chainID)

	//accNum added to txBldr
	accNum, err := cliCtx.GetAccountNumber(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithAccountNumber(accNum)

	//accSequence added
	accSeq, err := cliCtx.GetAccountSequence(DelegatorAddr)
	if err != nil {
		return err.Error()
	}
	txBldr = txBldr.WithSequence(accSeq)

	// build and sign the transaction
	txBytes, err := txBldr.BuildAndSign(delegatorName, password, msgs)
	if err != nil {
		return err.Error()
	}
	// broadcast to a Tendermint node
	resb, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(resb)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)

}

//get a delegation reward between delegator and validator
func GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//convert the validator string address to sdk form
	ValAddr, err := sdk.ValAddressFromBech32(validatorAddr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	//query the delegation rewards
	resp, err := cliCtx.QueryWithData("custom/distr/delegation_rewards", cdc.MustMarshalJSON(distr.NewQueryDelegationRewardsParams(DelAddr, ValAddr)))
	if err != nil {
		return err.Error()
	}

	var result sdk.DecCoins
	cdc.MustUnmarshalJSON(resp, &result)

	resbyte, err := cdc.MarshalJSON(result)
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

//get validator self bond shares
func GetValSelfBondShares (rootDir, node, chainID, validatorAddr string) string {
	//get the delegator string address from validatorAddr as self delegation
	_, valb, _ := bech32.DecodeAndConvert(validatorAddr)
	delegatorAddr, _ := bech32.ConvertAndEncode("cosmos", valb)
	return GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr)
}
//rewardcoins type sdk.Coins
type Delrewards struct {
	RewardsCoins 	sdk.DecCoins		`json:"rewards_coins"`
	Shares			sdk.Dec				`json:"delegation_shares"`
	ValidatorAddr 	sdk.ValAddress		`json:"validator_addr"`
}

//get all the delegation awards list including delegation ties
func GetDelegtorRewardsShares(rootDir, node, chainID, delegatorAddr string) string {
	//convert the delegator string address to sdk form
	DelAddr, err := sdk.AccAddressFromBech32(delegatorAddr)
	if err != nil {
		return err.Error()
	}

	//to be fixed, the trust-node was set true to passby the verifier function, need improvement
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true)
	if err := cliCtx.EnsureAccountExistsFromAddr(DelAddr); err != nil {
		return err.Error()
	}

	//get all the validators with delegation of the specific delegator
	ValAddrs, err := cliCtx.QueryWithData("custom/distr/delegator_validators", cdc.MustMarshalJSON(distr.NewQueryDelegatorParams(DelAddr)))
	if err != nil {
		return err.Error()
	}
	var validators []sdk.ValAddress
	if err := cdc.UnmarshalJSON(ValAddrs, &validators); err != nil {
		return err.Error()
	}

	var delrews []Delrewards
	//query the delegation rewards
	for _, valAddr := range validators{
		rewards, err := cliCtx.QueryWithData("custom/distr/delegation_rewards", cdc.MustMarshalJSON(distr.NewQueryDelegationRewardsParams(DelAddr, valAddr)))
		if err != nil {
			return err.Error()
		}
		var rewardsresult sdk.DecCoins
		cdc.MustUnmarshalJSON(rewards, &rewardsresult)

		// make a query to get the existing delegation shares
		key := staking.GetDelegationKey(DelAddr, valAddr)
		res, err := cliCtx.QueryStore(key, storeStake)
		if err != nil {
			return err.Error()
		}

		// parse out the delegation
		delegation, err := types.UnmarshalDelegation(cdc, res)
		if err != nil {
			return err.Error()
		}

		//create the unbond message
		sharesAmount := delegation.Shares

		delrew := Delrewards{
			rewardsresult,
			sharesAmount,
			valAddr,
		}

		delrews = append(delrews,delrew)

	}
	respbyte, err := cdc.MarshalJSON(delrews)
	if err != nil {
		return err.Error()
	}
	return string(respbyte)

}

//Only partial process with following sequence {Send coins (build -> sign -> Not send)}
func TransferB4send(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr string) string {
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
	//txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc)).WithGasPrices(feeStr).WithChainID(chainID)

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

	return string(hex.EncodeToString(txBytes))
}

//broadcast the tx
func BroadcastTransferTx(rootDir, node, chainID, txString, broadcastMode string) string {
	//initiate context
	cliCtx := newCLIContext(rootDir,node,chainID).
		WithCodec(cdc).
		WithAccountDecoder(cdc).WithTrustNode(true).WithBroadcastMode(broadcastMode)
	// broadcast to a Tendermint node
	txBytes, err := hex.DecodeString(txString)
	if err != nil {
		return err.Error()
	}
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err.Error()
	}
	resbyte, err := cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}