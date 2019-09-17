package slim

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/module"
	approve_client "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/client"
	bank_client "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/client"
	distribution_client "github.com/QOSGroup/litewallet/litewallet/slim/module/distribution/client"
	gov_client "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/client"
	stake_client "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/client"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
)

//only need the following arguments, it`s enough!
func QueryAccount(addr string) ([]byte, error) {
	qosAccount, err := module.GetAccountFromBech32Addr(addr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(qosAccount)
}

//only need the following arguments, it`s enough!
func Transfer(addrto, coinstr, privkey, chainid string) (string, error) {
	tx, err := bank_client.CreateSignedTransfer(addrto, coinstr, privkey, chainid)
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

// stake
func Delegation(addrto string, coins int64, privkey, chainid string) (string, error) {
	tx, err := stake_client.CreateSignedDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func UnbondDelegation(addrto string, coins int64, privkey, chainid string) (string, error) {
	tx, err := stake_client.CreateSignedUnbondDelegation(addrto, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func ReDelegation(remote, fromValidatorAddr, toValidatorAddr string, coins int64, privkey, chainid string) (string, error) {
	tx, err := stake_client.CreateReDelegationCommand(fromValidatorAddr, toValidatorAddr, coins, privkey, chainid)
	if err != nil {
		return "", err
	}
	return txs.BroadcastTx(tx)
}

func QueryValidatorInfo(remote, validatorAddr string) ([]byte, error) {
	validator, err := stake_client.QueryValidatorInfo(remote, validatorAddr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(validator)
}

func QueryValidators(remote string) ([]byte, error) {
	return stake_client.QueryValidators(remote)
}

func QueryValidatorMissedVoteInfo(remote, address string) ([]byte, error) {
	result, err := stake_client.QueryValidatorMissedVoteInfo(remote, address)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryDelegationInfo(remote, ownerAddr, delegatorAddr string) ([]byte, error) {
	result, err := stake_client.QueryDelegationInfo(remote, ownerAddr, delegatorAddr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryDelegations(remote, address string) ([]byte, error) {
	result, err := stake_client.QueryDelegations(remote, address)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryUnbondings(remote, address string) ([]byte, error) {
	result, err := stake_client.QueryUnbondings(remote, address)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryRedelegations(remote, address string) ([]byte, error) {
	result, err := stake_client.QueryRedelegations(remote, address)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

// approve
func QueryApprove(addrto, privkey string) ([]byte, error) {
	approve, err := approve_client.QueryApprove(addrto, privkey)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(approve)
}

// distribution
func QueryDelegatorIncomeInfo(remote, privkey, ownerAddr string) ([]byte, error) {
	result, err := distribution_client.QueryDelegatorIncomeInfo(remote, privkey, ownerAddr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryCommunityFeePool(remote string) ([]byte, error) {
	result, err := distribution_client.QueryCommunityFeePool(remote)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryProposal(remote string, pId uint64) ([]byte, error) {
	result, err := gov_client.QueryProposal(remote, pId)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryProposals(remote, depositor, voter, statusStr string) ([]byte, error) {
	result, err := gov_client.QueryProposals(remote, depositor, voter, statusStr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryVote(remote string, pId uint64, addrStr string) ([]byte, error) {
	result, err := gov_client.QueryVote(remote, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryVotes(remote string, pId uint64) ([]byte, error) {
	result, err := gov_client.QueryVotes(remote, pId)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryDeposit(remote string, pId uint64, addrStr string) ([]byte, error) {
	result, err := gov_client.QueryDeposit(remote, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryDeposits(remote string, pId uint64) ([]byte, error) {
	result, err := gov_client.QueryDeposits(remote, pId)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryTally(remote string, pId uint64, addrStr string) ([]byte, error) {
	result, err := gov_client.QueryTally(remote, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return txs.Cdc.MarshalJSON(result)
}

func QueryTx(remote, hashHex string) ([]byte, error) {
	txResponse, err := module.QueryTx(remote, hashHex)
	if err != nil {
		return nil, err
	}

	if txResponse.Empty() {
		return nil, fmt.Errorf("No transaction found with hash %s", hashHex)
	}
	return txs.Cdc.MarshalJSON(txResponse)
}

//func GetTx(tx string) string {
//	txBytes, err := hex.DecodeString(tx)
//	if err != nil {
//		return err.Error()
//	}
//	txhashs := strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
//	return string(txhashs)
//}
