package client

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"testing"
)

func TestQueryApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	Tout, err := QueryApprove(toAddr, privkey)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestCreateApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := CreateApprove(toAddr, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestIncreaseApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := IncreaseApprove(toAddr, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestDecreaseApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := DecreaseApprove(toAddr, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestUseApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := UseApprove(toAddr, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}

func TestCancelApprove(t *testing.T) {
	txs.SetBlockchainEntrance("47.103.78.91:26657", "forQmoonAddr")
	toAddr := "address1nzv9awha9606jp5rpqe2kujckddpyauggu56ru"
	coinstr := "10000qos"
	privkey := "GkPnj9kRit0IMQNyY3k3KgQapl4l0o1hCQg4yqk1iw0kyVH28bOOMahIjzKOnUPLgv7A5fX3wQjV6qPdGOWeVA=="
	chainid := "aquarius-1000"
	Tout, err := CancelApprove(toAddr, coinstr, privkey, chainid)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(Tout)
}
