package sdksource

import (
	"os/user"
	"testing"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCreateSeed(t *testing.T) {
	//usr, _ := user.Current()
	//rootDir := usr.HomeDir
	output := CreateSeed()
	t.Log(output)
}

func TestCreateAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "cm"
	password := "wm131421"
	seed := "tomorrow room limit true galaxy dove chicken fine resemble tonight record yellow"
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "easyzone2"
	password := "wm131421"
	seed := "style library milk jazz race dune disorder stay duck bunker garden favorite"
	output := RecoverKey(rootDir,name,password,seed)
	t.Log(output)
}

func TestUpdateKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "c33"
	oldpass := "wm131421"
	newpass := "wm131422"
	output := UpdateKey(rootDir, name, oldpass, newpass)
	t.Log(output)
}

func TestToken2Power(t *testing.T) {
	tokenInt := sdk.NewInt(int64(1000000))
	power := sdk.TokensToTendermintPower(tokenInt)
	t.Log(power)
}

func TestWalletAddressCheck(t *testing.T) {
	address := "0x1uyh63ddjrv944prku8sfn8vmmxluktl46dmy2e"
	output := WalletAddressCheck(address)
	t.Log(output)
}