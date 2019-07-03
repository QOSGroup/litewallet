package sdksource

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

//This package is used to add complementary support on codec, e.g. MsgSend
// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/MsgSend", nil)
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cryptoAmino.RegisterAmino(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&authtypes.BaseAccount{}, "cosmos-sdk/Account", nil)
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&authtypes.BaseVestingAccount{}, "cosmos-sdk/BaseVestingAccount", nil)
	cdc.RegisterConcrete(&authtypes.ContinuousVestingAccount{}, "cosmos-sdk/ContinuousVestingAccount", nil)
	cdc.RegisterConcrete(&authtypes.DelayedVestingAccount{}, "cosmos-sdk/DelayedVestingAccount", nil)
	cdc.RegisterConcrete(authtypes.StdTx{}, "cosmos-sdk/StdTx", nil)
}

// module bank codec for send message codec
var BankCdc *codec.Codec

func init() {
	BankCdc = codec.New()
	RegisterCodec(BankCdc)
	BankCdc.Seal()
}
