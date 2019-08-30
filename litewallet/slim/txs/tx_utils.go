package txs

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/pkg/errors"
)

const (
	PREF_ADD = "address"
	MaxGas   = 20000
)

type ITxBuilder func() (txs.ITx, error)

func BuildAndSignTx(privkey, chainid string, txBuilder ITxBuilder) (signedTx string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log := fmt.Sprintf("buildAndSignTx recovered: %v\n", r)
			signedTx = ""
			err = errors.New(log)
		}
	}()

	itx, err := txBuilder()
	if err != nil {
		return "", err
	}
	//toChainID := getChainID(ctx)
	//qcpMode := viper.GetBool(cflags.FlagQcp)
	//if qcpMode {
	//	fromChainID := viper.GetString(cflags.FlagQcpFrom)
	//	return BuildAndSignQcpTx(ctx, itx, fromChainID, toChainID)
	//} else {
	//	return BuildAndSignStdTx(ctx, []txs.ITx{itx}, "", toChainID)
	//}
	msg := BuildAndSignTxStd(itx, privkey, chainid)
	jasonpayload, err := Cdc.MarshalJSON(msg)
	if err != nil {
		return "", err
	}
	return string(jasonpayload), nil
}

func BuildAndSignTxStd(tx txs.ITx, privkey, chainid string) *txs.TxStd {
	gas := types.NewInt(int64(MaxGas))
	txStd := txs.NewTxStd(tx, chainid, gas)

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	from, err2 := types.GetAddrFromBech32(addrben32)
	if err2 != nil {
		fmt.Println(err2)
	}
	//Get "nonce" from the func RpcQueryAccount
	acc, _ := RpcQueryAccount(from)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	return signTxStd(txStd, key, chainid, qscnonce)
}

func signTxStd(txStd *txs.TxStd, priKey ed25519local.PrivKeyEd25519, chainid string, nonce int64) *txs.TxStd {
	//gas := NewBigInt(int64(MaxGas))
	//stx := txs.NewTxStd(sendTx, chainid, gas)

	signature, _ := txStd.SignTx(priKey, nonce, chainid)
	txStd.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priKey.PubKey(),
		Signature: signature,
		Nonce:     nonce,
	}}

	return txStd
}
