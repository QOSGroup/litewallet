package eth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"

	"github.com/cosmos/cosmos-sdk/crypto/keys/keyerror"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/crypto/bcrypt"
	tcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/armor"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tyler-smith/go-bip39"
	"log"
	"path/filepath"
	"strings"
)

const (
	blockTypePrivKey =	"TENDERMINT PRIVATE KEY"
	infoSuffix = "info"
	BcryptSecurityParameter = 12
)

// localInfo is the public information about a locally stored key
type LocalInfo struct {
	Name         string         `json:"name"`
	PubKey       string 		`json:"pubkey"`
	PrivKeyArmor string         `json:"privkey"`
	Address      string			`json:"address"`
}

type dbKeybase struct {
	db dbm.DB
}

type KeyOutput struct {
	Name      string                 `json:"name"`
	Type      string                 `json:"type"`
	Address   string                 `json:"address"`
	PubKey    string                 `json:"pubkey"`
	Mnemonic  string                 `json:"mnemonic,omitempty"`
	Denom  	  string 				 `json:"denom"`
}

//follow the cosmos hd implementation
func CreateAccount(rootDir, name, password, mnemonic string) string {
	if name == "" {
		err := errMissingName()
		return err.Error()
	}
	if password == "" {
		err := errMissingPassword()
		return err.Error()
	}
	//generate wallet with mnemonic
	if mnemonic == "" {
		return fmt.Sprintf("mnemonic is required")
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Sprintf("mnemonic is invalid")
	}
	//convert mnemonic string to seed byte
	seed := bip39.NewSeed(mnemonic, "")

	//dpath for the key base path derive:  m / purpose' / coin_type' / account' / change / address_index (m/44'/60'/0'/0/0)
	dpath, err := accounts.ParseDerivationPath(`m/44'/60'/0'/0/0`)
	if err != nil {
		return err.Error()
	}

	//fetch the masterKey for the wallet
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	//masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return err.Error()
	}

	key := masterKey

	for _, n := range dpath {
		key, err = key.Child(n)
		if err != nil {
			return err.Error()
		}
	}

	//generate the privateKey and pubKey
	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return err.Error()
	}
	//then the pubKey
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err.Error()
	}
	//the address, pubKey, PrivKey with hexString format
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	pubKeyHex := hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))[4:]

	//Armor and encrypt the privateKey
	privateKeyByte := crypto.FromECDSA(privateKeyECDSA)
	saltBytes, encBytes := encryptPrivKey(privateKeyByte, password)
	header := map[string]string{
		"kdf":  "bcrypt",
		"salt": fmt.Sprintf("%X", saltBytes),
	}
	priKeyAmor := armor.EncodeArmor(blockTypePrivKey, header, encBytes)
	//priKeyHex := hexutil.Encode(crypto.FromECDSA(privateKeyECDSA))[2:]

	//gather the local info into struct
	LInfo := &LocalInfo{
		Name:         name,
		PubKey:       pubKeyHex,
		PrivKeyArmor: priKeyAmor,
		Address: 	  address,
	}

	//write the local info by key
	key1 := []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
	serializeInfo, err := json.Marshal(LInfo)
	if err != nil {
		return err.Error()
	}

	//init a go level DB to store the key and Info, specify the ethkeys
	db, err := dbm.NewGoLevelDB("keys", filepath.Join(rootDir, "ethkeys"))
	if err != nil {
		return err.Error()
	}
	//Do not check whether the name conflict with existing names
	// List returns the keys from storage in alphabetical order.
	//var res []LocalInfo
	//iter := db.Iterator(nil, nil)
	//defer iter.Close()
	//for ; iter.Valid(); iter.Next() {
	//	key := string(iter.Key())
	//	// need to include only keys in storage that have an info suffix
	//	if strings.HasSuffix(key, infoSuffix) {
	//		info, err := readInfo(iter.Value())
	//		if err != nil {
	//			return err.Error()
	//		}
	//		res = append(res, info)
	//	}
	//}
	//// check if already exists
	//for _, info := range res {
	//	if info.Name == name {
	//		err = errKeyNameConflict(name)
	//		return err.Error()
	//	}
	//}

	db.SetSync(key1,serializeInfo)
	// store a pointer to the infokey by address for fast lookup
	addrKey := []byte(fmt.Sprintf("%s.%s", address, "addr"))
	db.SetSync(addrKey, key1)

	//fetch the result
	Ko := KeyOutput{LInfo.Name, "local", LInfo.Address,LInfo.PubKey,mnemonic,"ETH"}
	respbyte, _ := json.Marshal(Ko)
	return string(respbyte)

}

//List local account
func ListLocalAccount(rootDir string) string {
	//init a go level DB to store the key and Info, specify the ethkeys
	db, err := dbm.NewGoLevelDB("keys", filepath.Join(rootDir, "ethkeys"))
	if err != nil {
		return err.Error()
	}
	// List returns the keys from storage in alphabetical order.
	var res []LocalInfo
	iter := db.Iterator(nil, nil)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		key := string(iter.Key())
		// need to include only keys in storage that have an info suffix
		if strings.HasSuffix(key, infoSuffix) {
			info, err := readInfo(iter.Value())
			if err != nil {
				return err.Error()
			}
			res = append(res, info)
		}
	}
	// check if already exists
	var KoG []KeyOutput
	for _, info := range res {
		Ko := KeyOutput{info.Name, "local", info.Address,info.PubKey,"","ETH"}
		KoG = append(KoG, Ko)
	}
	Kos,_ := json.Marshal(KoG)
	return string(Kos)
}


// decoding info
func readInfo(bz []byte) (info LocalInfo, err error) {
	err = json.Unmarshal(bz, &info)
	return
}


// encrypt the given privKey with the passphrase using a randomly
// generated salt and the xsalsa20 cipher. returns the salt and the
// encrypted priv key.
func encryptPrivKey(privKey []byte, passphrase string) (saltBytes []byte, encBytes []byte) {
	saltBytes = tcrypto.CRandBytes(16)
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		cmn.Exit("Error generating bcrypt key from passphrase: " + err.Error())
	}
	key = tcrypto.Sha256(key) // get 32 bytes
	privKeyBytes := privKey
	return saltBytes, xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
}


// Unarmor and decrypt the private key.
func UnarmorDecryptPrivKey(armorStr string, passphrase string) (*ecdsa.PrivateKey, error) {
	var privKey *ecdsa.PrivateKey
	blockType, header, encBytes, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return privKey, err
	}
	if blockType != blockTypePrivKey {
		return privKey, fmt.Errorf("Unrecognized armor type: %v", blockType)
	}
	if header["kdf"] != "bcrypt" {
		return privKey, fmt.Errorf("Unrecognized KDF type: %v", header["KDF"])
	}
	if header["salt"] == "" {
		return privKey, fmt.Errorf("Missing salt bytes")
	}
	saltBytes, err := hex.DecodeString(header["salt"])
	if err != nil {
		return privKey, fmt.Errorf("Error decoding salt: %v", err.Error())
	}
	privKey, err = decryptPrivKey(saltBytes, encBytes, passphrase)
	return privKey, err
}

func decryptPrivKey(saltBytes []byte, encBytes []byte, passphrase string) (privKey *ecdsa.PrivateKey, err error) {
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		cmn.Exit("Error generating bcrypt key from passphrase: " + err.Error())
	}
	key = tcrypto.Sha256(key) // Get 32 bytes
	privKeyBytes, err := xsalsa20symmetric.DecryptSymmetric(encBytes, key)
	if err != nil && err.Error() == "Ciphertext decryption failed" {
		return privKey, keyerror.NewErrWrongPassword()
	} else if err != nil {
		return privKey, err
	}
	privKey, err = crypto.ToECDSA(privKeyBytes)
	return privKey, err
}

//Fetch private key for signning
func FetchtoSign(rootDir, name, password string) (privKey *ecdsa.PrivateKey, err error) {
	//init a go level DB to store the key and Info, specify the ethkeys
	db, err := dbm.NewGoLevelDB("keys", filepath.Join(rootDir, "ethkeys"))
	if err != nil {
		return nil, err
	}
	bs := db.Get(infoKey(name))
	if len(bs) == 0 {
		return nil, keyerror.NewErrKeyNotFound(name)
	}
	//get the LocalInfo by key
	Li, err := readInfo(bs)
	if err != nil {
		return nil, err
	}
	armorStr := Li.PrivKeyArmor
	//fmt.Println("PrivateKeyArmor",armorStr)
	return UnarmorDecryptPrivKey(armorStr,password)
}

func infoKey(name string) []byte {
	return []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
}

func errMissingName() error {
	return fmt.Errorf("you have to specify a name for the locally stored account")
}

func errMissingPassword() error {
	return fmt.Errorf("you have to specify a password for the locally stored account")
}

func errKeyNameConflict(name string) error {
	return fmt.Errorf("acount with name %s already exists", name)
}

//verify the signature
func SigVerify(rootDir,name,password string, data2sign []byte) {
	//Fetch the privateKey
	privateKey, err := FetchtoSign(rootDir,name,password)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	//data to be signed
	//data := []byte("hello")
	hash := crypto.Keccak256Hash(data2sign)

	//generate the signature
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//recover the pubkey from signature
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println("sigPublicKey vs publicKeyBytes", matches) // true


	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println("sigPublicKeyBytes vs publicKeyBytes", matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println("no recover verify", verified) // true

}
