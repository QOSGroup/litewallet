package client

import (
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"log"
	"strconv"
)

//广告商押金或赎回
func Extract(amount, privatekey, cointype, qscchainid string) string {
	var result ctypes.ResultInvest
	result.Code = ctypes.ResultCodeSuccess
	tx, berr := extract(amount, privatekey, cointype, qscchainid)
	if berr != "" {
		return berr
	}
	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("Extract err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

func extract(coins, privatekey, cointype, qscchainid string) (*txs.TxStd, string) {
	amount, err := strconv.Atoi(coins)
	if err != nil {
		return nil, ctypes.NewErrorResult("601", 0, "", "amount format error").Marshal()
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := types.NewInt(int64(ctxs.MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())
	investor, _ := types.GetAddrFromBech32(addrben32)

	acc, _ := ctxs.QueryAccount(investor)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &ctypes.CoinsTx{}
	it.Address = investor
	it.Cointype = cointype
	it.ChangeType = "2"
	it.Amount = types.NewInt(int64(amount))
	tx := ctxs.ExtractTx{it}
	fmt.Println(investor, amount, cointype, "2")
	tx2 := txs.NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}
