package account

import (
	"errors"
	"fmt"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/account"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/client/context"
	"github.com/QOSGroup/litewallet/litewallet/slim/base/types"
	"strings"

	"github.com/spf13/viper"
)

var (
	ErrAccountNotExsits = errors.New("account not exists")
)

func QueryAccount(cliCtx context.CLIContext, addrStr string) (account.Account, error) {
	var addr types.Address
	addr, err := GetAddrFromValue(addrStr)
	if err != nil {
		return nil, err
	}

	return queryAccount(cliCtx, addr.Bytes())
}

func queryAccount(cliCtx context.CLIContext, addr []byte) (account.Account, error) {
	path := account.BuildAccountStoreQueryPath()
	res, err := cliCtx.Query(string(path), account.AddressStoreKey(types.AccAddress(addr)))
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, ErrAccountNotExsits
	}

	var acc account.Account
	err = cliCtx.Codec.UnmarshalBinaryBare(res, &acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func GetAccount(ctx context.CLIContext, address []byte) (account.Account, error) {
	return queryAccount(ctx, address)
}

func GetAccountFromBech32Addr(ctx context.CLIContext, bech32Addr string) (account.Account, error) {

	addrBytes, err := types.AccAddressFromBech32(bech32Addr)

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid bech32Addr", bech32Addr)
	}

	return queryAccount(ctx, addrBytes)
}

func GetAccountNonce(ctx context.CLIContext, address []byte) (int64, error) {
	account, err := queryAccount(ctx, address)

	if err == ErrAccountNotExsits {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return account.GetNonce(), nil
}

func IsAccountExists(ctx context.CLIContext, address []byte) bool {
	_, err := queryAccount(ctx, address)

	if err != nil {
		return false
	}

	return true
}

func GetAddrFromFlag(ctx context.CLIContext, flag string) (types.AccAddress, error) {
	value := viper.GetString(flag)
	return GetAddrFromValue(value)
}

func GetAddrFromValue(value string) (types.AccAddress, error) {
	//prefix := types.GetAddressConfig().GetBech32AccountAddrPrefix()
	//if strings.HasPrefix(value, prefix) {
	//	addr, err := types.AccAddressFromBech32(value)
	//	if err == nil {
	//		return addr, nil
	//	} else {
	//		return types.AccAddress{}, fmt.Errorf("Address:%s is not a valid bech32 address. Error: %s", value, err.Error())
	//	}
	//}
	//
	//info, err := keys.GetKeyInfo(ctx, value)
	//if err != nil {
	//	return nil, fmt.Errorf("Name:%s not exsits in current keybase. Error: %s", value, err.Error())
	//}
	//
	//return info.GetAddress(), nil
	addr, err := types.AccAddressFromBech32(value)
	if err == nil {
		return addr, nil
	} else {
		return types.AccAddress{}, fmt.Errorf("Address:%s is not a valid bech32 address. Error: %s", value, err.Error())
	}
}

func GetValidatorAddrFromFlag(ctx context.CLIContext, flag string) (types.ValAddress, error) {
	value := viper.GetString(flag)
	return GetValidatorAddrFromValue(value)
}

func GetValidatorAddrFromValue(value string) (types.ValAddress, error) {
	prefix := types.GetAddressConfig().GetBech32ValidatorAddrPrefix()

	if strings.HasPrefix(value, prefix) {
		addr, err := types.ValAddressFromBech32(value)
		if err == nil {
			return addr, nil
		}
	}

	return types.ValAddress{}, fmt.Errorf("%s is not a validator address. it must start with %s", value, prefix)
}
