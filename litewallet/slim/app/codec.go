package app

import (
	approve_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/txs"
	stake_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	bank_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	gov_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/txs"
	go_amino "github.com/tendermint/go-amino"
)

//var cdc = go_amino.NewCodec()
//
//// 包初始化，注册codec
//func init() {
//	cryptoAmino.RegisterAmino(cdc)
//	cdc.RegisterInterface((*bacc.Account)(nil), nil)
//	RegisterCodec(cdc)
//}

func MakeCodec() *go_amino.Codec {
	//cdc := baseabci.MakeQBaseCodec()
	cdc := go_amino.NewCodec()
	RegisterCodec(cdc)
	return cdc
}

func RegisterCodec(cdc *go_amino.Codec) {
	//noPanicRegisterInterface(cdc)
	//ModuleBasics.RegisterCodec(cdc)
	//types.RegisterCodec(cdc)
	//cert.RegisterCodec(cdc)

	approve_txs.RegisterCodec(cdc)
	stake_txs.RegisterCodec(cdc)
	bank_txs.RegisterCodec(cdc)
	gov_txs.RegisterCodec(cdc)
}
