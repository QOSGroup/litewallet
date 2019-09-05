package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	tcommon "github.com/tendermint/tendermint/libs/common"
)

// qos端对TxQcp的执行结果
type QcpTxResult struct {
	Result              types.Result `json:"result"`              //对应TxQcp执行结果
	QcpOriginalSequence int64        `json:"qcporiginalsequence"` //此结果对应的TxQcp.Sequence
	QcpOriginalExtends  string       `json:"qcpextends"`          //此结果对应的 TxQcp.Extends
	Info                string       `json:"info"`                //结果信息
}

var _ ITx = (*QcpTxResult)(nil)

func (tx *QcpTxResult) IsOk() bool {
	return tx.Result.Code.IsOK()
}

// 获取签名字段
func (tx *QcpTxResult) GetSignData() []byte {
	ret := types.Int2Byte(int64(tx.Result.Code))
	ret = append(ret, tx.Result.Data...)
	//for _, event := range tx.Result.Events{
	//	ret = append(ret, []byte(event.Type)...)
	//	ret = append(ret, []byte(Extends2Byte(event.Attributes))...)
	//}
	ret = append(ret, types.Int2Byte(int64(tx.Result.GasUsed))...)
	ret = append(ret, types.Int2Byte(tx.QcpOriginalSequence)...)
	ret = append(ret, []byte(tx.QcpOriginalExtends)...)
	ret = append(ret, []byte(tx.Info)...)

	return ret
}

// 功能：构建 QcpTxReasult 结构体
func NewQcpTxResult(result types.Result, qcpSequence int64, qcpExtends, info string) *QcpTxResult {
	return &QcpTxResult{
		Result:              result,
		QcpOriginalSequence: qcpSequence,
		QcpOriginalExtends:  qcpExtends,
		Info:                info,
	}
}

// 功能：将common.KVPair转化成[]byte
// todo: test（amino序列化及反序列化的正确性）
func Extends2Byte(ext []tcommon.KVPair) (ret []byte) {
	if ext == nil || len(ext) == 0 {
		return nil
	}

	for _, kv := range ext {
		ret = append(ret, kv.GetKey()...)
		ret = append(ret, kv.GetValue()...)
	}

	return ret
}
