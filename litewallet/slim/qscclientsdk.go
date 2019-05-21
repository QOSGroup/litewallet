package slim

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/slim/funcInlocal/bech32local"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/slim/funcInlocal/ed25519local"
	"log"

)


type JianQianTx struct {
	FuncName string               //方法名 路由用
	Address  []Address            //签名者地址
	Args     []string             //参数列表
	Gas      BigInt
}


func (it JianQianTx) GetSignData() (ret []byte) {
	ret = append(ret, []byte(it.FuncName)...)
	for _, v := range it.Address {
		ret = append(ret, v.Bytes()...)
	}
	for _, v := range it.Args {
		ret = append(ret, []byte(v)...)
	}
	return
}


func CommHandler(funcName, privatekey, argstr,qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess

	var args []string
	err := json.Unmarshal([]byte(argstr), &args)
	if err != nil {
		log.Printf("CommHandler args err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}

	tx, berr := commHandler(funcName, privatekey, args,qscchainid)
	if berr != "" {
		return berr
	}
	js, err := Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("CommHandler err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

func commHandler(funcName, privatekey string, args []string,qscchainid string) (*TxStd, string) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	address, _ := getAddrFromBech32(addrben32)

	acc, _ := RpcQueryAccount(address)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	tx := &JianQianTx{}
	tx.Address = []Address{address}
	tx.FuncName = funcName
	tx.Args = args
	tx.Gas = gas


	tx2 := NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []Signature{Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}
