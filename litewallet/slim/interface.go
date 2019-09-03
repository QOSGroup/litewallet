package slim

import (
	"encoding/hex"
	"github.com/QOSGroup/litewallet/litewallet/slim/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"strings"
)

//only need the following arguments, it`s enough!
func QueryAccount(addr string) ([]byte, error) {
	qosAccount, err := client.GetAccountFromBech32Addr(addr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(qosAccount)
}

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

func DelegationSend(addrto string, coins int64, privkey, chainid string) (string, error) {
	tx, err := client.CreateSignedDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func UnbondDelegationSend(addrto string, coins int64, privkey, chainid string) (string, error) {
	tx, err := client.CreateSignedUnbondDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func QueryApproveSend(addrto, privkey string) ([]byte, error) {
	approve, err := client.QueryApprove(addrto, privkey)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(approve)
}

func CreateApproveSend(addrto string, coinsStr string, privkey, chainid string) (string, error) {
	tx, err := client.CreateApprove(addrto, coinsStr, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func IncreaseApprove(addrto string, coinsStr string, privkey, chainid string) (string, error) {
	tx, err := client.IncreaseApprove(addrto, coinsStr, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func DecreaseApproveSend(addrto string, coinsStr string, privkey, chainid string) (string, error) {
	tx, err := client.DecreaseApprove(addrto, coinsStr, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func UseApproveSend(addrto string, coinsStr string, privkey, chainid string) (string, error) {
	tx, err := client.DecreaseApprove(addrto, coinsStr, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func CancelApproveSend(addrto string, coinsStr string, privkey, chainid string) (string, error) {
	tx, err := client.DecreaseApprove(addrto, coinsStr, privkey, chainid)
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
