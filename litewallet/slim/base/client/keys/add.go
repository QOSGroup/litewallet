package keys
//
//import (
//	"fmt"
//	//"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
//	"github.com/QOSGroup/litewallet/litewallet/slim/txs"
//	"github.com/cosmos/cosmos-sdk/crypto/keys"
//	"github.com/spf13/viper"
//	"github.com/tendermint/tendermint/libs/cli"
//)
//
//const (
//	flagRecover = "recover"
//)
//
//func AddKey(name, pass string) {
//	//cliCtx := context.NewCLIContext("").WithCodec(txs.Cdc)
//	runAddCmd(name, pass)
//}
//
//func runAddCmd(name, pass string) (string, string, error) {
//	var kb keys.Keybase
//	var err error
//
//	kb, err = GetKeyBase(ctx)
//	if err != nil {
//		return "", "", err
//	}
//
//	_, err = kb.Get(name)
//	if err == nil {
//		if response, err := utils.GetConfirmation(
//			fmt.Sprintf("override the existing name %s", name), buf); err != nil || !response {
//			return "", "", err
//		}
//	}
//
//	//pass, err = utils.GetCheckPassword(
//	//	"Enter a passphrase for your key:",
//	//	"Repeat the passphrase:", buf)
//	//if err != nil {
//	//	return err
//	//}
//
//	if viper.GetBool(flagRecover) {
//		//seed, err := utils.GetSeed(
//		//	"Enter your recovery seed phrase:", buf)
//		//if err != nil {
//		//	return "", "", err
//		//}
//		//info, err := kb.Derive(name, seed, pass, hd.FullFundraiserPath)
//		//if err != nil {
//		//	return "", "", err
//		//}
//		//printCreate(ctx, info, "")
//	}
//	info, seed, err := kb.CreateEnMnemonic(name, pass)
//	if err != nil {
//		return "", "", err
//	}
//	return info, seed, nil
//}
//
//func printCreate(ctx context.CLIContext, info keys.Info, seed string) {
//	output := viper.Get(cli.OutputFlag)
//	switch output {
//	case "json":
//		out, err := Bech32KeyOutput(ctx, info)
//		if err != nil {
//			panic(err)
//		}
//		out.Seed = seed
//		var jsonString []byte
//		jsonString, err = ctx.Codec.MarshalJSONIndent(out, "", "  ")
//
//		if err != nil {
//			panic(err) // really shouldn't happen...
//		}
//		fmt.Println(string(jsonString))
//	default:
//		printKeyInfo(ctx, info, Bech32KeyOutput)
//
//		fmt.Println("**Important** write this seed phrase in a safe place.")
//		fmt.Println("It is the only way to recover your account if you ever forget your password.")
//		fmt.Println()
//		fmt.Println(seed)
//	}
//}
