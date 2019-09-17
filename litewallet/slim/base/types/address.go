package types

import (
	"bytes"
	"encoding/json"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/pkg/errors"
)

const (
	PREF_ADD = "address"
)

type Address []byte

func (add Address) Bytes() []byte {
	return add[:]
}

// 判断地址是否为空
func (add Address) Empty() bool {
	if len(add[:]) == 0 {
		return true
	}
	return false
}

// 判断两地址是否相同
func (add Address) EqualsTo(anotherAdd Address) bool {
	if add.Empty() && anotherAdd.Empty() {
		return true
	}
	return bytes.Compare(add.Bytes(), anotherAdd.Bytes()) == 0
}

func (add Address) String() string {
	bech32Addr, err := bech32local.ConvertAndEncode(PREF_ADD, add.Bytes())
	if err != nil {
		panic(err)
	}
	return bech32Addr
}

func (add Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(add.String())
}

// 将Bech32编码的地址Json进行UnMarshal
func (add *Address) UnmarshalJSON(bech32Addr []byte) error {
	var s string
	err := json.Unmarshal(bech32Addr, &s)
	if err != nil {
		return err
	}
	add2, err := GetAddrFromBech32(s)
	if err != nil {
		return err
	}
	*add = add2
	return nil
}

func GetAddrFromBech32(bech32Addr string) ([]byte, error) {
	if len(bech32Addr) == 0 {
		return nil, errors.New("decoding bech32 address failed: must provide an address")
	}
	prefix, bz, err := bech32local.DecodeAndConvert(bech32Addr)
	if prefix != "address" {
		return nil, errors.Wrap(err, "Invalid address prefix!")
	}
	return bz, err
}
