package main

import "C"
import (
	"encoding/json"
	"github.com/QOSGroup/litewallet/qos/account"
	"github.com/QOSGroup/litewallet/qos/wallet"
	db "github.com/tendermint/tm-db"
)

var _qosWallet wallet.Wallet

//export InitWallet
func InitWallet(name, storagePath *C.char) {
	walletDB := db.NewDB(C.GoString(name), db.GoLevelDBBackend, C.GoString(storagePath))
	am := account.NewAccountManager(walletDB)
	_qosWallet = wallet.NewWallet(am)
}

//export ProduceMnemonic
func ProduceMnemonic() *C.char {
	str, _ := json.Marshal(_qosWallet.GenerateMnemonic())
	return C.CString(string(str))
}

//export CreateAccount
func CreateAccount(password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.NewAccount("", C.GoString(password), ""))
	return C.CString(string(str))
}

//export CreateAccountWithName
func CreateAccountWithName(name, password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.NewAccount(C.GoString(name), C.GoString(password), ""))
	return C.CString(string(str))
}

//export CreateAccountWithMnemonic
func CreateAccountWithMnemonic(name, password, mnemonic *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.NewAccount(C.GoString(name), C.GoString(password), C.GoString(mnemonic)))
	return C.CString(string(str))
}

//export GetAccount
func GetAccount(address *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.FindAccount(C.GoString(address)))
	return C.CString(string(str))
}

//export GetAccountByName
func GetAccountByName(name *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.FindAccountByName(C.GoString(name)))
	return C.CString(string(str))
}

//export DeleteAccount
func DeleteAccount(address, password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.DeleteAccount(C.GoString(address), C.GoString(password)))
	return C.CString(string(str))
}


//export ExportAccount
func ExportAccount(address, password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.ProbeAccount(C.GoString(address), C.GoString(password)))
	return C.CString(string(str))
}


//export ImportMnemonic
func ImportMnemonic(mnemonic, password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.RecoverAccountFromMnemonic(C.GoString(mnemonic), C.GoString(password)))
	return C.CString(string(str))
}

//export ImportPrivateKey
func ImportPrivateKey(hexPrivateKey, password *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.RecoverAccountFromPrivateKey(C.GoString(hexPrivateKey), C.GoString(password)))
	return C.CString(string(str))
}

//export ListAllAccounts
func ListAllAccounts() *C.char {
	str, _ := json.Marshal(_qosWallet.ListAllAccounts())
	return C.CString(string(str))
}

//export Sign
func Sign(address, password, signStr *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.SignData(C.GoString(address), C.GoString(password), C.GoString(signStr), ""))
	return C.CString(string(str))
}

//export SignBase64
func SignBase64(address, password, base64Str *C.char) *C.char {
	str, _ := json.Marshal(_qosWallet.SignData(C.GoString(address), C.GoString(password), C.GoString(base64Str), "base64"))
	return C.CString(string(str))
}


func main()  {

}

