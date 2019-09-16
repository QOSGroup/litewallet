package module

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"testing"
)

func TestQueryTx(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	hash := "B5EECB27939D969556C61E00BDE8C910FFE3BE47BFA0356B57F58847D6502B70"
	Tout, err := QueryTx("47.103.78.91:26657", hash)
	if err != nil {
		t.Log(err)
	}
	t.Log(Tout)
}
