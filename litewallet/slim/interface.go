package slim

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

)
//the MaxGas under configing

const (
	AccountMapperName = "acc"      // 用户获取账户存储的store的键名
	accountStoreKey   = "account:" // 便于获取全部账户的通用存储键名，继承BaseAccount时，可根据不同业务设置存储前缀
    MaxGas = 20000

)


//genStdSendTx for the Tx send operation
// NewInt constructs BigInt from int64
func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}

// NewInt constructs BigInt from int64
//func NewInt(n int64) BigInt {
//	return BigInt{big.NewInt(n)}
//}

func (i BigInt) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

type BigInt struct {
	i *big.Int
}

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

// Add adds BigInt from another
func (i BigInt) Add(i2 BigInt) (res BigInt) {
	res = BigInt{add(i.i, i2.i)}
	// Check overflow
	if res.i.BitLen() > 255 {
		panic("BigInt overflow")
	}
	return
}

func (bi BigInt) IsNil() bool {
	return bi.i == nil
}

func (i BigInt) NilToZero() BigInt {
	if i.IsNil() {
		return ZeroInt()
	}
	return i
}

// ZeroInt returns BigInt value with zero
func ZeroInt() BigInt { return BigInt{big.NewInt(0)} }

func (i BigInt) String() string {
	return i.i.String()
}

// MarshalAmino defines custom encoding scheme
func (i BigInt) MarshalAmino() (string, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalAmino(i.i)
}

// UnmarshalAmino defines custom decoding scheme
func (i *BigInt) UnmarshalAmino(text string) error {
	if i.i == nil { // Necessary since default BigInt initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalAmino(i.i, text)
}

// MarshalJSON defines custom encoding scheme
func (i BigInt) MarshalJSON() ([]byte, error) {
	if i.i == nil { // Necessary since default Uint initialization has i.i as nil
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}

// UnmarshalJSON defines custom decoding scheme
func (i *BigInt) UnmarshalJSON(bz []byte) error {
	if i.i == nil { // Necessary since default BigInt initialization has i.i as nil
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
}

// MarshalAmino for custom encoding scheme
func marshalAmino(i *big.Int) (string, error) {
	bz, err := i.MarshalText()
	return string(bz), err
}

// UnmarshalAmino for custom decoding scheme
func unmarshalAmino(i *big.Int, text string) (err error) {
	return i.UnmarshalText([]byte(text))
}

// MarshalJSON for custom encoding scheme
// Must be encoded as a string for JSON precision
func marshalJSON(i *big.Int) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalJSON for custom decoding scheme
// Must be encoded as a string for JSON precision
func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}

// 函数：int64 转化为 []byte
func Int2Byte(in int64) []byte {
	var ret = bytes.NewBuffer([]byte{})
	err := binary.Write(ret, binary.BigEndian, in)
	if err != nil {
		log.Printf("Int2Byte error:%s", err.Error())
		return nil
	}

	return ret.Bytes()
}

type BaseCoin struct {
	Name   string `json:"coin_name"`
	Amount BigInt `json:"amount"`
}

type TxStd struct {
	ITx       ITx         `json:"itx"`      //ITx接口，将被具体Tx结构实例化
	Signature []Signature `json:"sigature"` //签名数组
	ChainID   string      `json:"chainid"`  //ChainID: 执行ITx.exec方法的链ID
	MaxGas    BigInt      `json:"maxgas"`   //Gas消耗的最大值
}

func (tx *TxStd) GetSignData() []byte {
	if tx.ITx == nil {
		panic("ITx shouldn't be nil in TxStd.GetSignData()")
		return nil
	}

	ret := tx.ITx.GetSignData()
	ret = append(ret, []byte(tx.ChainID)...)
	ret = append(ret, Int2Byte(tx.MaxGas.Int64())...)

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
	bz = append(bz, Int2Byte(nonce)...)
	bz = append(bz, []byte(qcpFromChainID)...)

	return bz
}

func (tx *TxStd) getSignData() []byte {
	if tx.ITx == nil {
		panic("ITx shouldn't be nil in TxStd.GetSignData()")
	}

	ret := tx.ITx.GetSignData()
	ret = append(ret, []byte(tx.ChainID)...)
	ret = append(ret, Int2Byte(tx.MaxGas.Int64())...)

	return ret
}

type ITx interface {
	GetSignData() []byte //获取签名字段
}

//var _ txs.ITx = (*TransferTx)(nil)

type Signature struct {
	Pubkey    ed25519local.PubKey `json:"pubkey"`    //可选
	Signature []byte              `json:"signature"` //签名内容
	Nonce     int64               `json:"nonce"`     //nonce的值
}

// 调用 NewTxStd后，需调用TxStd.SignTx填充TxStd.Signature(每个TxStd.Signer())
func NewTxStd(itx ITx, cid string, mgas BigInt) (rTx *TxStd) {
	rTx = &TxStd{
		itx,
		[]Signature{},
		cid,
		mgas,
	}

	return
}

func genStdSendTx(sendTx ITx, priKey ed25519local.PrivKeyEd25519, chainid string, nonce int64) *TxStd {
	gas := NewBigInt(int64(MaxGas))
	stx := NewTxStd(sendTx, chainid, gas)
	signature, _ := stx.SignTx(priKey, nonce, chainid)
	stx.Signature = []Signature{Signature{
		Pubkey:    priKey.PubKey(),
		Signature: signature,
		Nonce:     nonce,
	}}

	return stx
}

// 将地址转换成存储通用的key
func AddressStoreKey(addr Address) []byte {
	return append([]byte(accountStoreKey), addr.Bytes()...)
}

func getAddrFromBech32(bech32Addr string) ([]byte, error) {
	if len(bech32Addr) == 0 {
		return nil, errors.New("decoding bech32 address failed: must provide an address")
	}
	prefix, bz, err := bech32local.DecodeAndConvert(bech32Addr)
	if prefix != "address" {
		return nil, errors.Wrap(err, "Invalid address prefix!")
	}
	return bz, err
}

type Address []byte

func (add Address) Bytes() []byte {
	return add[:]
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
	add2, err := getAddrFromBech32(s)
	if err != nil {
		return err
	}
	*add = add2
	return nil
}

type BaseCoins []*BaseCoin
type QSCs = BaseCoins
type QSC = BaseCoin

func (coins BaseCoins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}

func (coin *BaseCoin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Name)
}

type TransItem struct {
	Address Address `json:"addr"` // 账户地址
	QOS     BigInt  `json:"qos"`  // QOS
	QSCs    QSCs    `json:"qscs"` // QSCs
}

type TransItems []TransItem

type TxTransfer struct {
	Senders   TransItems `json:"senders"`   // 发送集合
	Receivers TransItems `json:"receivers"` // 接收集合
}

//type TxTransfer struct {
//	Senders   []TransItem `json:"senders"`   // 发送集合
//	Receivers []TransItem `json:"receivers"` // 接收集合
//}

// 签名字节
func (tx TxTransfer) GetSignData() (ret []byte) {
	for _, sender := range tx.Senders {
		ret = append(ret, sender.Address...)
		ret = append(ret, (sender.QOS.NilToZero()).String()...)
		ret = append(ret, sender.QSCs.String()...)
	}
	for _, receiver := range tx.Receivers {
		ret = append(ret, receiver.Address...)
		ret = append(ret, (receiver.QOS.NilToZero()).String()...)
		ret = append(ret, receiver.QSCs.String()...)
	}

	return ret
}

func warpperTransItem(addr Address, coins []BaseCoin) TransItem {
	var ti TransItem
	ti.Address = addr
	ti.QOS = NewBigInt(0)

	for _, coin := range coins {
		if coin.Name == "qos" {
			ti.QOS = ti.QOS.Add(coin.Amount)
		} else {
			ti.QSCs = append(ti.QSCs, &coin)
		}
	}

	return ti
}

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
// Parse QOS and QSCs from string
// str example : 100qos,100qstar
func NewParseCoins(str string) (BigInt, QSCs, error) {
	if len(str) == 0 {
		return ZeroInt(), QSCs{}, nil
	}
	reDnm := `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt := `[[:digit:]]+`
	reSpc := `[[:space:]]*`
	reCoin := regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))

	arr := strings.Split(str, ",")
	qos := ZeroInt()
	qscs := QSCs{}
	for _, q := range arr {
		coin := reCoin.FindStringSubmatch(q)
		if len(coin) != 3 {
			return ZeroInt(), nil, fmt.Errorf("coins str: %s parse faild", q)
		}
		coin[2] = strings.TrimSpace(coin[2])
		amount, err := strconv.ParseInt(strings.TrimSpace(coin[1]), 10, 64)
		if err != nil {
			return ZeroInt(), nil, err
		}
		if strings.ToLower(coin[2]) == "qos" {
			qos = NewBigInt(amount)
		} else {
			qscs = append(qscs, &QSC{
				coin[2],
				NewBigInt(amount),
			})
		}

	}

	return qos, qscs, nil
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

func NewBigInt(n int64) BigInt {
	return BigInt{big.NewInt(n)}
}

type BaseAccount struct {
	AccountAddress Address             `json:"account_address"` // account address
	Publickey      ed25519local.PubKey `json:"public_key"`      // public key
	Nonce          int64               `json:"nonce"`           // identifies tx_status of an account
}

type QOSAccount struct {
	BaseAccount `json:"base_account"`
	QOS         BigInt `json:"qos"`  // coins in public chain
	QSCs        QSCs   `json:"qscs"` // varied QSCs
}

//only need the following arguments, it`s enough!
func QSCtransferSendStr(addrto, coinstr, privkey, chainid string) string {
	//var key ed25519local.PrivKeyEd25519
	//ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	//err := Cdc.UnmarshalJSON([]byte(ts), &key)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//priv := key
	//addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	//from, err2 := getAddrFromBech32(addrben32)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	////Get "nonce" from the func RpcQueryAccount
	//acc,_ := RpcQueryAccount(from)
	//var qscnonce int64
	//if acc!=nil{
	//	qscnonce = int64(acc.Nonce)
	//}
	//qscnonce++
	//
	//sendersStr := addrben32 + `,` + coinstr
	//senders, err := ParseTransItem(sendersStr)
	//if err != nil {
	//	return err.Error()
	//}
	//
	//receiversStr := addrto + `,` + coinstr
	//receivers, err := ParseTransItem(receiversStr)
	//if err != nil {
	//	return err.Error()
	//}
	//tn := TxTransfer{
	//	Senders:   senders,
	//	Receivers: receivers,
	//}
	jasonpayload, err := QSCCreateSignedTransfer(addrto, coinstr, privkey, chainid)
	if err != nil {
		fmt.Println(err)
	}
	datas := bytes.NewBuffer([]byte(jasonpayload))
	aurl := Accounturl + "txSend"
	req, _ := http.NewRequest("POST", aurl, datas)
	req.Header.Set("Content-Type", "application/json")
	clt := http.Client{}
	resp, _ := clt.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	output := string(body)
	return output
}

func QSCCreateSignedTransfer(addrto, coinstr, privkey, chainid string) (string, error) {
	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
	err := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err != nil {
		fmt.Println(err)
	}
	priv := key
	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	from, err2 := getAddrFromBech32(addrben32)
	if err2 != nil {
		fmt.Println(err2)
	}
	//Get "nonce" from the func RpcQueryAccount
	acc, _ := RpcQueryAccount(from)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	sendersStr := addrben32 + `,` + coinstr
	senders, err := ParseTransItem(sendersStr)
	if err != nil {
		return "", err
	}

	receiversStr := addrto + `,` + coinstr
	receivers, err := ParseTransItem(receiversStr)
	if err != nil {
		return "", err
	}
	tn := TxTransfer{
		Senders:   senders,
		Receivers: receivers,
	}

	msg := genStdSendTx(tn, priv, chainid, qscnonce)
	jasonpayload, err := Cdc.MarshalJSON(msg)
	if (err != nil) {
		return "", err
	}
	return string(jasonpayload), nil
}
//type InvestTx struct {
//	Std         *TxStd
//	ArticleHash []byte `json:"articleHash"` // 文章hash
//}


type InvestTx struct {
	Address      Address    `json:"address"`      // 投资者地址
	Invest       BigInt     `json:"investad"`     // 投资金额
	ArticleHash  []byte     `json:"articleHash"`  // 文章hash
	Gas          BigInt
	cointype     string
}



func (it InvestTx) GetSignData() (ret []byte) {
	ret = append(ret, it.ArticleHash...)
	ret = append(ret, it.Address.Bytes()...)
	ret = append(ret, []byte(it.cointype)...)
	ret = append(ret, Int2Byte(it.Invest.Int64())...)
	return
}

var _ ITx = (*InvestTx)(nil)

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
	if ri.Code==ResultCodeSuccess{
		return string(	hex.EncodeToString(ri.Result))
	}
	return string(ri.Result)
}

const coinsName = "AOE"

var tempAddr = Address("99999999999999999999")

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

	js, err := Cdc.MarshalBinaryBare(tx)
	if err != nil {
		fmt.Printf("investAd err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)

	return result.Marshal()
}

func investAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) (*TxStd, error) {
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
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	addrben32, _ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, err2 := getAddrFromBech32(addrben32)
	if err2 != nil {
		fmt.Println(err2)
	}
	var ccs []BaseCoin
	for _, coin := range cs {
		ccs = append(ccs, BaseCoin{
			Name:   coin.Denom,
			Amount: NewBigInt(coin.Amount.Int64()),
		})
	}
	//qos nonce fetched from the qosaccount query
	acc,_ := RpcQueryAccount(investor)
	var qscnonce int64
	if acc!=nil{
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	it := &InvestTx{}
	it.ArticleHash = []byte(articleHash)
	it.Address=investor
	it.cointype=ccs[0].Name
	it.Invest=ccs[0].Amount
	gas := NewBigInt(int64(MaxGas))
	tx2 := NewTxStd(it, QSCchainId, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, QSCchainId)
	tx2.Signature = []Signature{Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}

	return tx2, nil
}

type CoinsTx struct {
	Address Address
	Cointype string
	Amount BigInt
	ChangeType string     //0 plus  1 minus
}




//广告商押金或赎回
func Advertisers( amount, privatekey, cointype,isDeposit,qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	tx, berr := advertisers(amount, privatekey, cointype,isDeposit,qscchainid)
	if berr != "" {
		return berr
	}
	js, err := Cdc.MarshalBinaryBare(tx)
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
func advertisers( coins, privatekey, cointype,isDeposit,qscchainid string) (*TxStd, string) {
	amount, err := strconv.Atoi(coins)
	if err!=nil {
		return nil, NewErrorResult("601", 0, "", "amount format error").Marshal()
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32,_ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, _ := getAddrFromBech32(addrben32)

	acc,_ := RpcQueryAccount(investor)
	var qscnonce int64
	if acc!=nil{
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &CoinsTx{}
	it.Address = investor
	it.Cointype=cointype
	it.ChangeType=isDeposit
	it.Amount=NewBigInt(int64(amount))
	tx:=AdvertisersTx{it}
	fmt.Println(investor, amount, cointype, isDeposit)
	tx2 := NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []Signature{Signature{
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
	ret = append(ret, Int2Byte(tx.Tx.Amount.Int64())...)
	ret = append(ret, []byte(tx.Tx.Cointype)...)
	ret = append(ret, []byte(tx.Tx.ChangeType)...)
	return
}


type AuctionTx struct {
	ArticleHash string        // 文章hash
	Address     Address //qos地址
	CoinsType   string        //币种
	CoinAmount  BigInt //数量
	Gas BigInt
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
func AdvertisersTrue( privatekey,  coinsType, coinAmount,qscchainid string) string {
	return Advertisers(coinAmount,privatekey,coinsType,"2",qscchainid)
}

//成为非广告商 赎回押金
//privatekey             //用户私钥
//coinsType              //押金币种
//coinAmount             //押金数量
//qscchainid             //chainid
func AdvertisersFalse( privatekey,  coinsType, coinAmount,qscchainid string) string {
	return Advertisers(coinAmount,privatekey,coinsType,"1",qscchainid)
}


func GetTx(tx string)string{
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
func AcutionAd( articleHash, privatekey,  coinsType, coinAmount,qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	amount, err := strconv.Atoi(coinAmount)
	if err!=nil {
		result.Code = ResultCodeInternalError
		result.Reason = "AcutionAd invalid amount"
		return result.Marshal()
	}
	tx, buff := acutionAd( articleHash, privatekey,  coinsType, amount,qscchainid)
	if buff != "" {
		return buff
	}

	js, err := Cdc.MarshalBinaryBare(tx)
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
func acutionAd(articleHash, privatekey,  coinsType string,coinAmount int,qscchainid string) (*TxStd, string) {

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))
	addrben32,_ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	sendAddress, _ := getAddrFromBech32(addrben32)
	acc,_ := RpcQueryAccount(sendAddress)
	var qscnonce int64
	if acc!=nil{
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &AuctionTx{}
	it.ArticleHash = articleHash
	it.Address = sendAddress
	it.CoinsType = coinsType
	it.Gas =ZeroInt()
	it.CoinAmount = NewBigInt(int64(coinAmount))
	tx2 := NewTxStd(it, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []Signature{Signature{
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
	ret = append(ret, Int2Byte(tx.Tx.Amount.Int64())...)
	ret = append(ret, []byte(tx.Tx.Cointype)...)
	ret = append(ret, []byte(tx.Tx.ChangeType)...)
	return
}



//广告商押金或赎回
func Extract( amount, privatekey, cointype,qscchainid string) string {
	var result ResultInvest
	result.Code = ResultCodeSuccess
	tx, berr := extract(amount, privatekey, cointype,qscchainid)
	if berr != "" {
		return berr
	}
	js, err := Cdc.MarshalBinaryBare(tx)
	if err != nil {
		log.Printf("Extract err:%s", err.Error())
		result.Code = ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}
	result.Result = json.RawMessage(js)
	return result.Marshal()
}


func extract( coins, privatekey, cointype,qscchainid string) (*TxStd, string) {
	amount, err := strconv.Atoi(coins)
	if err!=nil {
		return nil, NewErrorResult("601", 0, "", "amount format error").Marshal()
	}

	var key ed25519local.PrivKeyEd25519
	ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privatekey + "\"}"
	err1 := Cdc.UnmarshalJSON([]byte(ts), &key)
	if err1 != nil {
		fmt.Println(err1)
	}
	priv := key
	gas := NewBigInt(int64(MaxGas))

	addrben32,_ := bech32local.ConvertAndEncode(PREF_ADD, key.PubKey().Address().Bytes())
	investor, _ := getAddrFromBech32(addrben32)

	acc,_ := RpcQueryAccount(investor)
	var qscnonce int64
	if acc!=nil{
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++
	it := &CoinsTx{}
	it.Address = investor
	it.Cointype=cointype
	it.ChangeType="2"
	it.Amount=NewBigInt(int64(amount))
	tx:=ExtractTx{it}
	fmt.Println(investor, amount, cointype, "2")
	tx2 := NewTxStd(tx, qscchainid, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, qscchainid)
	tx2.Signature = []Signature{Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}
	return tx2, ""
}

// Parse flags from string, Senders, eg: Arya,10qos,100qstar. multiple users separated by ';'
func ParseTransItem(str string) (TransItems, error) {
	items := make(TransItems, 0)
	tis := strings.Split(str, ";")
	for _, ti := range tis {
		if ti == "" {
			continue
		}

		addrAndCoins := strings.Split(ti, ",")
		if len(addrAndCoins) < 2 {
			return nil, fmt.Errorf("`%s` not match rules", ti)
		}

		addr, err := getAddrFromBech32(addrAndCoins[0])
		if err != nil {
			return nil, err
		}
		qos, qscs, err := NewParseCoins(strings.Join(addrAndCoins[1:], ","))
		if err != nil {
			return nil, err
		}
		items = append(items, TransItem{
			Address: addr,
			QOS:     qos,
			QSCs:    qscs,
		})
	}

	return items, nil
}
