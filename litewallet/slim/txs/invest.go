package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

type InvestTx struct {
	Address     types.Address `json:"address"`     // 投资者地址
	Invest      types.BigInt  `json:"investad"`    // 投资金额
	ArticleHash []byte        `json:"articleHash"` // 文章hash
	Gas         types.BigInt
	cointype    string
}

var _ txs.ITx = (*InvestTx)(nil)

func (it InvestTx) GetSignData() (ret []byte) {
	ret = append(ret, it.ArticleHash...)
	ret = append(ret, it.Address.Bytes()...)
	ret = append(ret, []byte(it.cointype)...)
	ret = append(ret, types.Int2Byte(it.Invest.Int64())...)
	return
}
