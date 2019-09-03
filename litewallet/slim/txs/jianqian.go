package txs

import "github.com/QOSGroup/litewallet/litewallet/slim/base/types"

type JianQianTx struct {
	FuncName string          //方法名 路由用
	Address  []types.Address //签名者地址
	Args     []string        //参数列表
	Gas      types.BigInt
}

func (it JianQianTx) GetSignData() (ret []byte) {
	ret = append(ret, []byte(it.FuncName)...)
	for _, v := range it.Address {
		ret = append(ret, v.Bytes()...)
	}
	for _, v := range it.Args {
		ret = append(ret, []byte(v)...)
	}
	return
}
