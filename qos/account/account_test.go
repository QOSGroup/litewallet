package account

import (
	"github.com/stretchr/testify/assert"
	tDB "github.com/tendermint/tm-db"
	"github.com/tyler-smith/go-bip39"
	"os"
	"path/filepath"
	"testing"
)

func TestAccountOperation(t *testing.T) {

	dbPath, _ := filepath.Abs("../db")
	defer os.RemoveAll(dbPath)

	db := tDB.NewDB("wallet", tDB.GoLevelDBBackend, dbPath)

	am := NewAccountManager(db)

	password := "1234567890"
	newPassword := "1223423445"

	for i := 0; i < 200; i++ {
		mnemonic, _ := genMnemonic()
		acc, err := am.NewAccountFromMnemonic("", password, mnemonic)
		assert.Nil(t, err)
		assert.NotNil(t, acc)
		assert.NotNil(t, acc.PrivateKeySecret)
	}

	l := am.ListAccounts()
	assert.Equal(t, len(l), 200)
	assert.Equal(t, l[0].Id, uint64(1))
	assert.Equal(t, l[100].Id, uint64(101))

	for i := 1; i <= 200; i++ {
		acc, err := am.getAccountByID(uint64(i))
		assert.Nil(t, err)
		assert.NotNil(t, acc)

		acc, err = am.getAccountByAddr(acc.Address)
		assert.Nil(t, err)
		assert.NotNil(t, acc)

		acc, err = am.QueryAccount(acc.Address, "")
		assert.Nil(t, err)
		assert.NotNil(t, acc)

		_, err = am.DecryptPrivateKey(acc, password)
		assert.Nil(t, err)

		acc, err = am.ModifyAccountPassword(acc, password, newPassword)
		assert.Nil(t, err)
		assert.NotNil(t, acc)

		_, err = am.DecryptPrivateKey(acc, password)
		assert.NotNil(t, err)

		err = am.DeleteAccount(acc, password)
		assert.NotNil(t, err)

		acc, err = am.getAccountByAddr(acc.Address)
		assert.Nil(t, err)
		assert.NotNil(t, acc)

		err = am.DeleteAccount(acc, newPassword)
		assert.Nil(t, err)

		acc, err = am.getAccountByAddr(acc.Address)
		assert.NotNil(t, err)
		assert.Nil(t, acc)

	}

	l = am.ListAccounts()
	assert.Equal(t, len(l), 0)

	mnemonic, _ := genMnemonic()
	mnemonic2, _ := genMnemonic()
	acc, err := am.NewAccountFromMnemonic("test1", password, mnemonic)
	assert.Nil(t, err)
	assert.NotNil(t, acc)

	acc, err = am.NewAccountFromMnemonic("test1", password, mnemonic2)
	assert.NotNil(t, err)
	assert.Nil(t, acc)

	acc, err = am.NewAccountFromMnemonic("test2", password, mnemonic)
	assert.NotNil(t, err)
	assert.Nil(t, acc)
}

func genMnemonic() (string, error) {
	bz, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(bz)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
