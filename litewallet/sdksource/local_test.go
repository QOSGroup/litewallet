package sdksource

import (
	"os/user"
	"testing"
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
	name := "cosmos"
	password := "wm131421"
	seed := "chair green bag foster frog sock buzz giant hover party welcome ill"
	output := CreateAccount(rootDir, name, password, seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "easyzone2"
	password := "wm131421"
	seed := "style library milk jazz race dune disorder stay duck bunker garden favorite"
	output := RecoverKey(rootDir, name, password, seed)
	t.Log(output)
}

func TestUpdateKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "easyzone2"
	oldpass := "wm131421"
	newpass := "wm131422"
	output := UpdateKey(rootDir, name, oldpass, newpass)
	t.Log(output)
}

// func TestToken2Power(t *testing.T) {
// 	tokenInt := sdk.NewInt(int64(1000000))
// 	power := sdk.TokensToTendermintPower(tokenInt)
// 	t.Log(power)
// }

func TestWalletAddressCheck(t *testing.T) {
	address := "cosmos1vk4ark02kc7ac9ctgegml66496a8nyz0wyfn33"
	output := WalletAddressCheck(address)
	t.Log(output)
}
