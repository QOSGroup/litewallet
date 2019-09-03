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

// acutionAd 竞拍广告
//articleHash            //广告位标识
//privatekey             //用户私钥
//coinsType              //竞拍币种
//coinAmount             //竞拍数量
//qscchainid             //chainid
func AcutionAd(articleHash, privatekey, coinsType, coinAmount, qscchainid string) string {
	var result ctypes.ResultInvest
	result.Code = ctypes.ResultCodeSuccess
	amount, err := strconv.Atoi(coinAmount)
	if err != nil {
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = "AcutionAd invalid amount"
		return result.Marshal()
	}
	tx, buff := acutionAd(articleHash, privatekey, coinsType, amount, qscchainid)
	if buff != "" {
		return buff
	}

	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("AcutionAd err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

// acutionAd 竞拍广告
func acutionAd(articleHash, privatekey, coinsType string, coinAmount int, qscchainid string) (*txs.TxStd, string) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := types.NewInt(int64(ctxs.MaxGas))
	addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())
	sendAddress, _ := types.GetAddrFromBech32(addrben32)
	acc, _ := ctxs.QueryAccount(sendAddress)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &ctxs.AuctionTx{}
	it.ArticleHash = articleHash
	it.Address = sendAddress
	it.CoinsType = coinsType
	it.Gas = types.ZeroInt()
	it.CoinAmount = types.NewInt(int64(coinAmount))
	tx2 := txs.NewTxStd(it, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}
