package client

import (
	"testing"
)

func TestQSCCreateSignedTransfer(t *testing.T) {
	//txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	addrto := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	chainid := "qos-test"
	Tout, err := CreateSignedTransfer(addrto, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
	}
	t.Log(Tout)
}
