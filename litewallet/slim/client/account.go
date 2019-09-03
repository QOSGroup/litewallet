package client

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
)

//const (
//	AccountMapperName = "acc"      // 用户获取账户存储的store的键名
//	accountStoreKey   = "account:" // 便于获取全部账户的通用存储键名，继承BaseAccount时，可根据不同业务设置存储前缀
//)

var (
	ErrAccountNotExsits = errors.New("account not exists")
)

func GetAccountFromBech32Addr(bech32Addr string) (*txs.QOSAccount, error) {
	addrBytes, err := types.GetAddrFromBech32(bech32Addr)

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid bech32Addr", bech32Addr)
	}

	//return queryAccount(addrBytes)
	return txs.RpcQueryAccount(addrBytes)
}

//func queryAccount(addr []byte) (*txs.QOSAccount, error) {
//	path := BuildAccountStoreQueryPath()
//	res, err := Query(string(path), AddressStoreKey(addr))
//	if err != nil {
//		return nil, err
//	}
//
//	if len(res) == 0 {
//		return nil, ErrAccountNotExsits
//	}
//
//	var acc *txs.QOSAccount
//	err = txs.Cdc.UnmarshalBinaryBare(res, &acc)
//	if err != nil {
//		return nil, err
//	}
//
//	return acc, nil
//}
//
//func BuildAccountStoreQueryPath() []byte {
//	return []byte(fmt.Sprintf("/store/%s/key", AccountMapperName))
//}
//
//// 将地址转换成存储通用的key
//func AddressStoreKey(addr types.Address) []byte {
//	return append([]byte(accountStoreKey), addr.Bytes()...)
//}
//
//func Query(path string, data []byte) (res []byte, err error) {
//	return query(path, common.HexBytes(data))
//}
//
//// query performs a query from a Tendermint node with the provided store name
//// and path.
//func query(path string, key common.HexBytes) (res []byte, err error) {
//	//todo
//	//node, err := GetNode()
//	//if err != nil {
//	//	return res, err
//	//}
//
//	opts := rpcclient.ABCIQueryOptions{
//		Height: int64(0),
//		Prove:  true,
//	}
//
//	result, err := txs.RPC.ABCIQueryWithOptions(path, key, opts)
//	if err != nil {
//		return res, err
//	}
//
//	resp := result.Response
//	if !resp.IsOK() {
//		return res, fmt.Errorf(resp.Log)
//	}
//
//	return resp.Value, nil
//}

//func GetAddrFromValue(value string) (types.Address, error) {
//	if strings.HasPrefix(value, types.PREF_ADD) {
//		addr, err := types.GetAddrFromBech32(value)
//		if err == nil {
//			return addr, nil
//		}
//	}
//
//	info, err := keys.GetKeyInfo(ctx, value)
//	if err != nil {
//		return nil, err
//	}
//
//	return info.GetAddress(), nil
//}
