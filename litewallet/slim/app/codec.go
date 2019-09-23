package app

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/baseabci"
	approve_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/txs"
	bank_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	gov_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/txs"
	stake_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/types"
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

var Cdc = MakeCodec()

func MakeCodec() *go_amino.Codec {
	cdc := baseabci.MakeQBaseCodec()
	//cdc := go_amino.NewCodec()
	ModuleRegisterCodec(cdc)
	types.RegisterCodec(cdc)
	return cdc
}

func ModuleRegisterCodec(cdc *go_amino.Codec) {
	//noPanicRegisterInterface(cdc)
	//ModuleBasics.RegisterCodec(cdc)
	//types.RegisterCodec(cdc)
	//cert.RegisterCodec(cdc)

	approve_txs.RegisterCodec(cdc)
	stake_txs.RegisterCodec(cdc)
	bank_txs.RegisterCodec(cdc)
	gov_txs.RegisterCodec(cdc)
}
