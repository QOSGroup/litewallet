package types

import (
	"fmt"
)

type BaseCoin struct {
	Name   string `json:"coin_name"`
	Amount BigInt `json:"amount"`
}

func (coin *BaseCoin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Name)
}

func (coins BaseCoins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}
