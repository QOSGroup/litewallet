package txs

import (
	"encoding/hex"
	"errors"
	qtypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	accountStoreKey = "account:" // 便于获取全部账户的通用存储键名，继承BaseAccount时，可根据不同业务设置存储前缀
)

type sendKVReq struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	PrivateKey string `json:"privatekey"`
	ChainID    string `json:"chainid"`
}

// IP initialization
var (
	Shost         string
	Mhost         string
	QSCAccounturl string
	QOSAccounturl string
	Accounturl    string
	KVurl         string
	QResulturl    string
	TRurl         string
	RPC           rpcclient.Client
)

//set Block Chain entrance hosts for both Qstars and Qmoon
func SetBlockchainEntrance(qstarshost, qmoonhost string) {
	Shost = qstarshost
	Mhost = qmoonhost
	QSCAccounturl = "http://" + Shost + "/QSCaccounts/"
	QOSAccounturl = "http://" + Shost + "/QOSaccounts/"
	Accounturl = "http://" + Shost + "/accounts/"
	KVurl = "http://" + Shost + "/kv/"
	QResulturl = "http://" + Shost + "/commits/"
	TRurl = "http://" + Mhost + "/nodes/"

	RPC = rpcclient.NewHTTP(Shost, "/websocket")
}

func init() {
	var sh string
	var mh string
	SetBlockchainEntrance(sh, mh)
}

//func QSCKVStoreSetPost(k, v, privkey, chain string) (result string) {
//	skr := sendKVReq{}
//	skr.Key = k
//	skr.Value = v
//	skr.PrivateKey = privkey
//	skr.ChainID = chain
//	payload, _ := Cdc.MarshalJSON(skr)
//	body := bytes.NewBuffer(payload)
//	req, _ := http.NewRequest("POST", KVurl, body)
//	req.Header.Set("Content-Type", "application/json")
//	//get the origin codes from the standard net/http package
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	respBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	output := string(respBody)
//	//fmt.Println(output)
//	return output
//}

//func QSCKVStoreGetQuery(k string) string {
//	kvurl := KVurl + k
//	resp, _ := http.Get(kvurl)
//	//	fmt.Println(KVurl)
//	if resp.StatusCode == http.StatusOK {
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//			//log.Fatal(err)
//		}
//		defer resp.Body.Close()
//		output := string(body)
//		return output
//	}
//	return "nil"
//}

//func QSCQueryAccountGet(addr string) string {
//	aurl := QSCAccounturl + addr
//	resp, _ := http.Get(aurl)
//	var body []byte
//	var err error
//	if resp.StatusCode == http.StatusOK {
//		body, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	defer resp.Body.Close()
//	output := string(body)
//	return output
//}

//func RPCCQSCQueryAccountGet(addr string) string {
//	aurl := QSCAccounturl + addr
//	resp, _ := http.Get(aurl)
//	var body []byte
//	var err error
//	if resp.StatusCode == http.StatusOK {
//		body, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	defer resp.Body.Close()
//	output := string(body)
//	return output
//}

////for QOS account query function
//func QOSQueryAccountGet(addr string) string {
//	aurl := QOSAccounturl + addr
//	resp, _ := http.Get(aurl)
//	var body []byte
//	var err error
//	if resp.StatusCode == http.StatusOK {
//		body, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	defer resp.Body.Close()
//	output := string(body)
//	return output
//}

////for QOSCommitResultCheck Restful interface
//func QOSCommitResultCheck(txhash, height string) string {
//	qstarskey := "heigth:" + height + ",hash:" + txhash
//	qrcurl := QResulturl + qstarskey
//	resp, _ := http.Get(qrcurl)
//	var body []byte
//	var err error
//	if resp.StatusCode == http.StatusOK {
//		body, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	defer resp.Body.Close()
//	output := fmt.Sprintf("This function has not been realized in QOS yet:\n%v", string(body))
//	return output
//
//}

//func TransferRecordsQuery(chainid, addr, cointype, offset, limit string) string {
//	trurl := TRurl + chainid + "/accounts/" + addr + "/transfer/?coin=" + cointype + "&offset=" + offset + "&limit=" + limit
//	resp, _ := http.Get(trurl)
//	var body []byte
//	var err error
//	if resp.StatusCode == http.StatusOK {
//		body, err = ioutil.ReadAll(resp.Body)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//
//	defer resp.Body.Close()
//	output := string(body)
//	return output
//}

//func Query(path string, key common.HexBytes) (res []byte, err error) {
//	opts := rpcclient.ABCIQueryOptions{
//		Height: 0,
//		Prove:  false,
//	}
//	result, err := RPC.ABCIQueryWithOptions(path, key, opts)
//	if err != nil {
//		return res, err
//	}
//	resp := result.Response
//	if !resp.IsOK() {
//		return res, errors.New("error query")
//	}
//
//	return resp.Value, nil
//}

func QueryAccount(addr qtypes.Address) (*types.QOSAccount, error) {
	key := AddressStoreKey(addr)
	res, err := Query("/store/acc/key", key)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, errors.New("account not exists")
	}
	var acc *types.QOSAccount
	err = Cdc.UnmarshalBinaryBare(res, &acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// 将地址转换成存储通用的key
func AddressStoreKey(addr qtypes.Address) []byte {
	return append([]byte(accountStoreKey), addr.Bytes()...)
}

//func Query(addr types.Address) ([]byte, error) {
func Query(path string, key []byte) ([]byte, error) {
	opts := rpcclient.ABCIQueryOptions{
		Height: 0,
		Prove:  true,
	}
	result, err := RPC.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return nil, err
	}
	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.New("query failed")
	}
	//var acc *QOSAccount
	//err = Cdc.UnmarshalBinaryBare(resp.Value, &acc)
	//if err != nil {
	//	return nil, err
	//}
	return resp.Value, nil
}

// 提交到联盟链上
func BroadcastTransferTxToQSC(txb string, broadcastModes string) string {
	txBytes, err := hex.DecodeString(txb)
	if err != nil {
		return err.Error()
	}
	var res *ctypes.ResultBroadcastTx
	switch broadcastModes {
	case "sync":
		res, err = RPC.BroadcastTxSync(txBytes)
		//默认异步
	default:
		res, err = RPC.BroadcastTxAsync(txBytes)
	}
	if err != nil {
		return err.Error()
	}
	resbyte, err := Cdc.MarshalJSON(res)
	if err != nil {
		return err.Error()
	}
	return string(resbyte)
}

func BroadcastTx(txb []byte) (string, error) {
	result, err := RPC.BroadcastTxSync(txb)
	if err != nil {
		return "", err
	}
	resbyte, err := Cdc.MarshalJSON(result)
	if err != nil {
		return "", err
	}
	return string(resbyte), nil
}
