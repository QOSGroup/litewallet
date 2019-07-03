package sdksource

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

//This package is used to add complementary support on codec, e.g. MsgSend
// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/MsgSend", nil)
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cryptoAmino.RegisterAmino(cdc)
}

// module bank codec for send message codec
var BankCdc *codec.Codec

func init() {
	BankCdc = codec.New()
	RegisterCodec(BankCdc)
	BankCdc.Seal()
}
