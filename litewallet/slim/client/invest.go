package client

import (
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

const coinsName = "AOE"

var (
	// Denominations can be 3 ~ 16 characters long.
	reDnm  = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt  = `[[:digit:]]+`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

func InvestAd(QOSchainId, QSCchainId, articleHash, coins, privatekey string) string {
	var result ctypes.ResultInvest
	result.Code = ctypes.ResultCodeSuccess

	tx, err := investAd(QOSchainId, QSCchainId, articleHash, coins, privatekey)
	if err != nil {
		fmt.Printf("investAd err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
		result.Reason = err.Error()
		return result.Marshal()
	}

	js, err := ctxs.Cdc.MarshalBinaryBare(tx)
	if err != nil {
		fmt.Printf("investAd err:%s", err.Error())
		result.Code = ctypes.ResultCodeInternalError
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
	addrben32, _ := bech32local.ConvertAndEncode(types.PREF_ADD, key.PubKey().Address().Bytes())
	investor, err2 := types.GetAddrFromBech32(addrben32)
	if err2 != nil {
		fmt.Println(err2)
	}
	var ccs []types.BaseCoin
	for _, coin := range cs {
		ccs = append(ccs, types.BaseCoin{
			Name:   coin.Denom,
			Amount: types.NewInt(coin.Amount.Int64()),
		})
	}
	//qos nonce fetched from the qosaccount query
	acc, _ := ctxs.QueryAccount(investor)
	var qscnonce int64
	if acc != nil {
		qscnonce = int64(acc.Nonce)
	}
	qscnonce++

	it := &ctxs.InvestTx{}
	it.ArticleHash = []byte(articleHash)
	it.Address = investor
	//it.cointype = ccs[0].Name
	it.Invest = ccs[0].Amount
	gas := types.NewInt(int64(ctxs.MaxGas))
	tx2 := txs.NewTxStd(it, QSCchainId, gas)
	signature2, _ := tx2.SignTx(priv, qscnonce, QSCchainId)
	tx2.Signature = []txs.Signature{txs.Signature{
		Pubkey:    priv.PubKey(),
		Signature: signature2,
		Nonce:     qscnonce,
	}}

	return tx2, nil
}
func ParseCoins(coinsStr string) (coins ctypes.Coins, err error) {
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
func ParseCoin(coinStr string) (coin ctypes.Coin, err error) {
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

	return ctypes.Coin{denomStr, ctypes.NewInt(int64(amount))}, nil
}
