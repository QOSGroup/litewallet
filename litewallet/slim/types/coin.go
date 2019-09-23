package types

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
)

const (
	QOSCoinName = "QOS"
)

type QSC = types.BaseCoin

func NewQSC(name string, amount types.BigInt) *QSC {
	return &QSC{
		name, amount,
	}
}

type QSCs = types.BaseCoins
