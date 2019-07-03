package sdksource

import (
	"os/user"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1vk4ark02kc7ac9ctgegml66496a8nyz0wyfn33"
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	acout := GetAccount(rootDir, node, chainId, addr)
	t.Log(acout)
}

func TestTransfer(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	fromName := "cosmos"
	password := "wm131421"
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "1000000stake"
	feeStr := "20stake"
	broadcastMode := "sync"
	transout := Transfer(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr, broadcastMode)
	t.Log(transout)
}

func TestDelegate(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorName := "cosmos"
	password := "wm131421"
	delegatorAddr := "cosmos1vk4ark02kc7ac9ctgegml66496a8nyz0wyfn33"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	delegationCoinStr := "1000000000stake"
	feeStr := "10stake"
	broadcastMode := "block"
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, broadcastMode)
	t.Log(delout)
}

func TestGetDelegationShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	getDelout := GetDelegationShares(rootDir, node, chainId, delegatorAddr, validatorAddr)
	t.Log(getDelout)
}

func TestUnbondingDelegation(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorName := "c34banker"
	password := "wm131421"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	Ubdshares := "20000000stake"
	feeStr := "1stake"
	broadcastMode := "block"
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, Ubdshares, feeStr, broadcastMode)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1ne8hnx92k8x7cvluvhkphtk5kpvzenvns7g48g"
	//validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getUbns := GetAllUnbondingDelegations(rootDir, node, chainId, delegatorAddr)
	t.Log(getUbns)
}

func TestGetBondValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	getBd := GetBondValidators(rootDir, node, chainId, delegatorAddr)
	t.Log(getBd)
}

func TestGetAllValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	getVals := GetAllValidators(rootDir, node, chainId)
	t.Log(getVals)
}

func TestGetAllDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	getDels := GetAllDelegations(rootDir, node, chainId, delegatorAddr)
	t.Log(getDels)
}

func TestWithdrawDelegationReward(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorName := "c34banker"
	password := "wm131421"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	feeStr := "1stake"
	broadcastMode := "block"
	withdrawRew := WithdrawDelegationReward(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, broadcastMode)
	t.Log(withdrawRew)
}

func TestGetDelegationRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1ne8hnx92k8x7cvluvhkphtk5kpvzenvns7g48g"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	getWithdraw := GetDelegationRewards(rootDir, node, chainId, delegatorAddr, validatorAddr)
	t.Log(getWithdraw)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	txHash := "1730528BCDD06039C82C27C87E6D19C8CA2D4540BD2DB34A3BD67EF04230B1FB"
	qTx := QueryTx(rootDir, node, chainId, txHash)
	t.Log(qTx)
}

func TestGetValSelfBondShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	validatorAddr := "cosmosvaloper1xwz2req975fqnvrrx9me7vwyz25paxflkx60pe"
	vsb := GetValSelfBondShares(rootDir, node, chainId, validatorAddr)
	t.Log(vsb)
}

func TestGetDelegtorRewardsShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	daa := GetDelegtorRewardsShares(rootDir, node, chainId, delegatorAddr)
	t.Log(daa)
}

func TestTransferB4send(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	fromName := "c34banker"
	password := "wm131421"
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "100stake"
	feeStr := "1stake"
	Tx := TransferB4send(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr)
	t.Log(Tx)

}

func TestBroadcastTransferTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	//fromName := "cosmos341"
	//password := "wm131421"
	//toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	//coinStr := "10000stake"
	//feeStr := "1stake"
	broadcastMode := "block"
	txString := "c201f0625dee0a3ea8a3619a0a149e4f7998aab1cdec33fc65ec1baed4b0582ccd931214b5bf6ae404ce83fe97cd16f0f885d4fc17e33e751a0c0a057374616b65120331303012100a0a0a057374616b6512013110c09a0c1a6a0a26eb5ae9872102ae6ea19e838daa09b8e614dbb59636bae7007b86b85c165d8100ba50b365d4861240647008bcdc33854d7270045e18ad28e61f03d062ce7a63c60aca6026fb3b4d812bbc1898fd08f0b726b0299c87be74f4870aa0ee099db8778a6cfe038685eb8f"
	Bt := BroadcastTransferTx(rootDir, node, chainId, txString, broadcastMode)
	t.Log(Bt)
}

func TestWithdrawDelegatorAllRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	delegatorName := "c34banker"
	password := "wm131421"
	delegatorAddr := "cosmos1xwz2req975fqnvrrx9me7vwyz25paxflnjw6d2"
	feeStr := "10stake"
	broadcastMode := "block"
	wda := WithdrawDelegatorAllRewards(rootDir, node, chainId, delegatorName, password, delegatorAddr, feeStr, broadcastMode)
	t.Log(wda)
}

func TestLocalGenTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "cosmosv34"
	fromName := "cosmos"
	password := "wm131421"
	toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	coinStr := "100stake"
	feeStr := "1stake"
	Txs := LocalGenTx(rootDir, node, chainId, fromName, password, toStr, coinStr, feeStr)
	//txb := []byte(Txs)
	t.Log(Txs)
}

func TestMsgSendGetSignBytes(t *testing.T) {
	addr1 := sdk.AccAddress([]byte("input"))
	addr2 := sdk.AccAddress([]byte("output"))
	coins := sdk.NewCoins(sdk.NewInt64Coin("atom", 10))
	var msg = NewMsgSend(addr1, addr2, coins)
	res := msg.GetSignBytes()

	expected := `{"type":"cosmos-sdk/MsgSend","value":{"amount":[{"amount":"10","denom":"atom"}],"from_address":"cosmos1d9h8qat57ljhcm","to_address":"cosmos1da6hgur4wsmpnjyg"}}`
	require.Equal(t, expected, string(res))
	t.Log(string(res))
}
