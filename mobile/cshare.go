package litewallet

import (
	"encoding/json"
	"github.com/QOSGroup/litewallet/eth"
	"github.com/QOSGroup/litewallet/qos/account"
	"github.com/QOSGroup/litewallet/qos/wallet"
	db "github.com/tendermint/tm-db"
)


//QOS wallet part
var _qosWallet wallet.Wallet

func InitWallet(name, storagePath string) {
	walletDB := db.NewDB(name, db.GoLevelDBBackend, storagePath)
	am := account.NewAccountManager(walletDB)
	_qosWallet = wallet.NewWallet(am)
}

func ProduceMnemonic() string {
	str, _ := json.Marshal(_qosWallet.GenerateMnemonic())
	return string(str)
}

func CreateAccount(password string) string {
	str, _ := json.Marshal(_qosWallet.NewAccount("", password, ""))
	return string(str)
}

func CreateAccountWithName(name, password string) string {
	str, _ := json.Marshal(_qosWallet.NewAccount(name, password, ""))
	return string(str)
}

func CreateAccountWithMnemonic(name, password, mnemonic string) string {
	str, _ := json.Marshal(_qosWallet.NewAccount(name, password, mnemonic))
	return string(str)
}

func GetAccount(address string) string {
	str, _ := json.Marshal(_qosWallet.FindAccount(address))
	return string(str)
}

func GetAccountByName(name string) string {
	str, _ := json.Marshal(_qosWallet.FindAccountByName(name))
	return string(str)
}

func DeleteAccount(address, password string) string {
	str, _ := json.Marshal(_qosWallet.DeleteAccount(address, password))
	return string(str)
}

func ExportAccount(address, password string) string {
	str, _ := json.Marshal(_qosWallet.ProbeAccount(address, password))
	return string(str)
}

func ImportMnemonic(mnemonic, password string) string {
	str, _ := json.Marshal(_qosWallet.RecoverAccountFromMnemonic(mnemonic, password))
	return string(str)
}

func ImportPrivateKey(hexPrivateKey, password string) string {
	str, _ := json.Marshal(_qosWallet.RecoverAccountFromPrivateKey(hexPrivateKey, password))
	return string(str)
}

func ListAllAccounts() string {
	str, _ := json.Marshal(_qosWallet.ListAllAccounts())
	return string(str)
}

func Sign(address, password, signStr string) string {
	str, _ := json.Marshal(_qosWallet.SignData(address, password, signStr, ""))
	return string(str)
}

func SignBase64(address, password, base64Str string) string {
	str, _ := json.Marshal(_qosWallet.SignData(address, password, base64Str, "base64"))
	return string(str)
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

//EthGetNonceAt provide the nonce at the latest block
func EthGetNonceAt(rootDir, node, fromName, password string) int64 {
	output := eth.GetNonceAt(rootDir, node, fromName, password)
	return output
}
