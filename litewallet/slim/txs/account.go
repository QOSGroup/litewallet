package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto"
)

type BaseAccount struct {
	AccountAddress types.Address `json:"account_address"` // account address
	Publickey      crypto.PubKey `json:"public_key"`      // public key
	Nonce          int64         `json:"nonce"`           // identifies tx_status of an account
}

type QOSAccount struct {
	BaseAccount `json:"base_account"`
	QOS         types.BigInt `json:"qos"`  // coins in public chain
	QSCs        types.QSCs   `json:"qscs"` // varied QSCs
}

func (account *QOSAccount) GetQOS() types.BigInt {
	return account.QOS.NilToZero()
}

// 设置QOS，币值必须大于等于0
func (account *QOSAccount) SetQOS(qos types.BigInt) error {
	//if qos.LT(types.ZeroInt()) {
	//	return errors.New("qos must gte zero")
	//}

	account.QOS = qos

	return nil
}
