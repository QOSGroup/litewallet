package eth

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
	"os/user"
	"testing"
)

func TestAddress(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a8
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	output := hexutil.Encode(hash.Sum(nil)[12:])
	t.Log(output)
}

func TestCreateAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "cm1"
	password := "wm131421"
	seed := "fence shell flame stove zebra occur hurry steel drip gather debate tuition crumble cigar hood swarm unaware plunge lake artist snack skate between police"
	output := CreateAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestListLocalAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	output := ListLocalAccount(rootDir)
	t.Log(output)
}

func TestRecoverAccount(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "eth5"
	password := "wm131421"
	seed := "monster soap pipe grief tourist marine turkey scatter because fade actual robust"
	output := RecoverAccount(rootDir,name,password,seed)
	t.Log(output)
}

func TestFetchtoSign(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "cm1"
	password := "wm131421"
	output, err := FetchtoSign(rootDir,name,password)
	if err != nil {
		return
	}
	t.Log(output)
}


func TestSigVerify(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	name := "cm1"
	password := "wm131421"
	data2sign := []byte("hello")
	SigVerify(rootDir,name,password,data2sign)
}