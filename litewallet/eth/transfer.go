package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"strconv"
)

func TransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount string, GasLimit int64) string {
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//get the nonce from the fromAddress to be dumped into tx
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//amount convertion to wei
	Amount, err := strconv.ParseFloat(amount,64)
	if err != nil {
		log.Fatal(err)
	}
	Amountwei := Amount*1000000000000000000
	value := big.NewInt(int64(Amountwei))

	//value := big.NewInt(amount)

	//gasPrice fethced from ethgasstation then convert the gasPrice of string to gwei
	gasAmount, err := strconv.ParseFloat(gasPrice,64)
	if err != nil {
		log.Fatal(err)
	}
	gasgwei := gasAmount*1000000000
	bigGas := big.NewInt(int64(gasgwei))

	//concert the to Address to byte format
	toAddress := common.HexToAddress(toAddr)

	gasLimit := uint64(GasLimit)
	//Generate the Tx body, the data field is nil for just sending ETH
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, bigGas, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//sign the Tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//SendTransaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()

}

//Transfer with ERC20 token
func TransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice string, GasLimit int64) string {
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//get the nonce for this tx
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//value is zero here for ERC20 tx
	value := big.NewInt(0) // in wei (0 eth)
	//set the amount and gasPrice for this Tx
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice fethced from ethgasstation then convert the gasPrice of string to gwei
	gasAmount, err := strconv.ParseFloat(gasPrice,64)
	if err != nil {
		log.Fatal(err)
	}
	gasgwei := gasAmount*1000000000
	bigGas := big.NewInt(int64(gasgwei))

	//the receiptant address
	toAddress := common.HexToAddress(toAddr)

	//begin smart contract part
	tokenAddress := common.HexToAddress(tokenAddr)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	//convert the tokenValue to wei in ERC20
	vamount, err := strconv.ParseFloat(tokenValue,64)
	if err != nil {
		log.Fatal(err)
	}
	vwei := vamount*1000000000000000000
	vstring := strconv.FormatFloat(vwei, 'f', -1, 32)

	Tamount := new(big.Int)
	//1000 token to transfer
	Tamount.SetString(vstring,10)

	paddedAmount := common.LeftPadBytes(Tamount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To: &tokenAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//create a transaction
	gasLimit := uint64(GasLimit)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, bigGas, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()
}

//Deprecated in cshare for mobile!
// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func GetPendingNonceAt(rootDir, node, fromName, password string) int64 {
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//get the nonce from the fromAddress to be dumped into tx
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	nonceInt := int64(nonce)
	//noncestr := strconv.FormatUint(nonce,10)
	return nonceInt
}

//Speedup Tnx with Pending nonce
func SpeedTransferETH(rootDir, node, fromName, password,toAddr, gasPrice, amount string, GasLimit, pendingNonce int64) string {
	//fromName generated from keyspace locally
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}
	//amount convertion to wei
	Amount, err := strconv.ParseFloat(amount,64)
	if err != nil {
		log.Fatal(err)
	}
	Amountwei := Amount*1000000000000000000
	value := big.NewInt(int64(Amountwei))

	//value := big.NewInt(amount)

	//gasPrice fethced from ethgasstation then convert the gasPrice of string to gwei
	gasAmount, err := strconv.ParseFloat(gasPrice,64)
	if err != nil {
		log.Fatal(err)
	}
	gasgwei := gasAmount*1000000000
	bigGas := big.NewInt(int64(gasgwei))

	//concert the to Address to byte format
	toAddress := common.HexToAddress(toAddr)

	gasLimit := uint64(GasLimit)

	//Generate the Tx body, the data field is nil for just sending ETH
	nonce := uint64(pendingNonce)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, bigGas, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//sign the Tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//SendTransaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()


}

func SpeedTransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr, tokenValue, gasPrice string, GasLimit, pendingNonce int64) string {
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}
	if fromName == "" {
		fmt.Println("no fromName input!")
	}
	//Fetch the privateKey to sign
	privateKey, err := FetchtoSign(rootDir, fromName, password)
	if err != nil {
		log.Fatal(err)
	}
	//value is zero here for ERC20 tx
	value := big.NewInt(0) // in wei (0 eth)
	//set the amount and gasPrice for this Tx
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gasPrice fethced from ethgasstation then convert the gasPrice of string to gwei
	gasAmount, err := strconv.ParseFloat(gasPrice,64)
	if err != nil {
		log.Fatal(err)
	}
	gasgwei := gasAmount*1000000000
	bigGas := big.NewInt(int64(gasgwei))

	//the receiptant address
	toAddress := common.HexToAddress(toAddr)

	//begin smart contract part
	tokenAddress := common.HexToAddress(tokenAddr)
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	//convert the tokenValue to wei in ERC20
	vamount, err := strconv.ParseFloat(tokenValue,64)
	if err != nil {
		log.Fatal(err)
	}
	vwei := vamount*1000000000000000000
	vstring := strconv.FormatFloat(vwei, 'f', -1, 32)

	Tamount := new(big.Int)
	//1000 token to transfer
	Tamount.SetString(vstring,10)

	paddedAmount := common.LeftPadBytes(Tamount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
	//	To: &tokenAddress,
	//	Data: data,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//create a transaction
	gasLimit := uint64(GasLimit)
	nonce := uint64(pendingNonce)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, bigGas, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	return signedTx.Hash().Hex()

}
