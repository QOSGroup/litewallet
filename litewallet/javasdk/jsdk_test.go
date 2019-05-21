package javasdk

import (
	"os/user"
	"testing"
)

func TestTransferAsync(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	fromName := "local"
	password := "wm131421"
	toStr := "cosmos1mrf49r22adtd8juv6kvg8dxly32qlj7rg47644"
	coinStr := "1stake"
	feeStr := "1token"
	transout := TransferAsync(rootDir,node,chainId,fromName,password,toStr,coinStr,feeStr)
	t.Log(transout)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://localhost:26657"
	chainId := "test4matt"
	Txhash := "0BA029449967228DB14E7ECCFF9B97C5963807DB07D32CF180CBD545BBE59CFC"
	qout := QueryTx(rootDir, node, chainId, Txhash)
	t.Log(qout)
}