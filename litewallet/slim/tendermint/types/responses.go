package types

import (
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/types"
)

// Result of querying for a tx
type ResultTx struct {
	Hash     cmn.HexBytes      `json:"hash"`
	Height   int64             `json:"height"`
	Index    uint32            `json:"index"`
	TxResult ResponseDeliverTx `json:"tx_result"`
	Tx       types.Tx          `json:"tx"`
	Proof    types.TxProof     `json:"proof,omitempty"`
}

type ResponseDeliverTx struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log                  string   `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string   `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted            int64    `protobuf:"varint,5,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed              int64    `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Events               []Event  `protobuf:"bytes,7,rep,name=events" json:"events,omitempty"`
	Codespace            string   `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type Event struct {
	Type                 string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes           []common.KVPair `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
