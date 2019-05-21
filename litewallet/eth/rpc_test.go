package eth

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	output := GetAccount(node,addr)
	t.Log(output)
}

func TestGetAccountERC20(t *testing.T) {
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0xB1b59AB97D172D0bf46bBC8F8646c4b2fe6B451A"
	tokenAddr := "0xc5d0ac103d253ca6fad4ec3170391ffab6fe5bb8"
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
	amount := int64(200000000000000000)
	gasLimit := uint64(21000)
	output := TransferETH(rootDir,node,name,password,toAddr,amount,gasLimit)
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
	gasLimit := uint64(210000)
	output := TransferERC20(rootDir,node,name,password,toAddr,tokenAddr,gasLimit)
	t.Log(output)
}