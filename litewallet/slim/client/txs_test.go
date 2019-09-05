package client

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"testing"
)

func TestQueryTx(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	hash := "20819ED6BE5E999C8F80D058C92918BB96ED212AE38B1AB568D9DE1366A4C74A"
	Tout, err := QueryTx(hash)
	if err != nil {
		t.Log(err)
	}
	t.Log(Tout)
}
