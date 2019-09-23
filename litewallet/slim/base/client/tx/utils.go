package tx

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	rpcclient "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/client"
	qtxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"runtime/debug"
)

const (
	MaxGas = 2000000
)

type ITxBuilder func() (txs.ITx, error)

func BuildAndSignTx(ctx context.CLIContext, privkey, chainId string, txBuilder ITxBuilder) (signedTx []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			log := fmt.Sprintf("buildAndSignTx recovered: %v\n", string(debug.Stack()))
			signedTx = nil
			err = errors.New(log)
		}
	}()

	itx, err := txBuilder()
	if err != nil {
		return nil, err
	}
	//toChainID := getChainID(ctx)
	//qcpMode := viper.GetBool(cflags.FlagQcp)
	//if qcpMode {
	//	fromChainID := viper.GetString(cflags.FlagQcpFrom)
	//	return BuildAndSignQcpTx(ctx, itx, fromChainID, toChainID)
	//} else {
	//	return BuildAndSignStdTx(ctx, []txs.ITx{itx}, "", toChainID)
	//}
	msg, err := BuildAndSignStdTx(ctx, itx, privkey, chainId)
	if err != nil {
		return nil, err
	}
	jmsg, _ := ctx.Codec.MarshalJSON(msg)
	fmt.Println(string(jmsg))

	jsonPayload, err := ctx.Codec.MarshalBinaryBare(msg)
	if err != nil {
		return nil, err
	}
	return jsonPayload, nil
}

func BuildAndSignStdTx(ctx context.CLIContext, tx txs.ITx, privkey, chainId string) (*txs.TxStd, error) {
	gas := types.NewInt(int64(MaxGas))
	txStd := txs.NewTxStd(tx, chainId, gas)

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err := ctx.Codec.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())
	from, err := account.GetAddrFromValue(addrben32)
	if err != nil {
		return nil, err
	}

	//Get "nonce" from the func RpcQueryAccount
	acc, _ := qtxs.QueryAccount(from)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	//var actualNonce int64
	//nonce, err := getDefaultAccountNonce(ctx, from.Bytes())
	//if err != nil || nonce < 0 {
	//	return nil, err
	//}
	//actualNonce = nonce + 1

	return signStdTx(key, qscnonce, txStd, chainId, ""), nil
}

func signStdTx(priKey ed25519local.PrivKeyEd25519, nonce int64, txStd *txs.TxStd, chainid string, fromChainID string) *txs.TxStd {
	//gas := NewBigInt(int64(MaxGas))
	//stx := txs.NewTxStd(sendTx, chainid, gas)

	signature, _ := txStd.SignTx(priKey, nonce, "", chainid)
	//sigdata := txStd.BuildSignatureBytes(nonce, fromChainID)
	//sig, pubkey := signData(ctx, signerKeyName, sigdata)
	txStd.Signature = append(txStd.Signature, txs.Signature{
		Pubkey:    priKey.PubKey(),
		Signature: signature,
		Nonce:     nonce,
	})

	return txStd
}

func getDefaultAccountNonce(ctx context.CLIContext, address []byte) (int64, error) {
	if ctx.NonceNodeURI == "" {
		return account.GetAccountNonce(ctx, address)
	}

	//NonceNodeURI不为空,使用NonceNodeURI查询account nonce值
	rpc := rpcclient.NewHTTP(ctx.NonceNodeURI, "/websocket")
	newCtx := context.NewCLIContext("").WithClient(rpc).WithCodec(ctx.Codec)

	return account.GetAccountNonce(newCtx, address)
}
