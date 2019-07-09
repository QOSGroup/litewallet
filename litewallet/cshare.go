package litewallet

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/QOSGroup/litewallet/litewallet/eth"
	"github.com/QOSGroup/litewallet/litewallet/sdksource"
	"github.com/QOSGroup/litewallet/litewallet/slim"
)

//create the seed(mnemonic) for the account generation
func CreateSeed() string {
	output := sdksource.CreateSeed()
	return output
}

//WalletAddressCheck for different chains
func WalletAddressCheck(addr string) string {
	output := sdksource.WalletAddressCheck(addr)
	return output
}

//create account
func CosmosCreateAccount(rootDir, name, password, seed string) string {
	output := sdksource.CreateAccount(rootDir, name, password, seed)
	return output
}

//recover key
func CosmosRecoverKey(rootDir, name, password, seed string) string {
	output := sdksource.RecoverKey(rootDir, name, password, seed)
	return output
}

//update password
func CosmosUpdateKey(rootDir, name, oldpass, newpass string) string {
	output := sdksource.UpdateKey(rootDir, name, oldpass, newpass)
	return output
}

//get account info
func CosmosGetAccount(rootDir, node, chainID, addr string) string {
	output := sdksource.GetAccount(rootDir, node, chainID, addr)
	return output
}

//transfer
func CosmosTransfer(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr, broadcastMode string) string {
	output := sdksource.Transfer(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr, broadcastMode)
	return output
}

//delegate
func CosmosDelegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, broadcastMode string) string {
	output := sdksource.Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, broadcastMode)
	return output
}

//get a specific delegation shares
func CosmosGetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}

//for unbond delegation shares from specific validator
func CosmosUnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, Ubdshares, feeStr, broadcastMode string) string {
	output := sdksource.UnbondingDelegation(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, Ubdshares, feeStr, broadcastMode)
	return output
}

//get all unbonding delegations from a specific delegator
func CosmosGetAllUnbondingDelegations(rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetAllUnbondingDelegations(rootDir, node, chainID, delegatorAddr)
	return output
}

//Get bonded validators
func CosmosGetBondValidators(rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetBondValidators(rootDir, node, chainID, delegatorAddr)
	return output
}

//get all the validators
func CosmosGetAllValidators(rootDir, node, chainID string) string {
	output := sdksource.GetAllValidators(rootDir, node, chainID)
	return output
}

//get all delegations from the delegator
func CosmosGetAllDelegations(rootDir, node, chainID, delegatorAddr string) string {
	output := sdksource.GetAllDelegations(rootDir, node, chainID, delegatorAddr)
	return output
}

//Withdraw rewards from a specific validator
func CosmosWithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr, broadcastMode string) string {
	output := sdksource.WithdrawDelegationReward(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr, feeStr, broadcastMode)
	return output
}

//get a delegation reward between delegator and validator
func CosmosGetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr string) string {
	output := sdksource.GetDelegationRewards(rootDir, node, chainID, delegatorAddr, validatorAddr)
	return output
}

//query the tx result by txHash generated via async broadcast
func CosmosQueryTx(rootDir, node, chainId, txHash string) string {
	output := sdksource.QueryTx(rootDir, node, chainId, txHash)
	return output
}

func CosmosGetValSelfBondShares(rootDir, node, chainID, validatorAddr string) string {
	output := sdksource.GetValSelfBondShares(rootDir, node, chainID, validatorAddr)
	return output
}

func CosmosGetDelegtorRewardsShares(rootDir, node, chainId, delegatorAddr string) string {
	output := sdksource.GetDelegtorRewardsShares(rootDir, node, chainId, delegatorAddr)
	return output
}

func CosmosWithdrawDelegatorAllRewards(rootDir, node, chainID, delegatorName, password, delegatorAddr, feeStr, broadcastMode string) string {
	output := sdksource.WithdrawDelegatorAllRewards(rootDir, node, chainID, delegatorName, password, delegatorAddr, feeStr, broadcastMode)
	return output
}

func CosmosQueryQueryTxsWithTags(rootDir, node, chainID, addr string, page, limit int) string {
	output := sdksource.QueryTxsWithTags(rootDir, node, chainID, addr, page, limit)
	return output
}

//QOS wallet part begin from here
func QOSAccountCreate(password string) string {
	output := slim.AccountCreateStr(password)
	return output
}

func QOSAccountCreateFromSeed(mncode string) string {
	output := slim.AccountCreateFromSeed(mncode)
	return output
}

//for QSCKVStoreset
func QSCKVStoreSet(k, v, privkey, chain string) string {
	output := slim.QSCKVStoreSetPost(k, v, privkey, chain)
	return output
}

//for QSCKVStoreGet
func QSCKVStoreGet(k string) string {
	output := slim.QSCKVStoreGetQuery(k)
	return output
}

//for QSCQueryAccount
func QSCQueryAccount(addr string) string {
	output := slim.QSCQueryAccountGet(addr)
	return output
}

//for QOSQueryAccount
func QOSQueryAccount(addr string) string {
	output := slim.QOSQueryAccountGet(addr)
	return output
}

//for AccountRecovery
func QOSAccountRecover(mncode, password string) string {
	output := slim.AccountRecoverStr(mncode, password)
	return output
}

//for IP input
func QOSSetBlockchainEntrance(sh, mh string) {
	slim.SetBlockchainEntrance(sh, mh)
}

//for PubAddrRetrieval
func QOSPubAddrRetrieval(priv string) string {
	//	fmt.Println("Please input host including IP and port for initialization on Qstar deamon:")
	output := slim.PubAddrRetrievalStr(priv)
	return output
}

//for QSCtransferSend
func QSCtransferSend(addrto, coinstr, privkey, chainid string) string {
	output := slim.QSCtransferSendStr(addrto, coinstr, privkey, chainid)
	return output
}

//for QOSCommitResultCheck
func QOSCommitResultCheck(txhash, height string) string {
	output := slim.QOSCommitResultCheck(txhash, height)
	return output
}

func QOSJQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) string {
	output := slim.JQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey)
	return output
}

func QOSAesEncrypt(key, plainText string) string {
	output := slim.AesEncrypt(key, plainText)
	return output
}

func QOSAesDecrypt(key, cipherText string) string {
	output := slim.AesDecrypt(key, cipherText)
	return output
}

func QOSTransferRecordsQuery(chainid, addr, cointype, offset, limit string) string {
	output := slim.TransferRecordsQuery(chainid, addr, cointype, offset, limit)
	return output
}

func CosmosTransferB4send(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr string) string {
	output := sdksource.TransferB4send(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr)
	return output
}

func CosmosBroadcastTransferTx(rootDir, node, chainID, txString, broadcastMode string) string {
	output := sdksource.BroadcastTransferTx(rootDir, node, chainID, txString, broadcastMode)
	return output
}

//for AdvertisersTrue
func QOSAdvertisersTrue(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := slim.AdvertisersTrue(privatekey, coinsType, coinAmount, qscchainid)
	return output
}

//for AdvertisersFalse
func QOSAdvertisersFalse(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := slim.AdvertisersFalse(privatekey, coinsType, coinAmount, qscchainid)
	return output
}

//for GetTx
func QOSGetTx(tx string) string {
	output := slim.GetTx(tx)
	return output
}

func QOSGetBlance(addrs string) string {
	path := fmt.Sprintf("/store/%s/%s", "aoeaccount", "key")
	output, _ := slim.Query(path, []byte(addrs))
	var basecoin *slim.BaseCoins
	//err=json.Unmarshal(resp.Value,&basecoin)
	slim.Cdc.UnmarshalBinaryBare(output, &basecoin)
	result, _ := json.Marshal(basecoin)
	return string(result)
}

func QOSGetBlanceByCointype(addrs, cointype string) string {
	result := QOSGetBlance(addrs)
	var qsc slim.QSCs
	json.Unmarshal([]byte(result), &qsc)
	for _, v := range qsc {
		if strings.ToUpper(v.Name) == strings.ToUpper(cointype) {
			return v.Amount.String()
		}
	}
	return "0"
}

// acutionAd 竞拍广告
//articleHash            //广告位标识
//privatekey             //用户私钥
//coinsType              //竞拍币种
//coinAmount             //竞拍数量
//qscchainid             //chainid
func QOSAcutionAd(articleHash, privatekey, coinsType, coinAmount, qscchainid string) string {
	output := slim.AcutionAd(articleHash, privatekey, coinsType, coinAmount, qscchainid)
	return output
}

//for Extract
func QOSExtract(privatekey, coinsType, coinAmount, qscchainid string) string {
	output := slim.Extract(privatekey, coinsType, coinAmount, qscchainid)
	return output
}

// 提交到联盟链上
func QOSBroadcastTransferTxToQSC(txstring, broadcastModes string) string {
	return slim.BroadcastTransferTxToQSC(txstring, broadcastModes)
}

func QOSCommHandler(funcName, privatekey, args, qscchainid string) string {
	output := slim.CommHandler(funcName, privatekey, args, qscchainid)
	return output
}

//From here, Eth wallet part start
func EthCreateAccount(rootDir, name, password, seed string) string {
	output := eth.CreateAccount(rootDir, name, password, seed)
	return output
}

func EthRecoverAccount(rootDir, name, password, seed string) string {
	output := eth.RecoverAccount(rootDir, name, password, seed)
	return output
}

func EthGetAccount(node, addr string) string {
	output := eth.GetAccount(node, addr)
	return output
}

func EthGetErc20Account(node, addr, tokenAddr string) string {
	output := eth.GetAccountERC20(node, addr, tokenAddr)
	return output
}

func EthTransferETH(rootDir, node, name, password, toAddr, gasPrice, amount string, gasLimit int64) string {
	output := eth.TransferETH(rootDir, node, name, password, toAddr, gasPrice, amount, gasLimit)
	return output
}

func EthTransferErc20(rootDir, node, name, password, toAddr, tokenAddr, tokenValue, gasPrice string, gasLimit int64) string {
	output := eth.TransferERC20(rootDir, node, name, password, toAddr, tokenAddr, tokenValue, gasPrice, gasLimit)
	return output
}

//Deprecated!
//func EthGetPendingNonceAt(rootDir, node, fromName, password string) int64 {
//	output := eth.GetPendingNonceAt(rootDir, node, fromName, password)
//	return output
//}

func EthSpeedTransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount string, GasLimit, pendingNonce int64) string {
	output := eth.SpeedTransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount, GasLimit, pendingNonce)
	return output
}

func EthSpeedTransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice string, GasLimit, pendingNonce int64) string {
	output := eth.SpeedTransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice, GasLimit, pendingNonce)
	return output
}
