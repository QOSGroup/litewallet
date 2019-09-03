package slim

import (
	"encoding/hex"
	"github.com/QOSGroup/litewallet/litewallet/slim/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"strings"
)

//only need the following arguments, it`s enough!
func TransferSend(addrto, coinstr, privkey, chainid string) (string, error) {
	tx, err := client.CreateSignedTransfer(addrto, coinstr, privkey, chainid)
	if err != nil {
		return "", err
	}
	//datas := bytes.NewBuffer([]byte(jasonpayload))
	//aurl := txs.Accounturl + "txSend"
	//req, _ := http.NewRequest("POST", aurl, datas)
	//req.Header.Set("Content-Type", "application/json")
	//clt := http.Client{}
	//resp, _ := clt.Do(req)
	//body, err := ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()
	//output := string(body)
	return txs.BroadcastTx(tx)
}

//only need the following arguments, it`s enough!
func DelegationSend(addrto string, coins int64, privkey, chainid string) (string, error) {
	tx, err := client.CreateSignedDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

//only need the following arguments, it`s enough!
func UnbondDelegationSend(addrto string, coins int64, privkey, chainid string)  (string, error) {
	tx, err := client.CreateSignedUnbondDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func GetTx(tx string) string {
	txBytes, err := hex.DecodeString(tx)
	if err != nil {
		return err.Error()
	}
	txhashs := strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
	return string(txhashs)
}
