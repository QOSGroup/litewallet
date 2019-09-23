package client

import (
	"fmt"
	qcliacc "github.com/QOSGroup/litewallet/litewallet/slim/base/client/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/tx"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/txs"
	btypes "github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	bank_txs "github.com/QOSGroup/litewallet/litewallet/slim/module/bank/txs"
	"github.com/QOSGroup/litewallet/litewallet/slim/module/bank/types"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/bech32local"
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/crypto/funcInlocal/ed25519local"
	ctxs "github.com/QOSGroup/litewallet/litewallet/slim/txs"
	qtypes "github.com/QOSGroup/litewallet/litewallet/slim/types"
	"strings"
)

func CreateTransfer(cliCtx context.CLIContext, addrto, coinstr, privkey, chainid string) ([]byte, error) {
	return tx.BuildAndSignTx(cliCtx, privkey, chainid, func() (txs.ITx, error) {
		var key ed25519local.PrivKeyEd25519
		ts := "{\"type\": \"tendermint/PrivKeyEd25519\",\"value\": \"" + privkey + "\"}"
		err := ctxs.Cdc.UnmarshalJSON([]byte(ts), &key)
		if err != nil {
			fmt.Println(err)
		}
		addrben32, _ := bech32local.ConvertAndEncode(btypes.PREF_ADD, key.PubKey().Address().Bytes())

		sendersStr := addrben32 + `,` + coinstr
		senders, err := parseTransItem(cliCtx, sendersStr)
		if err != nil {
			return nil, err
		}

		receiversStr := addrto + `,` + coinstr
		receivers, err := parseTransItem(cliCtx, receiversStr)
		if err != nil {
			return nil, err
		}
		return bank_txs.TxTransfer{
			Senders:   senders,
			Receivers: receivers,
		}, nil
	})
}

// Parse flags from string
func parseTransItem(cliCtx context.CLIContext, str string) (types.TransItems, error) {
	items := make(types.TransItems, 0)
	tis := strings.Split(str, ";")
	for _, ti := range tis {
		if ti == "" {
			continue
		}

		addrAndCoins := strings.Split(ti, ",")
		if len(addrAndCoins) < 2 {
			return nil, fmt.Errorf("`%s` not match rules", ti)
		}

		addr, err := qcliacc.GetAddrFromValue(addrAndCoins[0])
		if err != nil {
			return nil, err
		}
		qos, qscs, err := qtypes.ParseCoins(strings.Join(addrAndCoins[1:], ","))
		if err != nil {
			return nil, err
		}
		items = append(items, types.TransItem{
			Address: addr,
			QOS:     qos,
			QSCs:    qscs,
		})
	}

	return items, nil
}

//// Parse QOS and QSCs from string
//// str example : 100qos,100qstar
//func NewParseCoins(str string) (btypes.BigInt, types.QSCs, error) {
//	if len(str) == 0 {
//		return btypes.ZeroInt(), btypes.QSCs{}, nil
//	}
//	reDnm := `[[:alpha:]][[:alnum:]]{2,15}`
//	reAmt := `[[:digit:]]+`
//	reSpc := `[[:space:]]*`
//	reCoin := regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
//
//	arr := strings.Split(str, ",")
//	qos := btypes.ZeroInt()
//	qscs := btypes.QSCs{}
//	for _, q := range arr {
//		coin := reCoin.FindStringSubmatch(q)
//		if len(coin) != 3 {
//			return btypes.ZeroInt(), nil, fmt.Errorf("coins str: %s parse faild", q)
//		}
//		coin[2] = strings.TrimSpace(coin[2])
//		amount, err := strconv.ParseInt(strings.TrimSpace(coin[1]), 10, 64)
//		if err != nil {
//			return btypes.ZeroInt(), nil, err
//		}
//		if strings.ToLower(coin[2]) == "qos" {
//			qos = btypes.NewInt(amount)
//		} else {
//			qscs = append(qscs, &types.QSC{
//				coin[2],
//				btypes.NewInt(amount),
//			})
//		}
//
//	}
//
//	return qos, qscs, nil
//}
