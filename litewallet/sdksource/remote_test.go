package sdksource

import (
	"os/user"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	addr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	acout := GetAccount(rootDir, node, chainId, addr)
	t.Log(acout)
}

func TestTransfer(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	fromName := "cosmostest"
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
	chainId := "gaiav2"
	delegatorName := "cosmostest"
	password := "wm131421"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	validatorAddr := "cosmosvaloper109hhp349cay3zqczruyh0mtcrwp2emrcwa4474"
	delegationCoinStr := "10000000stake"
	feeStr := "10stake"
	broadcastMode := "async"
	delout := Delegate(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, delegationCoinStr, feeStr, broadcastMode)
	t.Log(delout)
}

func TestGetDelegationShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	validatorAddr := "cosmosvaloper109hhp349cay3zqczruyh0mtcrwp2emrcwa4474"
	getDelout := GetDelegationShares(rootDir, node, chainId, delegatorAddr, validatorAddr)
	t.Log(getDelout)
}

func TestUnbondingDelegation(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorName := "cosmostest"
	password := "wm131421"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	validatorAddr := "cosmosvaloper109hhp349cay3zqczruyh0mtcrwp2emrcwa4474"
	Ubdshares := "2000000stake"
	feeStr := "10stake"
	broadcastMode := "block"
	unbondDel := UnbondingDelegation(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, Ubdshares, feeStr, broadcastMode)
	t.Log(unbondDel)
}

func TestGetAllUnbondingDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	//validatorAddr := "cosmosvaloper1a8e4nvxw26c9ug9x687s65vxquszu3j82zezuc"
	getUbns := GetAllUnbondingDelegations(rootDir, node, chainId, delegatorAddr)
	t.Log(getUbns)
}

func TestGetBondValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	getBd := GetBondValidators(rootDir, node, chainId, delegatorAddr)
	t.Log(getBd)
}

func TestGetAllValidators(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	getVals := GetAllValidators(rootDir, node, chainId)
	t.Log(getVals)
}

func TestGetAllDelegations(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	getDels := GetAllDelegations(rootDir, node, chainId, delegatorAddr)
	t.Log(getDels)
}

func TestWithdrawDelegationReward(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorName := "cosmostest"
	password := "wm131421"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	validatorAddr := "cosmosvaloper1r6zq2r3d2cddz59p5cat2cazeavh8uq99wxw8l"
	feeStr := "10stake"
	broadcastMode := "async"
	withdrawRew := WithdrawDelegationReward(rootDir, node, chainId, delegatorName, password, delegatorAddr, validatorAddr, feeStr, broadcastMode)
	t.Log(withdrawRew)
}

func TestGetDelegationRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	validatorAddr := "cosmosvaloper1r6zq2r3d2cddz59p5cat2cazeavh8uq99wxw8l"
	getWithdraw := GetDelegationRewards(rootDir, node, chainId, delegatorAddr, validatorAddr)
	t.Log(getWithdraw)
}

func TestQueryTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	txHash := "25418A6B8CC7DB18F335D3CDD50FA8977A82DEFCE7C152C67EFD21285CB004A9"
	qTx := QueryTx(rootDir, node, chainId, txHash)
	t.Log(qTx)
}

func TestGetValSelfBondShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	validatorAddr := "cosmosvaloper1r6zq2r3d2cddz59p5cat2cazeavh8uq99wxw8l"
	vsb := GetValSelfBondShares(rootDir, node, chainId, validatorAddr)
	t.Log(vsb)
}

func TestGetDelegtorRewardsShares(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	daa := GetDelegtorRewardsShares(rootDir, node, chainId, delegatorAddr)
	t.Log(daa)
}

func TestTransferB4send(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	fromName := "cosmostest"
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
	chainId := "gaiav2"
	//fromName := "cosmos341"
	//password := "wm131421"
	//toStr := "cosmos1kklk4eqye6pla97dzmc03pw5lst7x0n4zt8syw"
	//coinStr := "10000stake"
	//feeStr := "1stake"
	broadcastMode := "block"
	txString := "c201282816a90a3ea8a3619a0a142080c64d44759f0eac8a476b661a7dd705760fd91214b5bf6ae404ce83fe97cd16f0f885d4fc17e33e751a0c0a057374616b65120331303012100a0a0a057374616b6512013110c09a0c1a6a0a26eb5ae9872103aa7393cb4998d47f13df3845f08d5af1c0901b4b144a03ff763df8578f4913261240683a26f4e5467ba52a5868cbc42e88820627498124aee0d415803ce42734b3091b96142e129135a48e7537790858e647e36122bce81b779b638bef094ec84cc5"
	Bt := BroadcastTransferTx(rootDir, node, chainId, txString, broadcastMode)
	t.Log(Bt)
}

func TestWithdrawDelegatorAllRewards(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	delegatorName := "cosmostest"
	password := "wm131421"
	delegatorAddr := "cosmos1yzqvvn2ywk0saty2ga4kvxna6uzhvr7e5g5jlf"
	feeStr := "10stake"
	broadcastMode := "block"
	wda := WithdrawDelegatorAllRewards(rootDir, node, chainId, delegatorName, password, delegatorAddr, feeStr, broadcastMode)
	t.Log(wda)
}

func TestLocalGenTx(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "tcp://192.168.1.184:26657"
	chainId := "gaiav2"
	fromName := "cosmostest"
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
