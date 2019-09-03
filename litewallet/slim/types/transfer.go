package types

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

type TransItem struct {
	Address types.Address `json:"addr"` // 账户地址
	QOS     types.BigInt  `json:"qos"`  // QOS
	QSCs    types.QSCs    `json:"qscs"` // QSCs
}

type TransItems []TransItem

//func warpperTransItem(addr types.Address, coins []types.BaseCoin) TransItem {
//	var ti TransItem
//	ti.Address = addr
//	ti.QOS = types.BigInt{big.NewInt(0)}
//
//	for _, coin := range coins {
//		if coin.Name == "qos" {
//			ti.QOS = ti.QOS.Add(coin.Amount)
//		} else {
//			ti.QSCs = append(ti.QSCs, &coin)
//		}
//	}
//
//	return ti
//}
