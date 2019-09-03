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
	name := "cosmostest"
	password := "wm131421"
	seed := "very online issue brain swarm deer thunder sustain pact jelly lift return"
	output := CreateAccount(rootDir, name, password, seed)
	t.Log(output)
}

func TestRecoverKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "easyzonetest"
	password := "wm131421"
	seed := "very online issue brain swarm deer thunder sustain pact jelly lift return"
	output := RecoverKey(rootDir, name, password, seed)
	t.Log(output)
}

func TestUpdateKey(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "easyzonetest"
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
