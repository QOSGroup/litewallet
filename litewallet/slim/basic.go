package slim

import (
	"fmt"
	btxs "github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/bank/client"
	txs2 "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bip39local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/respwrap"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"log"
)

type ResultCreateAccount struct {
	PubKey   string `json:"pubKey"`
	PrivKey  string `json:"privKey"`
	Addr     string `json:"addr"`
	Mnemonic string `json:"mnemonic"`
	Type     string `json:"type"`
	Denom    string `json:"denom"`
}

type PrivkeyAmino struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PubkeyAmino struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

const (
	// Bech32 prefixes
	//Bech32PrefixAccPub = "cosmosaccpub"
	AccountResultType = "local"
	DenomQOS          = "qos"
	PREF_ADD          = "address"
)

func AccountCreate(password string) *ResultCreateAccount {
	entropy, _ := bip39local.NewEntropy(256)
	mnemonic, _ := bip39local.NewMnemonic(entropy)
	if len(password) == 0 {
		password = "DNWTTY"
	}
	seedo := bip39local.NewSeed(mnemonic, password)

	key := ed25519local.GenPrivKeyFromSecret(seedo)
	//pub := key.PubKey().Bytes()
	pub := key.PubKey()
	pubkeyAmino, _ := txs.Cdc.MarshalJSON(pub)
	var pubkeyAminoStc PubkeyAmino
	err := txs.Cdc.UnmarshalJSON(pubkeyAmino, &pubkeyAminoStc)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pubkeyAminoStr := pubkeyAminoStc.Value

	addr := key.PubKey().Address()
	//bech32Pub, _ := bech32local.ConvertAndEncode(Bech32PrefixAccPub, pub)
	bech32Addr, _ := bech32local.ConvertAndEncode(PREF_ADD, addr.Bytes())

	privkeyAmino, _ := txs.Cdc.MarshalJSON(key)
	var privkeyAminoStc PrivkeyAmino
	err1 := txs.Cdc.UnmarshalJSON(privkeyAmino, &privkeyAminoStc)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	privkeyAminoStr := privkeyAminoStc.Value

	Type := AccountResultType

	result := &ResultCreateAccount{}
	result.PubKey = pubkeyAminoStr
	result.PrivKey = privkeyAminoStr
	result.Addr = bech32Addr
	result.Mnemonic = mnemonic
	result.Type = Type
	result.Denom = DenomQOS

	return result
}

//convert the output to json string format
func AccountCreateStr(password string) string {
	acc := AccountCreate(password)
	result, _ := respwrap.ResponseWrapper(txs.Cdc, acc, nil)
	out := string(result)

	return out
}

func AccountRecoverStr(mncode, password string) string {
	if len(password) == 0 {
		password = "DNWTTY"
	}
	// add mnemonics validation
	if bip39local.IsMnemonicValid(mncode) == false {
		err := errors.Errorf("Invalid mnemonic!")
		resp, _ := respwrap.ResponseWrapper(txs.Cdc, nil, err)
		return string(resp)

	}

	seed := bip39local.NewSeed(mncode, password)
	key := ed25519local.GenPrivKeyFromSecret(seed)
	pub := key.PubKey()
	pubkeyAmino, _ := txs.Cdc.MarshalJSON(pub)
	var pubkeyAminoStc PubkeyAmino
	err := txs.Cdc.UnmarshalJSON(pubkeyAmino, &pubkeyAminoStc)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pubkeyAminoStr := pubkeyAminoStc.Value

	addr := key.PubKey().Address()
	//bech32Pub, _ := bech32local.ConvertAndEncode("cosmosaccpub", pub)
	bech32Addr, _ := bech32local.ConvertAndEncode(PREF_ADD, addr.Bytes())

	privkeyAmino, _ := txs.Cdc.MarshalJSON(key)
	var privkeyAminoStc PrivkeyAmino
	err1 := txs.Cdc.UnmarshalJSON(privkeyAmino, &privkeyAminoStc)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	privkeyAminoStr := privkeyAminoStc.Value

	Type := AccountResultType
	result := &ResultCreateAccount{}
	result.PubKey = pubkeyAminoStr
	result.PrivKey = privkeyAminoStr
	result.Addr = bech32Addr
	result.Mnemonic = mncode
	result.Type = Type
	result.Denom = DenomQOS

	resp, _ := respwrap.ResponseWrapper(txs.Cdc, result, nil)
	out := string(resp)
	return out
}

type PubAddrRetrieval struct {
	PubKey string `json:"pubKey"`
	Addr   string `json:"addr"`
}

func PubAddrRetrievalStr(s string) string {
	//change the private unmarshal format according to the other pack
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + s + "\"}"
	var key ed25519local.PrivKeyEd25519

	err := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	pub := key.PubKey()
	pubkeyAmino, _ := txs.Cdc.MarshalJSON(pub)
	var pubkeyAminoStc PubkeyAmino
	err1 := txs.Cdc.UnmarshalJSON(pubkeyAmino, &pubkeyAminoStc)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	pubkeyAminoStr := pubkeyAminoStc.Value

	addr := key.PubKey().Address()
	//bech32Pub, _ := bech32local.ConvertAndEncode(Bech32PrefixAccPub, pub)
	bech32Addr, _ := bech32local.ConvertAndEncode(PREF_ADD, addr.Bytes())

	result := &PubAddrRetrieval{}
	result.PubKey = pubkeyAminoStr
	result.Addr = bech32Addr

	resp, _ := respwrap.ResponseWrapper(txs.Cdc, result, nil)
	out := string(resp)
	return out
}

//new account result with field of Denom
type AccountKeyOut struct {
	PubKey   string `json:"pubKey"`
	PrivKey  string `json:"privKey"`
	Addr     string `json:"addr"`
	Mnemonic string `json:"mnemonic"`
	Type     string `json:"type"`
	Denom    string `json:"denom"`
}

//add new function for Account Creation with seed input
func AccountCreateFromSeed(mncode string) string {
	// add mnemonics validation
	if bip39local.IsMnemonicValid(mncode) == false {
		err := errors.Errorf("Invalid mnemonic!")
		resp, _ := respwrap.ResponseWrapper(txs.Cdc, nil, err)
		return string(resp)

	}

	var defaultBIP39Passphrase = ""
	seed := bip39local.NewSeed(mncode, defaultBIP39Passphrase)
	key := ed25519local.GenPrivKeyFromSecret(seed)
	pub := key.PubKey()
	pubkeyAmino, _ := txs.Cdc.MarshalJSON(pub)
	var pubkeyAminoStc PubkeyAmino
	err := txs.Cdc.UnmarshalJSON(pubkeyAmino, &pubkeyAminoStc)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pubkeyAminoStr := pubkeyAminoStc.Value

	addr := key.PubKey().Address()
	//bech32Pub, _ := bech32local.ConvertAndEncode("cosmosaccpub", pub)
	bech32Addr, _ := bech32local.ConvertAndEncode(PREF_ADD, addr.Bytes())

	privkeyAmino, _ := txs.Cdc.MarshalJSON(key)
	var privkeyAminoStc PrivkeyAmino
	err1 := txs.Cdc.UnmarshalJSON(privkeyAmino, &privkeyAminoStc)
	if err1 != nil {
		log.Fatalln(err1.Error())
	}
	privkeyAminoStr := privkeyAminoStc.Value

	Type := AccountResultType
	Denom := DenomQOS

	result := &AccountKeyOut{}
	result.PubKey = pubkeyAminoStr
	result.PrivKey = privkeyAminoStr
	result.Addr = bech32Addr
	result.Mnemonic = mncode
	result.Type = Type
	result.Denom = Denom

	resp, _ := respwrap.ResponseWrapper(txs.Cdc, result, nil)
	out := string(resp)
	return out

}

//Local Tx generation
func LocalTxGen(fromStr, toStr, coinstr, chainid, privkey string, nonce int64) []byte {
	sendersStr := fromStr + `,` + coinstr
	senders, err := client.ParseTransItem(sendersStr)
	if err != nil {
		err.Error()
	}

	receiversStr := toStr + `,` + coinstr
	receivers, err := client.ParseTransItem(receiversStr)
	if err != nil {
		err.Error()
	}

	tn := txs2.TxTransfer{
		Senders:   senders,
		Receivers: receivers,
	}

	gas := types.NewInt(int64(20000))
	txStd := btxs.NewTxStd(tn, chainid, gas)

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err1 := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := crypto.PrivKey(key)

	signature, _ := txStd.SignTx(priv, nonce, chainid)
	txStd.Signature = []btxs.Signature{btxs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature,
		Nonce:     nonce,
	}}
	msg := txStd
	jasonpayload, err := txs.Cdc.MarshalBinaryBare(msg)
	if err != nil {
		fmt.Println(err)
	}

	return jasonpayload
}
