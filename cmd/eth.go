package cmd

import (
	"github.com/QOSGroup/litewallet/litewallet/eth"
	"github.com/spf13/cobra"
)

// ethCmd
var (
	rootDir, node, fromName, password, toAddr, gasPrice, amount string
	gasLimit                                                    int64
	ethCmd                                                      = &cobra.Command{
		Use:   "eth",
		Short: "eth cli command",
	}
	ethTransferETHCmd = &cobra.Command{
		Use:   "transfer",
		Short: "transfer cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.TransferETH(rootDir, node, fromName, password, toAddr, gasPrice, amount, gasLimit)
			cmd.Println(res)
		},
	}
)

func init() {
	ethTransferETHCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	ethTransferETHCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	ethTransferETHCmd.PersistentFlags().StringVar(&fromName, "fromName", "", "fromName")
	ethTransferETHCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	ethTransferETHCmd.PersistentFlags().StringVar(&toAddr, "toAddr", "", "toAddr")
	ethTransferETHCmd.PersistentFlags().StringVar(&gasPrice, "gasPrice", "", "gasPrice")
	ethTransferETHCmd.PersistentFlags().StringVar(&amount, "amount", "", "amount")
	ethTransferETHCmd.PersistentFlags().Int64Var(&gasLimit, "gasLimit", 0, "gasLimit")

	ethCmd.AddCommand(ethTransferETHCmd)
}
