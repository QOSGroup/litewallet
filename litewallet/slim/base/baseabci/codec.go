package baseabci

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	cryptoAmino "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/encoding/amino"
	go_amino "github.com/tendermint/go-amino"
	//cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

func MakeQBaseCodec() *go_amino.Codec {

	var cdc = go_amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	RegisterCodec(cdc)

	return cdc
}

//// RegisterAmino registers all crypto related types in the given (amino) codec.
//func RegisterAmino(cdc *go_amino.Codec) {
//	//cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
//	//cdc.RegisterConcrete(ed25519.PubKeyEd25519{}, ed25519.PubKeyAminoName, nil)
//	//cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName, nil)
//	//cdc.RegisterConcrete(multisig.PubKeyMultisigThreshold{}, multisig.PubKeyMultisigThresholdAminoRoute, nil)
//	//
//	//cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
//	//cdc.RegisterConcrete(ed25519.PrivKeyEd25519{}, ed25519.PrivKeyAminoName, nil)
//	//cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName, nil)
//
//	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
//	cdc.RegisterConcrete(ed25519local.PubKeyEd25519{}, ed25519local.Ed25519PubKeyAminoRoute, nil)
//
//	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
//	cdc.RegisterConcrete(ed25519local.PrivKeyEd25519{}, ed25519local.Ed25519PrivKeyAminoRoute, nil)
//}

func RegisterCodec(cdc *go_amino.Codec) {
	txs.RegisterCodec(cdc)
	account.RegisterCodec(cdc)
	//keys.RegisterCodec(cdc)
	//consensus.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
}
