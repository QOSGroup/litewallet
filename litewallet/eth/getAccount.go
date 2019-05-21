package eth

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/lcd/cosmoswallet/eth/contracts_erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func GetAccount(node, addr string) string{
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//convert the addr string to common.Address type
	address := common.HexToAddress(addr)

	//get the latest block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(header.Number.Int64())

	balance, err := client.BalanceAt(context.Background(), address, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	accountStr := ethValue.String() + "ETH"
	return accountStr
}

func GetAccountERC20(node, addr, tokenAddr string) string {
	//setup the client, here use the infura own project "eth_wallet" node="https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	client, err := ethclient.Dial(node)
	if err != nil {
		log.Fatal(err)
	}

	//ERC20 Token QT Address
	tokenAddress := common.HexToAddress(tokenAddr)
	instance, err := contracts_erc20.NewContractsErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	//convert the addr string to common.Address type
	address := common.HexToAddress(addr)
	//Enter smart contract querying
	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	//details of the token in ERC20 standards: including symbol and decimals
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	digit := int(decimals)

	//format the output in wei
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(digit)))

	accountStr := ethValue.String() + symbol
	return accountStr
}