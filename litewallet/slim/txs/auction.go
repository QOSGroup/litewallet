package txs

import "github.com/QOSGroup/litewallet/litewallet/slim/base/types"

type AuctionTx struct {
	ArticleHash string        // 文章hash
	Address     types.AccAddress //qos地址
	CoinsType   string        //币种
	CoinAmount  types.BigInt  //数量
	Gas         types.BigInt
}

func (tx AuctionTx) GetSignData() (ret []byte) {
	ret = append(ret, tx.Address...)
	ret = append(ret, tx.CoinsType...)
	ret = append(ret, []byte(tx.CoinAmount.String())...)
	return ret
}
