package slim

import (
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/QOSGroup/litewallet/litewallet/slim/http"
	txs2 "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"log"
)

type JianQianTx struct {
	FuncName string    //方法名 路由用
	Address  []types.Address //签名者地址
	Args     []string  //参数列表
	Gas      types.BigInt
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

func CommHandler(funcName, privatekey, argstr, qscchainid string) string {
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

	tx, berr := commHandler(funcName, privatekey, args, qscchainid)
	if berr != "" {
		return berr
	}
	js, err := txs2.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("CommHandler err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

func commHandler(funcName, privatekey string, args []string, qscchainid string) (*txs.TxStd, string) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := txs2.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	address, _ := types.GetAddrFromBech32(addrben32)

	acc, _ := http.RpcQueryAccount(address)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	tx := &JianQianTx{}
	tx.Address = []types.Address{address}
	tx.FuncName = funcName
	tx.Args = args
	tx.Gas = gas

	tx2 := txs.NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}
