package types

import (
	"regexp"
)

const (
	ETH    = "ETH"
	COSMOS = "COSMOS"
	QOS    = "QOS"
	OTHERS = "OTHERS"
)

var (
	cosmosAddressPattern = regexp.MustCompile("^cosmos1[0-9a-z]{38}$")
	qosAddressPattern    = regexp.MustCompile("^qosacc1[0-9a-z]{38}$")
	ethAddressPattern    = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
)

func AddressType(addr string) string {
	switch {
	case ethAddressPattern.MatchString(addr):
		return ETH

	case cosmosAddressPattern.MatchString(addr):
		return COSMOS

	case qosAddressPattern.MatchString(addr):
		return QOS
	default:
		return OTHERS
	}
}
