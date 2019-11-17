package account

import (
	"github.com/QOSGroup/litewallet/types"
)

var (
	accountIndexKey      = []byte("AccountIndex") //value: id
	accountPrefixKey     = []byte("acc:")         //key: acc:[id] value: account
	accountNamePrefixKey = []byte("acc-name:")    //key: acc-name:[name] value: id
	accountAddrPrefixKey = []byte("acc-addr:")    //key: acc-addr:[addr] value: id
	accountSaltPrefixKey = []byte("acc-salt:")    //key: acc-salt:[id] value: salt
)

func AccountIndexKey() []byte {
	return accountIndexKey
}

func AccountKey(id uint64) []byte {
	return append(accountPrefixKey, types.Uint64ToBytes(id)...)
}

func AccountNameKey(name string) []byte {
	return append(accountNamePrefixKey, []byte(name)...)
}

func AccountAddressKey(addr string) []byte {
	return append(accountAddrPrefixKey, []byte(addr)...)
}

func AccountSaltKey(id uint64) []byte {
	return append(accountSaltPrefixKey, types.Uint64ToBytes(id)...)
}
