package slim

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"log"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//the MaxGas under configing

const (
	AccountMapperName = "acc"      // 用户获取账户存储的store的键名
	accountStoreKey   = "account:" // 便于获取全部账户的通用存储键名，继承BaseAccount时，可根据不同业务设置存储前缀
	MaxGas            = 20000
)

// NewTransfer ... Deprecated!
//func NewTransfer(sender Address, receiver Address, coin []BaseCoin) ITx {
//	var sendTx TxTransfer
//
//	sendTx.Senders = append(sendTx.Senders, warpperTransItem(sender, coin))
//	sendTx.Receivers = append(sendTx.Receivers, warpperTransItem(receiver, coin))
//
//	return sendTx
//}

func (coins Coins) Len() int           { return len(coins) }
func (coins Coins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }
func (coins Coins) Swap(i, j int)      { coins[i], coins[j] = coins[j], coins[i] }

var _ sort.Interface = Coins{}

type Coins []Coin

func (coins Coins) Sort() Coins {
	sort.Sort(coins)
	return coins
}

func (coins Coins) IsZero() bool {
	for _, coin := range coins {
		if !coin.IsZero() {
			return false
		}
	}
	return true
}

func (coins Coins) IsValid() bool {
	switch len(coins) {
	case 0:
		return true
	case 1:
		return !coins[0].IsZero()
	default:
		lowDenom := coins[0].Denom
		for _, coin := range coins[1:] {
			if coin.Denom <= lowDenom {
				return false
			}
			if coin.IsZero() {
				return false
			}
			// we compare each coin against the last denom
			lowDenom = coin.Denom
		}
		return true
	}
}

func ParseCoins(coinsStr string) (coins Coins, err error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin, err := ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}

	// Sort coins for determinism.
	coins.Sort()

	// Validate coins before returning.
	if !coins.IsValid() {
		return nil, fmt.Errorf("parseCoins invalid: %#v", coins)
	}

	return coins, nil
}

type Int struct {
	i *big.Int
}

func (i Int) IsZero() bool {
	return i.i.Sign() == 0
}

func (i Int) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

//genStdSendTx for the Tx send operation
// NewInt constructs BigInt from int64
func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount Int    `json:"amount"`
}

func (coin Coin) IsZero() bool {
	return coin.Amount.IsZero()
}

var (
	// Denominations can be 3 ~ 16 characters long.
	reDnm  = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt  = `[[:digit:]]+`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

func ParseCoin(coinStr string) (coin Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		err = fmt.Errorf("invalid coin expression: %s", coinStr)
		return
	}
	denomStr, amountStr := matches[2], matches[1]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return
	}

	return Coin{denomStr, NewInt(int64(amount))}, nil
}

func NewBigInt(n int64) types.BigInt {
	return types.NewInt(n)
}

// 将地址转换成存储通用的key
func AddressStoreKey(addr types.Address) []byte {
	return append([]byte(accountStoreKey), addr.Bytes()...)
}

//type InvestTx struct {
//	Std         *TxStd
//	ArticleHash []byte `json:"articleHash"` // 文章hash
//}

type InvestTx struct {
	Address     types.Address `json:"address"`     // 投资者地址
	Invest      types.BigInt  `json:"investad"`    // 投资金额
	ArticleHash []byte        `json:"articleHash"` // 文章hash
	Gas         types.BigInt
	cointype    string
}

func (it InvestTx) GetSignData() (ret []byte) {
	ret = append(ret, it.ArticleHash...)
	ret = append(ret, it.Address.Bytes()...)
	ret = append(ret, []byte(it.cointype)...)
	ret = append(ret, types.Int2Byte(it.Invest.Int64())...)
	return
}

var _ txs.ITx = (*InvestTx)(nil)

const (
	ResultCodeSuccess       = "0"
	ResultCodeQstarsTimeout = "-2"
	ResultCodeQOSTimeout    = "-1"
	ResultCodeInternalError = "500"
)

type ResultInvest struct {
	Code   string          `json:"code"`
	Height int64           `json:"height"`
	Hash   string          `json:"hash,omitempty"`
	Reason string          `json:"reason,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

func InternalError(reason string) ResultInvest {
	return NewErrorResult(ResultCodeInternalError, 0, "", reason)
}

func NewErrorResult(code string, height int64, hash string, reason string) ResultInvest {
	var result ResultInvest
	result.Height = height
	result.Hash = hash
	result.Code = code
	result.Reason = reason

	return result
}

func (ri ResultInvest) Marshal() string {
	//jsonBytes, err := json.MarshalIndent(ri, "", "  ")
	//if err != nil {
	//	fmt.Printf("InvestAd err:%s", err.Error())
	//	return InternalError(err.Error()).Marshal()
	//}
	if ri.Code == ResultCodeSuccess {
		return string(hex.EncodeToString(ri.Result))
	}
	return string(ri.Result)
}

const coinsName = "AOE"

var tempAddr = types.Address("99999999999999999999")

func JQInvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess

	tx, err := investAd(QOSchainId, QSCchainId, articleHash, coins, privatekey)
	if err != nil {
		fmt.Printf("investAd err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}

	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		fmt.Printf("investAd err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)

	return result.Marshal()
}

func investAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) (*txs.TxStd, error) {
	cs, err := ParseCoins(coins)
	if err != nil {
		return nil, err
	}
	for _, v := range cs {
		if v.Denom != coinsName {
			return nil, errors.New("only support AOE")
		}
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, err2 := types.GetAddrFromBech32(addrben32)
	if err2 != nil {
		fmt.Println(err2)
	}
	var ccs []types.BaseCoin
	for _, coin := range cs {
		ccs = append(ccs, types.BaseCoin{
			Name:   coin.Denom,
			Amount: NewBigInt(coin.Amount.Int64()),
		})
	}
	//qos nonce fetched from the qosaccount query
	acc, _ := ctxs.RpcQueryAccount(investor)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	it := &InvestTx{}
	it.ArticleHash = []byte(articleHash)
	it.Address = investor
	it.cointype = ccs[0].Name
	it.Invest = ccs[0].Amount
	gas := NewBigInt(int64(MaxGas))
	tx2 := txs.NewTxStd(it, QSCchainId, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, QSCchainId)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}

	return tx2, nil
}

type CoinsTx struct {
	Address    types.Address
	Cointype   string
	Amount     types.BigInt
	ChangeType string //0 plus  1 minus
}

//广告商押金或赎回
func Advertisers(amount, privatekey, cointype, isDeposit, qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	tx, berr := advertisers(amount, privatekey, cointype, isDeposit, qscchainid)
	if berr != "" {
		return berr
	}
	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("Advertisers err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

// investAd 投资广告
func advertisers(coins, privatekey, cointype, isDeposit, qscchainid string) (*txs.TxStd, string) {
	amount, err := strconv.Atoi(coins)
	if err != nil {
		return nil, NewErrorResult("601", 0, "", "amount format error").Marshal()
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, _ := types.GetAddrFromBech32(addrben32)

	acc, _ := ctxs.RpcQueryAccount(investor)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &CoinsTx{}
	it.Address = investor
	it.Cointype = cointype
	it.ChangeType = isDeposit
	it.Amount = NewBigInt(int64(amount))
	tx := AdvertisersTx{it}
	fmt.Println(investor, amount, cointype, isDeposit)
	tx2 := txs.NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}

type AdvertisersTx struct {
	Tx *CoinsTx
}

func (tx AdvertisersTx) GetSignData() (ret []byte) {
	ret = append(ret, tx.Tx.Address.Bytes()...)
	ret = append(ret, types.Int2Byte(tx.Tx.Amount.Int64())...)
	ret = append(ret, []byte(tx.Tx.Cointype)...)
	ret = append(ret, []byte(tx.Tx.ChangeType)...)
	return
}

type AuctionTx struct {
	ArticleHash string        // 文章hash
	Address     types.Address //qos地址
	CoinsType   string        //币种
	CoinAmount  types.BigInt  //数量
	Gas         types.BigInt
}

func (tx AuctionTx) GetSignData() (ret []byte) {
	ret = append(ret, tx.Address...)
	ret = append(ret, tx.CoinsType...)
	ret = append(ret, []byte(tx.CoinAmount.String())...)
	return ret
}

//成为广告商
//privatekey             //用户私钥
//coinsType              //押金币种
//coinAmount             //押金数量
//qscchainid             //chainid
func AdvertisersTrue(privatekey, coinsType, coinAmount, qscchainid string) string {
	return Advertisers(coinAmount, privatekey, coinsType, "2", qscchainid)
}

//成为非广告商 赎回押金
//privatekey             //用户私钥
//coinsType              //押金币种
//coinAmount             //押金数量
//qscchainid             //chainid
func AdvertisersFalse(privatekey, coinsType, coinAmount, qscchainid string) string {
	return Advertisers(coinAmount, privatekey, coinsType, "1", qscchainid)
}

func GetTx(tx string) string {
	txBytes, err := hex.DecodeString(tx)
	if err != nil {
		return err.Error()
	}
	txhashs := strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes)))
	return string(txhashs)
}

// acutionAd 竞拍广告
//articleHash            //广告位标识
//privatekey             //用户私钥
//coinsType              //竞拍币种
//coinAmount             //竞拍数量
//qscchainid             //chainid
func AcutionAd(articleHash, privatekey, coinsType, coinAmount, qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	amount, err := strconv.Atoi(coinAmount)
	if err != nil {
		result.Code = ResultCodeInternalError
		result.Reason = "AcutionAd invalid amount"
		return result.Marshal()
	}
	tx, buff := acutionAd(articleHash, privatekey, coinsType, amount, qscchainid)
	if buff != "" {
		return buff
	}

	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("AcutionAd err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

// acutionAd 竞拍广告
func acutionAd(articleHash, privatekey, coinsType string, coinAmount int, qscchainid string) (*txs.TxStd, string) {

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))
	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	sendAddress, _ := types.GetAddrFromBech32(addrben32)
	acc, _ := ctxs.RpcQueryAccount(sendAddress)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &AuctionTx{}
	it.ArticleHash = articleHash
	it.Address = sendAddress
	it.CoinsType = coinsType
	it.Gas = types.ZeroInt()
	it.CoinAmount = NewBigInt(int64(coinAmount))
	tx2 := txs.NewTxStd(it, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}

type ExtractTx struct {
	Tx *CoinsTx
}

func (tx ExtractTx) GetSignData() (ret []byte) {
	ret = append(ret, tx.Tx.Address.Bytes()...)
	ret = append(ret, types.Int2Byte(tx.Tx.Amount.Int64())...)
	ret = append(ret, []byte(tx.Tx.Cointype)...)
	ret = append(ret, []byte(tx.Tx.ChangeType)...)
	return
}

//广告商押金或赎回
func Extract(amount, privatekey, cointype, qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	tx, berr := extract(amount, privatekey, cointype, qscchainid)
	if berr != "" {
		return berr
	}
	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("Extract err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}

func extract(coins, privatekey, cointype, qscchainid string) (*txs.TxStd, string) {
	amount, err := strconv.Atoi(coins)
	if err != nil {
		return nil, NewErrorResult("601", 0, "", "amount format error").Marshal()
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, _ := types.GetAddrFromBech32(addrben32)

	acc, _ := ctxs.RpcQueryAccount(investor)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &CoinsTx{}
	it.Address = investor
	it.Cointype = cointype
	it.ChangeType = "2"
	it.Amount = NewBigInt(int64(amount))
	tx := ExtractTx{it}
	fmt.Println(investor, amount, cointype, "2")
	tx2 := txs.NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}
