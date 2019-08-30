package slim

import (
	"bytes"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/http"
)

//only need the following arguments, it`s enough!
func QSCTransferSendStr(addrto, coinstr, privkey, chainid string) string {
	jasonpayload, err := client.CreateSignedTransfer(addrto, coinstr, privkey, chainid)
	if err != nil {
		fmt.Println(err)
	}
	datas := bytes.NewBuffer([]byte(jasonpayload))
	//aurl := txs.Accounturl + "txSend"
	//req, _ := http.NewRequest("POST", aurl, datas)
	//req.Header.Set("Content-Type", "application/json")
	//clt := http.Client{}
	//resp, _ := clt.Do(req)
	//body, err := ioutil.ReadAll(resp.Body)
	//defer resp.Body.Close()
	//output := string(body)
	output, err := http.Request(datas)
	return output
}

//only need the following arguments, it`s enough!
func QSCDelegationSendStr(addrto string, coins int64, privkey, chainid string) string {
	jasonpayload, err := client.CreateSignedDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		fmt.Println(err)
	}
	datas := bytes.NewBuffer([]byte(jasonpayload))
	output, err := http.Request(datas)
	return output
}

//only need the following arguments, it`s enough!
func QSCUnbondDelegationSendStr(addrto string, coins int64, privkey, chainid string) string {
	jasonpayload, err := client.CreateSignedUnbondDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		fmt.Println(err)
	}
	datas := bytes.NewBuffer([]byte(jasonpayload))
	output, err := http.Request(datas)
	return output
}
