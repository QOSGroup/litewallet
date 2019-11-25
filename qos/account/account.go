package account

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/QOSGroup/js-keys/keys"
	"github.com/QOSGroup/litewallet/types"
	db "github.com/tendermint/tm-db"
	"strings"
	"sync"
)

type Account struct {
	Id               uint64 `json:"id"`
	Name             string `json:"name"`
	PubKey           string `json:"public_key"`
	Address          string `json:"address"`
	PrivateKeySecret string `json:"private_key_enc"` //base64
}

type AccountWithMnemonic struct {
	*Account
	Mnemonic string `json:"mnemonic"`
}

func NewAccountWithMnemonic(account *Account, mnemonic string) *AccountWithMnemonic {
	return &AccountWithMnemonic{Account: account, Mnemonic: mnemonic}
}

type AccountPrivacy struct {
	*Account
	PrivateKey string `json:"private_key"` //hex
}

func NewAccountPrivacy(account *Account, privateKey string) *AccountPrivacy {
	return &AccountPrivacy{Account: account, PrivateKey: privateKey}
}

type AccountManager struct {
	db   db.DB
	lock sync.Mutex
}

func NewAccountManager(db db.DB) *AccountManager {
	return &AccountManager{
		db: db,
	}
}

func (manager *AccountManager) NewAccountFromMnemonic(name, password, mnemonic string) (*AccountWithMnemonic, error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	id := manager.increaseAccountIndex()
	if len(name) == 0 {
		name = fmt.Sprintf("Account%d", id)
	}

	oldAcc, _ := manager.getAccountByName(name)
	if oldAcc != nil {
		return nil, types.NewErrKeyAlreadyExists(name)
	}

	priKeyBz, pubKeyBz, err := keys.DeriveQOSKey(mnemonic)
	if err != nil {
		return nil, err
	}

	pubKey, _ := keys.Bech32ifyQOSAccPubKey(pubKeyBz)
	addr, _ := keys.Bech32ifyQOSAccAddressFromPubKey(pubKeyBz)
	oldAcc, _ = manager.getAccountByAddr(addr)
	if oldAcc != nil {
		return nil, types.NewErrKeyAlreadyExists(addr)
	}

	encBytes, saltBytes := types.Encrypt(priKeyBz, password)
	acc := &Account{
		Id:               id,
		Name:             name,
		PubKey:           pubKey,
		Address:          addr,
		PrivateKeySecret: types.EncodeBase64(encBytes),
	}

	manager.saveAccount(acc, saltBytes)

	return NewAccountWithMnemonic(acc, mnemonic), nil
}

func (manager *AccountManager) ImportAccount(priKeyBz []byte, password string) (*Account, error) {

	manager.lock.Lock()
	defer manager.lock.Unlock()

	id := manager.increaseAccountIndex()
	name := fmt.Sprintf("Account%d", id)
	oldAcc, _ := manager.getAccountByName(name)
	if oldAcc != nil {
		return nil, types.NewErrKeyAlreadyExists(name)
	}

	pubKeyBz := priKeyBz[32:]
	addr, err := keys.Bech32ifyQOSAccAddressFromPubKey(pubKeyBz)
	if err != nil {
		return nil, err
	}

	oldAcc, _ = manager.getAccountByAddr(addr)
	if oldAcc != nil {
		return nil, types.NewErrKeyAlreadyExists(addr)
	}

	pubKey, err := keys.Bech32ifyQOSAccPubKey(pubKeyBz)
	if err != nil {
		return nil, err
	}

	encBytes, saltBytes := types.Encrypt(priKeyBz, password)
	acc := &Account{
		Id:               id,
		Name:             name,
		PubKey:           pubKey,
		Address:          addr,
		PrivateKeySecret: types.EncodeBase64(encBytes),
	}

	manager.saveAccount(acc, saltBytes)
	return acc, nil
}

func (manager *AccountManager) QueryAccount(addr, name string) (*Account, error) {
	if len(addr) != 0 {
		return manager.getAccountByAddr(addr)
	} else if len(name) != 0 {
		return manager.getAccountByName(name)
	}

	return nil, types.NewErrKeyNotFound("")
}

func (manager *AccountManager) ListAccounts() []*Account {

	accounts := make([]*Account, 0)

	iterator := manager.db.Iterator(accountPrefixKey, types.PrefixEndBytes(accountPrefixKey))
	defer iterator.Close()

	for {
		if !iterator.Valid() {
			break
		}

		v := iterator.Value()
		var acc Account
		json.Unmarshal(v, &acc)
		accounts = append(accounts, &acc)

		iterator.Next()
	}

	return accounts
}

func (manager *AccountManager) DeleteAccount(acc *Account, password string) error {

	saltBytes, err := manager.getAccountSalt(acc)
	if err != nil {
		return err
	}

	if _, err := manager.verifyPassword(acc, saltBytes, password); err != nil {
		return err
	}

	manager.deleteAccount(acc)
	return nil
}

func (manager *AccountManager) ModifyAccountPassword(acc *Account, oldPassword, newPassword string) (*Account, error) {

	saltBytes, err := manager.getAccountSalt(acc)
	if err != nil {
		return nil, err
	}

	priKeyBz, err := manager.verifyPassword(acc, saltBytes, oldPassword)
	if err != nil {
		return nil, err
	}

	encBytes, saltBytes := types.Encrypt(priKeyBz, newPassword)
	acc.PrivateKeySecret = types.EncodeBase64(encBytes)

	manager.saveAccount(acc, saltBytes)
	return acc, nil
}

func (manager *AccountManager) DecryptPrivateKey(acc *Account, passphrase string) (string, error) {

	saltBytes, err := manager.getAccountSalt(acc)
	if err != nil {
		return "", err
	}

	bz, err := manager.verifyPassword(acc, saltBytes, passphrase)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(types.EncodeHex(bz)), nil
}

func (manager *AccountManager) increaseAccountIndex() (index uint64) {
	currentIndex := manager.getCurrentAccountIndex()
	currentIndex = currentIndex + 1
	manager.setCurrentAccountIndex(currentIndex)
	return currentIndex
}

func (manager *AccountManager) getCurrentAccountIndex() (index uint64) {
	bz := manager.db.Get(AccountIndexKey())
	if len(bz) == 0 {
		return
	}
	return types.BytesToUint64(bz)
}

func (manager *AccountManager) setCurrentAccountIndex(index uint64) {
	manager.db.Set(AccountIndexKey(), types.Uint64ToBytes(index))
}

func (manager *AccountManager) saveAccount(acc *Account, saltBytes []byte) {
	bz, _ := json.Marshal(acc)
	idBytes := types.Uint64ToBytes(acc.Id)

	batch := manager.db.NewBatch()
	batch.Set(AccountKey(acc.Id), bz)
	batch.Set(AccountAddressKey(acc.Address), idBytes)
	batch.Set(AccountNameKey(acc.Name), idBytes)
	batch.Set(AccountSaltKey(acc.Id), saltBytes)
	batch.Write()
}

func (manager *AccountManager) deleteAccount(acc *Account) {
	manager.db.Delete(AccountKey(acc.Id))
	manager.db.Delete(AccountAddressKey(acc.Address))
	manager.db.Delete(AccountNameKey(acc.Name))
	manager.db.Delete(AccountSaltKey(acc.Id))
}

func (manager *AccountManager) getAccountByID(id uint64) (acc *Account, err error) {
	bz := manager.db.Get(AccountKey(id))
	if len(bz) == 0 {
		return nil, types.NewErrKeyNotFound(fmt.Sprintf("%d", id))
	}

	acc = &Account{}
	err = json.Unmarshal(bz, acc)
	return
}

func (manager *AccountManager) getAccountByName(name string) (acc *Account, err error) {
	bz := manager.db.Get(AccountNameKey(name))
	if len(bz) == 0 {
		return nil, types.NewErrKeyNotFound(name)
	}
	return manager.getAccountByID(binary.BigEndian.Uint64(bz))
}

func (manager *AccountManager) getAccountByAddr(addr string) (acc *Account, err error) {
	bz := manager.db.Get(AccountAddressKey(addr))
	if len(bz) == 0 {
		return nil, types.NewErrKeyNotFound(addr)
	}
	return manager.getAccountByID(binary.BigEndian.Uint64(bz))
}

func (manager *AccountManager) getAccountSalt(acc *Account) ([]byte, error) {
	bz := manager.db.Get(AccountSaltKey(acc.Id))
	if len(bz) == 0 {
		return nil, types.NewErrKeyNotFound(fmt.Sprintf("%d", acc.Id))
	}
	return bz, nil
}

func (manager *AccountManager) verifyPassword(acc *Account, saltBytes []byte, passphrase string) ([]byte, error) {
	return types.Decrypt(types.MustDecodeBase64(acc.PrivateKeySecret), saltBytes, passphrase)
}
