package eth

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0x189f91780c97ed13c242ce3f1568f97a446cab88"
	output := GetAccount(node,addr)
	t.Log(output)
}

func TestGetAccountERC20(t *testing.T) {
	node := "https://mainnet.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	tokenAddr := "0x7b188A8b3A2113621895Fb35fC67a779CAFFA92D"
	output := GetAccountERC20(node,addr,tokenAddr)
	t.Log(output)
}

func TestTransferETH(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	name := "easyzone"
	password := "wm131421"
	toAddr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	amount := "200000000000000000"
	gasLimit := int64(21000)
	gasPrice := "3"
	output := TransferETH(rootDir,node,name,password,toAddr,gasPrice,amount,gasLimit)
	t.Log(output)
}

func TestTransferERC20(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	name := "eth"
	password := "wm131421"
	toAddr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	tokenAddr := "0xc5d0ac103d253ca6fad4ec3170391ffab6fe5bb8"
	//amount := int64(200000000000000000)
	gasLimit := int64(210000)
	tokenValue := "0.34"
	gasPrice := "3"
	output := TransferERC20(rootDir,node,name,password,toAddr,tokenAddr,tokenValue,gasPrice,gasLimit)
	t.Log(output)
}

//func TestFormatFloat(t *testing.T) {
//	tokenValue := "0.32"
//	vamount, err := strconv.ParseFloat(tokenValue,32)
//	if err != nil {
//		log.Fatal(err)
//	}
//	vwei := vamount*1000000000000000000
//	vstring := strconv.FormatFloat(vwei, 'f', -1, 32)
//	//t.Log(vstring)
//	Tamount := new(big.Int)
//	//1000 token to transfer
//	Tamount.SetString(vstring,10)
//	t.Log(Tamount)
//}