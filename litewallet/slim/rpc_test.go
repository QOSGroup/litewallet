package slim

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/module"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"testing"
)

var chainId = "aquarius-1001"

func TestQueryAccount(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	addr := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	bytes, err := QueryAccount(addr)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(bytes))
}

//func TestQSCQueryAccountGet(t *testing.T) {
//	txs.SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
//	addr := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
//	Aout := txs.QSCQueryAccountGet(addr)
//	t.Log(Aout)
//}

func TestTransferSend(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	addrto := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	Tout, err := Transfer(addrto, coinstr, privkey, chainId)
	if err != nil {
		t.Log(err)
	}
	t.Log(Tout)
}

//func TestQSCCreateSignedTransfer(t *testing.T) {
//	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
//	addrto := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
//	coinstr := "10000qos"
//	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
//	chainid := "qos-test"
//	Tout, err := client.CreateSignedTransfer(addrto, coinstr, privkey, chainid)
//	if err != nil {
//		t.Log(err)
//	}
//	t.Log(Tout)
//}

//stake
func TestDelegationSend(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	//validatorAddress := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	validatorAddress := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	//coinstr := "10000qos"
	coins := int64(1000)
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	Tout, err := Delegation(validatorAddress, coins, privkey, chainId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestUnbondDelegationSend(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	validatorAddress := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	//coinstr := "10000qos"
	coins := int64(500)
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	Tout, err := UnbondDelegation(validatorAddress, coins, privkey, chainId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestRedelegations(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	formValidatorAddr := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	toValidatorAddr := "address1demk4rqhfc5ewlsefa2g805ycnzse26mr534nu"
	coins := int64(500)
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	Tout, err := ReDelegation("47.103.78.91:26657", formValidatorAddr, toValidatorAddr, coins, privkey, chainId)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestQueryValidatorInfo(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	validatorAddr := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	Tout, err := QueryValidatorInfo("47.103.78.91:26657", validatorAddr)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryValidators(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryValidators("47.103.78.91:26657")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryValidatorMissedVoteInfo(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	address := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	Tout, err := QueryValidatorMissedVoteInfo("47.103.78.91:26657", address)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryDelegationInfo(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	ownerAddr := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	delegatorAddr := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Tout, err := QueryDelegationInfo("47.103.78.91:26657", ownerAddr, delegatorAddr)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryDelegations(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	address := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Tout, err := QueryDelegations("47.103.78.91:26657", address)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryUnbondings(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	address := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Tout, err := QueryUnbondings("47.103.78.91:26657", address)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryRedelegations(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	address := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Tout, err := QueryRedelegations("47.103.78.91:26657", address)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

//approve
func TestQueryApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	Tout, err := QueryApprove(toAddr, privkey)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

//distribution
func TestQueryDelegatorIncomeInfo(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	ownerAddr := "address1tl0gdpdwjz2g77s7qcf0jvr4sxc0szw0nlk08t"
	Tout, err := QueryDelegatorIncomeInfo("47.103.78.91:26657", privkey, ownerAddr)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryCommunityFeePool(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryCommunityFeePool("47.103.78.91:26657")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryProposal(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryProposal("47.103.78.91:26657", 1)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryProposals(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryProposals("47.103.78.91:26657", "", "", "deposit_period")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryVote(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryVote("47.103.78.91:26657", 1, "")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryVotes(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryVotes("47.103.78.91:26657", 1)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryDeposit(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryDeposit("47.103.78.91:26657", 1, "")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryDeposits(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryDeposits("47.103.78.91:26657", 1)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryTally(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	Tout, err := QueryTally("47.103.78.91:26657", 1, "")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

func TestQueryTx(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	hashHex := "B5EECB27939D969556C61E00BDE8C910FFE3BE47BFA0356B57F58847D6502B70"
	Tout, err := QueryTx("47.103.78.91:26657", hashHex)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(Tout))
}

//func TestTransferTxToQSC(t *testing.T) {
//	t.Log("1111111111111111")
//
//	txs.SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
//	t.Log("2222222222222222")
//
//	privkey := "wnEmxnWFgT93M5a9l7aPTdkxM8MLoenyMe60sD/8rqzslA7MvfoHydXqL4QGbplLhIlEbLAZ/0ue9G1rjBFMfQ=="
//	chainid := "test-chain-xHEkEv"
//	Tout := client.AdvertisersTrue(privkey, "ATOM", "100000", chainid)
//	t.Log(Tout)
//
//	result := txs.BroadcastTransferTxToQSC(Tout, "sync")
//	t.Log(result)
//}

//func TestQueryAccount(t *testing.T) {
//	t.Log("1111111111111111")
//
//	txs.SetBlockchainEntrance("47.105.52.237:26657", "forQmoonAddr")
//	t.Log("2222222222222222")
//
//	privkey := "j7wbT1lfG0ZRObMKwyYkVFpwRW0OlzFcgVL6ZwQJjeicPcStxppyWMpD1tfFg2YEcDu5JH+M6g+EhSBmjHq0bg=="
//
//	var key ed25519local.PrivKeyEd25519
//	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
//	err1 := txs.Cdc.UnmarshalJSON([]byte(ts), &key)
//	if err1 != nil {
//		fmt.Println(err1)
//	}
//
//	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
//	address, _ := types.GetAddrFromBech32(addrben32)
//
//	fmt.Println("address", string(address))
//
//	account, err := txs.RpcQueryAccount(address)
//
//	if err != nil {
//		t.Error(err.Error())
//	} else {
//		t.Log("111", account.Nonce)
//	}
//}

func TestCommHandler(t *testing.T) {
	t.Log("1111111111111111")

	txs.SetBlockchainEntrance("localhost:26657", "forQmoonAddr")
	t.Log("2222222222222222")

	privkey := "9QkouVPl29N2v1lBO1+azUDqm38fAgs6d3Xo8DcnCus7xjMqsavhc190xCGzZuXcjapUahi7Y7v2DD4hzVCAsQ=="
	chainid := "test-chain-dpYlL3"

	args := "[\"address1y9r4pjjnvkmpvw46de8tmwunw4nx4qnz2ax5ux\",\"0\",\"abcde\",\"20\",\"20\",\"10\",\"50\",\"20\",\"3\",\"ATOM\"]"

	Tout := module.CommHandler("ArticleTx", privkey, args, chainid)
	t.Log(Tout)

	result := txs.BroadcastTransferTxToQSC(Tout, "sync")
	t.Log(result)
}

func TestLocalTxGen(t *testing.T) {
	fromStr := "address1vpszt2jp2j8m5l3mutvqserzuu9uylmzydqaj9"

	toStr := "address1eep59h9ez4thymept8nxl0padlrc6r78fsjmp3"

	coinstr := "2qos"
	//generate singed Tx
	chainid := "capricorn-2000"
	nonce := int64(1)
	//gas := NewBigInt(int64(0))

	//PrivKey output
	privkey := "sV5sRbwnR8DddL5e4UC1ntKPiOtGEaOFAqvePTfhJFI9GcC28zmPURSUI6C1oBlnk2ykBcAtIbYUazuCexWyqg=="

	jasonpayload := LocalTxGen(fromStr, toStr, coinstr, chainid, privkey, nonce)

	t.Log("msg\n", string(jasonpayload))
}

////HTTP POST to QOS chain
//func TestHttpBrToChain(t *testing.T) {
//	fromStr := "address1vpszt2jp2j8m5l3mutvqserzuu9uylmzydqaj9"
//	toStr := "address1eep59h9ez4thymept8nxl0padlrc6r78fsjmp3"
//	coinstr := "2qos"
//	//generate singed Tx
//	chainid := "capricorn-1000"
//	nonce := int64(1)
//	//gas := NewBigInt(int64(0))
//	//PrivKey output
//	privkey := "sV5sRbwnR8DddL5e4UC1ntKPiOtGEaOFAqvePTfhJFI9GcC28zmPURSUI6C1oBlnk2ykBcAtIbYUazuCexWyqg=="
//
//	jasonpayload := LocalTxGen(fromStr, toStr, coinstr, chainid, privkey, nonce)
//
//	//tbt := new(types.Tx)
//	//err := Cdc.UnmarshalJSON(jasonpayload, tbt)
//	//if err != nil {
//	//	fmt.Println(err)
//	//}
//	//
//	//txBytes, err := Cdc.MarshalBinaryBare(jasonpayload)
//	//if err != nil {
//	//	panic("use cdc encode object fail")
//	//}
//
//	client := rpc_client.NewHTTP("tcp://192.168.1.183:26657", "/websocket")
//	result, err := client.BroadcastTxCommit(jasonpayload)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	t.Log(result)
//}
