package client

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"testing"
)

func TestCreateSignedDelegation(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	validatorAddress := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	//coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := CreateSignedDelegation(validatorAddress, 100, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestCreateSignedUnbondDelegation(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	validatorAddress := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	//coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	chainid := "aquarius-1000"
	Tout, err := CreateSignedUnbondDelegation(validatorAddress, 1000, privkey, chainid)
	if err != nil {
		t.Log(err)
	}
	t.Log(Tout)
}
