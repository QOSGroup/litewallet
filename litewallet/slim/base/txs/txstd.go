package txs

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/pkg/errors"
)

type ITx interface {
	GetSignData() []byte //获取签名字段
}

type TxStd struct {
	ITx       ITx          `json:"itx"`      //ITx接口，将被具体Tx结构实例化
	Signature []Signature  `json:"sigature"` //签名数组
	ChainID   string       `json:"chainid"`  //ChainID: 执行ITx.exec方法的链ID
	MaxGas    types.BigInt `json:"maxgas"`   //Gas消耗的最大值
}

var _ types.Tx = (*TxStd)(nil)

type Signature struct {
	Pubkey    ed25519local.PubKey `json:"pubkey"`    //可选
	Signature []byte              `json:"signature"` //签名内容
	Nonce     int64               `json:"nonce"`     //nonce的值
}

// Type: just for implements types.Tx
func (tx *TxStd) Type() string {
	return "txstd"
}

func (tx *TxStd) GetSignData() []byte {
	if tx.ITx == nil {
		panic("ITx shouldn't be nil in TxStd.GetSignData()")
		return nil
	}

	ret := tx.ITx.GetSignData()
	ret = append(ret, []byte(tx.ChainID)...)
	ret = append(ret, types.Int2Byte(tx.MaxGas.Int64())...)

	return ret
}

// 签名：每个签名者外部调用此方法
func (tx *TxStd) SignTx(privkey ed25519local.PrivKey, nonce int64, fromChainID string) (signedbyte []byte, err error) {
	if tx.ITx == nil {
		return nil, errors.New("Signature txstd err(itx is nil)")
	}

	bz := tx.BuildSignatureBytes(nonce, fromChainID)
	signedbyte, err = privkey.Sign(bz)
	if err != nil {
		return nil, err
	}

	return
}

func (tx *TxStd) BuildSignatureBytes(nonce int64, qcpFromChainID string) []byte {
	bz := tx.getSignData()
	bz = append(bz, types.Int2Byte(nonce)...)
	bz = append(bz, []byte(qcpFromChainID)...)

	return bz
}

func (tx *TxStd) getSignData() []byte {
	if tx.ITx == nil {
		panic("ITx shouldn't be nil in TxStd.GetSignData()")
	}

	ret := tx.ITx.GetSignData()
	ret = append(ret, []byte(tx.ChainID)...)
	ret = append(ret, types.Int2Byte(tx.MaxGas.Int64())...)

	return ret
}

// 调用 NewTxStd后，需调用TxStd.SignTx填充TxStd.Signature(每个TxStd.Signer())
func NewTxStd(itx ITx, cid string, mgas types.BigInt) (rTx *TxStd) {
	rTx = &TxStd{
		itx,
		[]Signature{},
		cid,
		mgas,
	}

	return
}
