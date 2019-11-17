package eth

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

func RecoverAccount(rootDir, name, password, mnemonic string) string {
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
	output := CreateAccount(rootDir, name, password, mnemonic)
	return output
}
