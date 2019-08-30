package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
)

type BaseAccount struct {
	AccountAddress types.Address       `json:"account_address"` // account address
	Publickey      ed25519local.PubKey `json:"public_key"`      // public key
	Nonce          int64               `json:"nonce"`           // identifies tx_status of an account
}

type QOSAccount struct {
	BaseAccount `json:"base_account"`
	QOS         types.BigInt `json:"qos"`  // coins in public chain
	QSCs        types.QSCs   `json:"qscs"` // varied QSCs
}
