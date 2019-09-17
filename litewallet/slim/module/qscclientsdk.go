package module

import (
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"log"
)

func CommHandler(funcName, privatekey, argstr, qscchainid string) string {
	var result ctypes.ResultInvest
	result.Code = ctypes.ResultCodeSuccess

	var args []string
	err := json.Unmarshal([]byte(argstr), &args)
	if err != nil {
		log.Printf("CommHandler args err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}

	tx, berr := commHandler(funcName, privatekey, args, qscchainid)
	if berr != "" {
		return berr
	}
	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("CommHandler err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

func commHandler(funcName, privatekey string, args []string, qscchainid string) (*txs.TxStd, string) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := types.NewInt(int64(ctxs.MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())
	address, _ := types.GetAddrFromBech32(addrben32)

	acc, _ := ctxs.QueryAccount(address)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	tx := &ctxs.JianQianTx{}
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
