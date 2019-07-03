package slim

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/tendermint/go-amino"
)

type Codec = amino.Codec

func NewCodec() *Codec {
	cdc := amino.NewCodec()
	return cdc
}

var Cdc *Codec

func init() {
	cdc := NewCodec()
	RegisterAmino(cdc)
	RegisterCodec(cdc)
	Cdc = cdc.Seal()
}

// RegisterAmino registers all crypto related types in the given (amino) codec.
func RegisterAmino(cdc *amino.Codec) {
	cdc.RegisterInterface((*ed25519local.PubKey)(nil), nil)
	cdc.RegisterConcrete(ed25519local.PubKeyEd25519{},
		ed25519local.Ed25519PubKeyAminoRoute, nil)

	cdc.RegisterInterface((*ed25519local.PrivKey)(nil), nil)
	cdc.RegisterConcrete(ed25519local.PrivKeyEd25519{},
		ed25519local.Ed25519PrivKeyAminoRoute, nil)
}

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(&Signature{}, "qbase/txs/signature", nil)
	cdc.RegisterConcrete(&TxStd{}, "qbase/txs/stdtx", nil)
	cdc.RegisterInterface((*ITx)(nil), nil)
	cdc.RegisterConcrete(&TxTransfer{}, "qos/txs/TxTransfer", nil)

	cdc.RegisterConcrete(&QOSAccount{}, "qos/types/QOSAccount", nil)
	cdc.RegisterConcrete(&BaseAccount{}, "qbase/account/BaseAccount", nil)
	cdc.RegisterConcrete(&InvestTx{}, "qstars/InvestTx", nil)
	cdc.RegisterConcrete(&AdvertisersTx{}, "jianqian/AdvertisersTx", nil)
	cdc.RegisterConcrete(&AuctionTx{}, "jianqian/AuctionTx", nil)
	cdc.RegisterConcrete(&ExtractTx{}, "jianqian/ExtractTx", nil)
	cdc.RegisterConcrete(&JianQianTx{}, "jianqian/JianQianTx", nil)

}
