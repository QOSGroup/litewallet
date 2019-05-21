package main

import "C"
import (
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/javasdk"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/sdksource"
)

//export GetAccount
func GetAccount(rootDir,node,chainId,addr *C.char) *C.char {
	output := sdksource.GetAccount(C.GoString(rootDir),C.GoString(node), C.GoString(chainId), C.GoString(addr))
	return C.CString(output)
}

//export RecoverKey
func RecoverKey(rootDir,name,password,seed *C.char) *C.char {
	output := sdksource.RecoverKey(C.GoString(rootDir),C.GoString(name), C.GoString(password), C.GoString(seed))
	return C.CString(output)
}


//export TransferAsync
func TransferAsync(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr *C.char) *C.char {
	output := javasdk.TransferAsync(C.GoString(rootDir),C.GoString(node), C.GoString(chainId), C.GoString(fromName),C.GoString(password), C.GoString(toStr), C.GoString(coinStr), C.GoString(feeStr))
	return C.CString(output)
}

//export QueryTx
func QueryTx(rootDir, node, chainId, Txhash *C.char) *C.char {
	output := javasdk.QueryTx(C.GoString(rootDir),C.GoString(node), C.GoString(chainId), C.GoString(Txhash))
	return C.CString(output)
}

func main() {

}