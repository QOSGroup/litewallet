package app

import (
	approve_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/approve/txs"
	bank_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	gov_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/txs"
	stake_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
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
	//cdc := baseabci.MakeQBaseCodec()
	cdc := go_amino.NewCodec()
	RegisterAmino(cdc)
	RegisterCodec(cdc)
	return cdc
}

// RegisterAmino registers all crypto related types in the given (amino) codec.
func RegisterAmino(cdc *go_amino.Codec) {
	//cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	//cdc.RegisterConcrete(ed25519.PubKeyEd25519{}, ed25519.PubKeyAminoName, nil)
	//cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
	//cdc.RegisterConcrete(multisig.PubKeyMultisigThreshold{}, multisig.PubKeyMultisigThresholdAminoRoute, nil)
	//
	//cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	//cdc.RegisterConcrete(ed25519.PrivKeyEd25519{}, ed25519.PrivKeyAminoName, nil)
	//cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName, nil)

	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519local.PubKeyEd25519{}, ed25519local.Ed25519PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519local.PrivKeyEd25519{}, ed25519local.Ed25519PrivKeyAminoRoute, nil)
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
