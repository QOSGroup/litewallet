package types

import (
	"encoding/hex"
	"encoding/json"
)

const (
	ResultCodeSuccess       = "0"
	ResultCodeInternalError = "500"
)

type ResultInvest struct {
	Code   string          `json:"code"`
	Height int64           `json:"height"`
	Hash   string          `json:"hash,omitempty"`
	Reason string          `json:"reason,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func InternalError(reason string) ResultInvest {
	return NewErrorResult(ResultCodeInternalError, 0, "", reason)
}

func NewErrorResult(code string, height int64, hash string, reason string) ResultInvest {
	var result ResultInvest
	result.Height = height
	result.Hash = hash
	result.Code = code
	result.Reason = reason

	return result
}

func (ri ResultInvest) Marshal() string {
	//jsonBytes, err := json.MarshalIndent(ri, "", "  ")
	//if err != nil {
	//	fmt.Printf("InvestAd err:%s", err.Error())
	//	return InternalError(err.Error()).Marshal()
	//}
	if ri.Code == ResultCodeSuccess {
		return string(hex.EncodeToString(ri.Result))
	}
	return string(ri.Result)
}
