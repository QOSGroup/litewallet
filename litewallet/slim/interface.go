package slim

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/app"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	"github.com/QOSGroup/litewallet/litewallet/slim/module"
	approve_client "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/client"
	bank_client "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/client"
	distribution_client "github.com/QOSGroup/litewallet/litewallet/slim/module/distribution/client"
	gov_client "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/client"
	stake_client "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/client"
)

//only need the following arguments, it`s enough!
func QueryAccount(remote, addr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	qosAccount, err := account.QueryAccount(cliCtx, addr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(qosAccount)
}

//only need the following arguments, it`s enough!
func Transfer(remote, addrto, coinstr, privkey, chainid string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	tx, err := bank_client.CreateTransfer(cliCtx, addrto, coinstr, privkey, chainid)
	if err != nil {
		return nil, err
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
	result, err := cliCtx.BroadcastTxSync(tx)
	if err != nil {
		return nil, err
	}
	return cliCtx.Codec.MarshalJSON(result)
}

// stake
func Delegation(remote, addrto string, coins int64, privkey, chainid string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	tx, err := stake_client.CreateDelegation(cliCtx, addrto, coins, privkey, chainid)
	if err != nil {
		return nil, err
	}
	result, err := cliCtx.BroadcastTxSync(tx)
	if err != nil {
		return nil, err
	}
	return cliCtx.Codec.MarshalJSON(result)
}

func UnbondDelegation(remote, addrto string, coins int64, privkey, chainid string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	tx, err := stake_client.CreateUnbondDelegation(cliCtx, addrto, coins, privkey, chainid)
	if err != nil {
		return nil, err
	}
	result, err := cliCtx.BroadcastTxSync(tx)
	if err != nil {
		return nil, err
	}
	return cliCtx.Codec.MarshalJSON(result)
}

func ReDelegation(remote, fromValidatorAddr, toValidatorAddr string, coins int64, privkey, chainid string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	tx, err := stake_client.CreateReDelegationCommand(cliCtx, fromValidatorAddr, toValidatorAddr, coins, privkey, chainid)
	if err != nil {
		return nil, err
	}
	result, err := cliCtx.BroadcastTxSync(tx)
	if err != nil {
		return nil, err
	}
	return cliCtx.Codec.MarshalJSON(result)
}

func QueryValidatorInfo(remote, validatorAddr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	validator, err := stake_client.QueryValidatorInfo(cliCtx, validatorAddr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(validator)
}

func QueryValidators(remote string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	return stake_client.QueryValidators(cliCtx)
}

func QueryValidatorMissedVoteInfo(remote, address string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := stake_client.QueryValidatorMissedVoteInfo(cliCtx, address)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryDelegationInfo(remote, ownerAddr, delegatorAddr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := stake_client.QueryDelegationInfo(cliCtx, ownerAddr, delegatorAddr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryDelegations(remote, address string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := stake_client.QueryDelegations(cliCtx, address)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryUnbondings(remote, address string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := stake_client.QueryUnbondings(cliCtx, address)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryRedelegations(remote, address string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := stake_client.QueryRedelegations(cliCtx, address)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

// approve
func QueryApprove(remote, addrto, privkey string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	approve, err := approve_client.QueryApprove(cliCtx, addrto, privkey)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(approve)
}

// distribution
func QueryDelegatorIncomeInfo(remote, privkey, validatorAddr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := distribution_client.QueryDelegatorIncomeInfo(cliCtx, privkey, validatorAddr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryCommunityFeePool(remote string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := distribution_client.QueryCommunityFeePool(cliCtx)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryProposal(remote string, pId uint64) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryProposal(cliCtx, pId)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryProposals(remote, depositor, voter, statusStr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryProposals(cliCtx, depositor, voter, statusStr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryVote(remote string, pId uint64, addrStr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryVote(cliCtx, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryVotes(remote string, pId uint64) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryVotes(cliCtx, pId)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryDeposit(remote string, pId uint64, addrStr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryDeposit(cliCtx, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryDeposits(remote string, pId uint64) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryDeposits(cliCtx, pId)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryTally(remote string, pId uint64, addrStr string) ([]byte, error) {
	var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	result, err := gov_client.QueryTally(cliCtx, pId, addrStr)
	if err != nil {
		return nil, err
	}
	return app.Cdc.MarshalJSON(result)
}

func QueryTx(remote, hashHex string) ([]byte, error) {
	//var cliCtx = context.NewCLIContext(remote).WithCodec(app.Cdc)
	txResponse, err := module.QueryTx(remote, hashHex)
	if err != nil {
		return nil, err
	}

	if txResponse.Empty() {
		return nil, fmt.Errorf("No transaction found with hash %s", hashHex)
	}
	return app.Cdc.MarshalJSON(txResponse)
}

//func GetTx(tx string) string {
//	txBytes, err := hex.DecodeString(tx)
//	if err != nil {
//		return err.Error()
//	}
//	txhashs := strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
//	return string(txhashs)
//}
