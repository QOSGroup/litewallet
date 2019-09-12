package context

import (
	"github.com/QOSGroup/litewallet/litewallet/slim/tendermint/rpc/client"
)

// CLIContext implements a typical CLI context created in SDK modules for
// transaction handling and queries.
type CLIContext struct {
	Client  *client.HTTP
	NodeURI string
}

// NewCLIContext returns a new initialized CLIContext with parameters from the
// command line using Viper.
func NewCLIContext(remote string) CLIContext {
	rpc := client.NewHTTP(remote, "/websocket")

	return CLIContext{
		Client:  rpc,
		NodeURI: remote,
	}
}
