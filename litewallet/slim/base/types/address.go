package types

import (
	"encoding/json"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/pkg/errors"
)

type Address []byte

func (add Address) Bytes() []byte {
	return add[:]
}

func (add Address) String() string {
	bech32Addr, err := bech32local.ConvertAndEncode("address", add.Bytes())
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
