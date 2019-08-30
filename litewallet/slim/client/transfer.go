package client

import (
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	ctypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"regexp"
	"strconv"
	"strings"
)

func CreateSignedTransfer(addrto, coinstr, privkey, chainid string) (string, error) {
	return ctxs.BuildAndSignTx(privkey, chainid, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode("address", key.PubKey().Address().Bytes())

		sendersStr := addrben32 + `,` + coinstr
		senders, err := ParseTransItem(sendersStr)
		if err != nil {
			return nil, err
		}

		receiversStr := addrto + `,` + coinstr
		receivers, err := ParseTransItem(receiversStr)
		if err != nil {
			return nil, err
		}
		return ctxs.TxTransfer{
			Senders:   senders,
			Receivers: receivers,
		}, nil
	})
}

// Parse flags from string, Senders, eg: Arya,10qos,100qstar. multiple users separated by ';'
func ParseTransItem(str string) (ctypes.TransItems, error) {
	items := make(ctypes.TransItems, 0)
	tis := strings.Split(str, ";")
	for _, ti := range tis {
		if ti == "" {
			continue
		}

		addrAndCoins := strings.Split(ti, ",")
		if len(addrAndCoins) < 2 {
			return nil, fmt.Errorf("`%s` not match rules", ti)
		}

		addr, err := types.GetAddrFromBech32(addrAndCoins[0])
		if err != nil {
			return nil, err
		}
		qos, qscs, err := NewParseCoins(strings.Join(addrAndCoins[1:], ","))
		if err != nil {
			return nil, err
		}
		items = append(items, ctypes.TransItem{
			Address: addr,
			QOS:     qos,
			QSCs:    qscs,
		})
	}

	return items, nil
}

// Parse QOS and QSCs from string
// str example : 100qos,100qstar
func NewParseCoins(str string) (types.BigInt, types.QSCs, error) {
	if len(str) == 0 {
		return types.ZeroInt(), types.QSCs{}, nil
	}
	reDnm := `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt := `[[:digit:]]+`
	reSpc := `[[:space:]]*`
	reCoin := regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))

	arr := strings.Split(str, ",")
	qos := types.ZeroInt()
	qscs := types.QSCs{}
	for _, q := range arr {
		coin := reCoin.FindStringSubmatch(q)
		if len(coin) != 3 {
			return types.ZeroInt(), nil, fmt.Errorf("coins str: %s parse faild", q)
		}
		coin[2] = strings.TrimSpace(coin[2])
		amount, err := strconv.ParseInt(strings.TrimSpace(coin[1]), 10, 64)
		if err != nil {
			return types.ZeroInt(), nil, err
		}
		if strings.ToLower(coin[2]) == "qos" {
			qos = types.NewInt(amount)
		} else {
			qscs = append(qscs, &types.QSC{
				coin[2],
				types.NewInt(amount),
			})
		}

	}

	return qos, qscs, nil
}
