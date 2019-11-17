package litewallet

import (
	"encoding/json"
	"github.com/QOSGroup/js-keys/keys"
	"github.com/QOSGroup/litewallet/types"
	"os"
	"testing"
)

func TestBind(t *testing.T) {

	path := "db"
	defer os.RemoveAll(path)

	InitWallet("qos-wallet", path)

	response := ProduceMnemonic()
	t.Log(response)

	m := make(map[string]string)
	json.Unmarshal([]byte(response), &m)

	r1 := CreateAccountWithName("Account3", "123456")
	r2 := CreateAccountWithMnemonic("account-2", "qwerty", m["data"])
	r3 := CreateAccount("12345678")
	r31 := CreateAccount("12345678")
	t.Log(r1)
	t.Log(r2)
	t.Log(r3)
	t.Log(r31)

	m1 := make(map[string]interface{})
	json.Unmarshal([]byte(r1), &m1)
	d1 := m1["data"].(map[string]interface{})
	addr1 := d1["address"].(string)

	m2 := make(map[string]interface{})
	json.Unmarshal([]byte(r2), &m2)
	d2 := m2["data"].(map[string]interface{})
	addr2 := d2["address"].(string)

	r4 := GetAccountByName("Account3")
	t.Log(r4)

	r41 := GetAccountByName("Account5")
	t.Log(r41)

	r5 := GetAccount(addr1)
	t.Log(r5)

	r52 := GetAccount(addr2)
	t.Log(r52)

	r6 := ExportAccount(addr1, "123456")
	t.Log(r6)

	r62 := ExportAccount(addr1, "1234567")
	t.Log(r62)

	r63 := ExportAccount(addr2, "qwerty")
	t.Log(r63)

	r64 := ExportAccount(addr2, "1234567")
	t.Log(r64)

	la1 := ListAllAccounts()
	t.Log(la1)

	r7 := DeleteAccount(addr1, "1234567")
	t.Log(r7)

	r72 := DeleteAccount(addr1, "123456")
	t.Log(r72)

	r8 := GetAccount(addr1)
	t.Log(r8)

	la2 := ListAllAccounts()
	t.Log(la2)

	s1 := Sign(addr2, "qwerty", "你好, QOS钱包")
	t.Log(s1)

	s12 := Sign(addr2, "qwerty2", "你好, QOS钱包")
	t.Log(s12)

	s2 := SignBase64(addr2, "qwerty", types.EncodeBase64([]byte("你好, QOS钱包")))
	t.Log(s2)

	r9 := DeleteAccount(addr2, "qwerty")
	t.Log(r9)

	la3 := ListAllAccounts()
	t.Log(la3)

	password := "11111111"
	mn1 := getMnemonic()
	mn2 := getMnemonic()
	mn3 := getMnemonic()

	t.Log(CreateAccountWithMnemonic("", password, mn1))
	t.Log(ImportMnemonic(mn1, password))
	t.Log(ImportMnemonic(mn2, password))
	t.Log(ImportMnemonic(mn3, password))

	mn4 := getMnemonic()
	mn5 := getMnemonic()
	mn6 := getMnemonic()

	pk1, _, _ := keys.DeriveQOSKey(mn4)
	pk2, _, _ := keys.DeriveQOSKey(mn5)
	pk3, _, _ := keys.DeriveQOSKey(mn6)

	t.Log(CreateAccountWithMnemonic("", password, mn4))
	t.Log(ImportPrivateKey(types.EncodeHex(pk1), password))
	t.Log(ImportPrivateKey(types.EncodeHex(pk2), password))
	t.Log(ImportPrivateKey(types.EncodeHex(pk3), password))
}

func getMnemonic() string {
	response := ProduceMnemonic()
	m := make(map[string]string)
	json.Unmarshal([]byte(response), &m)
	return m["data"]
}
