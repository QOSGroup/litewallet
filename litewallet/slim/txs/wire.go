package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	txs3 "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	gov_types "github.com/QOSGroup/litewallet/litewallet/slim/module/gov/types"
	txs2 "github.com/QOSGroup/litewallet/litewallet/slim/module/stake/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/core/types"
	qtypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"github.com/tendermint/go-amino"
)

//
//import (
//	"github.com/QOSGroup/qbase/account"
//	"github.com/QOSGroup/qbase/baseabci"
//	"github.com/QOSGroup/qbase/context"
//	"github.com/QOSGroup/qbase/mapper"
//	"github.com/QOSGroup/qbase/txs"
//	"github.com/tendermint/go-amino"
//	"github.com/tendermint/tendermint/crypto"
//	"github.com/tendermint/tendermint/crypto/ed25519"
//	cmn "github.com/tendermint/tendermint/libs/common"
//	"log"
//	"reflect"
//)
//
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

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterConcrete(&ctypes.ResultTx{}, "qbase/types/ResultTx", nil)

	//cdc.RegisterConcrete(&txs.QcpTxResult{}, "qbase/txs/qcpresult", nil)
	cdc.RegisterConcrete(&txs.Signature{}, "qbase/txs/signature", nil)
	cdc.RegisterConcrete(&txs.TxStd{}, "qbase/txs/stdtx", nil)
	//cdc.RegisterConcrete(&txs.TxQcp{}, "qbase/txs/qcptx", nil)
	cdc.RegisterInterface((*txs.ITx)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)

	cdc.RegisterConcrete(&txs3.TxTransfer{}, "transfer/txs/TxTransfer", nil)

	cdc.RegisterConcrete(&txs2.TxCreateDelegation{}, "stake/txs/TxCreateDelegation", nil)
	cdc.RegisterConcrete(&txs2.TxUnbondDelegation{}, "stake/txs/TxUnbondDelegation", nil)
	cdc.RegisterConcrete(&txs2.TxCreateReDelegation{}, "stake/txs/TxCreateReDelegation", nil)

	cdc.RegisterConcrete(&qtypes.QOSAccount{}, "qos/types/QOSAccount", nil)
	//cdc.RegisterConcrete(&BaseAccount{}, "qbase/account/BaseAccount", nil)
	cdc.RegisterInterface((*account.Account)(nil), nil)
	cdc.RegisterConcrete(&account.BaseAccount{}, "qbase/account/BaseAccount", nil)

	//cdc.RegisterConcrete(&txs4.TxCreateApprove{}, "approve/txs/TxCreateApprove", nil)
	//cdc.RegisterConcrete(&txs4.TxIncreaseApprove{}, "approve/txs/TxIncreaseApprove", nil)
	//cdc.RegisterConcrete(&txs4.TxDecreaseApprove{}, "approve/txs/TxDecreaseApprove", nil)
	//cdc.RegisterConcrete(&txs4.TxUseApprove{}, "approve/txs/TxUseApprove", nil)
	//cdc.RegisterConcrete(&txs4.TxCancelApprove{}, "approve/txs/TxCancelApprove", nil)

	cdc.RegisterInterface((*gov_types.ProposalContent)(nil), nil)
	cdc.RegisterConcrete(&gov_types.TextProposal{}, "gov/TextProposal", nil)
	cdc.RegisterConcrete(&gov_types.TaxUsageProposal{}, "gov/TaxUsageProposal", nil)
	cdc.RegisterConcrete(&gov_types.ParameterProposal{}, "gov/ParameterProposal", nil)
	//cdc.RegisterConcrete(&gov_types.ModifyInflationProposal{}, "gov/ModifyInflationProposal", nil)
	cdc.RegisterConcrete(&gov_types.SoftwareUpgradeProposal{}, "gov/SoftwareUpgradeProposal", nil)

	cdc.RegisterConcrete(&InvestTx{}, "qstars/InvestTx", nil)
	cdc.RegisterConcrete(&AdvertisersTx{}, "jianqian/AdvertisersTx", nil)
	cdc.RegisterConcrete(&AuctionTx{}, "jianqian/AuctionTx", nil)
	cdc.RegisterConcrete(&ExtractTx{}, "jianqian/ExtractTx", nil)
	cdc.RegisterConcrete(&JianQianTx{}, "jianqian/JianQianTx", nil)
}

// amino codec to marshal/unmarshal
//var typeRegistry = make(map[string]reflect.Type)
//var Cdc *amino.Codec
//
//type ABCICodeType uint32
//type Tags cmn.KVPairs
//
//type QstarsBaseApp struct {
//	Transactions    BaseXTransaction
//	Baseapp         *baseabci.BaseApp
//	TransactionList []BaseXTransaction
//	Logger          log.Logger
//	RootDir         string
//}
//type BaseXTransaction interface {
//	mapper.IMapper
//	RegisterCdc(cdc *amino.Codec)
//	StartX(base *QstarsBaseApp) error
//}
//type Result struct {
//
//	// Code is the response code, is stored back on the chain.
//	Code ABCICodeType
//
//	// Data is any data returned from the app.
//	Data []byte
//
//	// Log is just debug information. NOTE: nondeterministic.
//	Log string
//
//	// GasWanted is the maximum units of work we allow this tx to perform.
//	GasWanted int64
//
//	// GasUsed is the amount of gas actually consumed. NOTE: unimplemented
//	GasUsed int64
//
//	// Tx fee amount and denom.
//	FeeAmount int64
//	FeeDenom  string
//
//	// Tags are used for transaction indexing and pubsub.
//	Tags Tags
//}
//
//func MakeCodec() *amino.Codec {
//	cdc := MakeQBaseCodec()
//	for k, _ := range typeRegistry {
//		txs, err := newStruct(k)
//		if err == false {
//			panic("reflect transaction is error.")
//		}
//		t := txs.(BaseXTransaction)
//		t.RegisterCdc(cdc)
//	}
//	//kvstore.NewKVStub().RegisterKVCdc(cdc)
//	//bank.NewBankStub().RegisterKVCdc(cdc)
//	return cdc
//}
//
////
//func newStruct(name string) (interface{}, bool) {
//	elem, ok := typeRegistry[name]
//	if !ok {
//		return nil, false
//	}
//	return reflect.New(elem).Elem().Interface(), true
//}
//
//func init() {
//	Cdc = MakeCodec()
//}
//
//func MakeQBaseCodec() *amino.Codec {
//
//	var cdc = amino.NewCodec()
//	//RegisterAmino(cdc)
//	RegisterCodec(cdc)
//
//	return cdc
//}
//
//func RegisterCodec(cdc *amino.Codec) {
//	//txs.RegisterCodec(cdc)
//	//account.RegisterCodec(cdc)
//}

//func RegisterAmino(cdc *amino.Codec) {
//	// These are all written here instead of
//	cdc.RegisterInterface((*crypto.PubKey)(nil), nil)
//	cdc.RegisterConcrete(ed25519.PubKeyEd25519{},
//		"tendermint/PubKeyEd25519", nil)
//
//	cdc.RegisterInterface((*crypto.PrivKey)(nil), nil)
//	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{},
//		"tendermint/PrivKeyEd25519", nil)
//
//}
