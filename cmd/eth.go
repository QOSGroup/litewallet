package cmd

import (
	"github.com/QOSGroup/litewallet/litewallet/eth"
	"github.com/spf13/cobra"
)

// ethCmd
var (
	addr, name, mnemonic, tokenAddr, tokenValue                 string
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

	ethTransferERC20Cmd = &cobra.Command{
		Use:   "transferERC20",
		Short: "transferERC20 cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.TransferERC20(rootDir, node, fromName, password, toAddr, tokenAddr,
				tokenValue, gasPrice, gasLimit)
			cmd.Println(res)
		},
	}

	ethCreateAccountCmd = &cobra.Command{
		Use:   "createAccount",
		Short: "createAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.CreateAccount(rootDir, name, password, mnemonic)
			cmd.Println(res)
		},
	}

	ethGetAccountCmd = &cobra.Command{
		Use:   "getAccount",
		Short: "getAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.GetAccount(node, addr)
			cmd.Println(res)
		},
	}

	ethListLocalAccountCmd = &cobra.Command{
		Use:   "listLocalAccount",
		Short: "listLocalAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.ListLocalAccount(rootDir)
			cmd.Println(res)
		},
	}

	ethRecoverAccountCmd = &cobra.Command{
		Use:   "recoverAccount",
		Short: "recoverAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := eth.RecoverAccount(rootDir, name, password, mnemonic)
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

	ethTransferERC20Cmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&node, "node", "", "node")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&fromName, "fromName", "", "fromName")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&password, "password", "", "password")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&toAddr, "toAddr", "", "toAddr")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&tokenAddr, "tokenAddr", "", "tokenAddr")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&tokenValue, "tokenValue", "", "tokenValue")
	ethTransferERC20Cmd.PersistentFlags().StringVar(&gasPrice, "gasPrice", "", "gasPrice")
	ethTransferERC20Cmd.PersistentFlags().Int64Var(&gasLimit, "gasLimit", 0, "gasLimit")

	ethGetAccountCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	ethGetAccountCmd.PersistentFlags().StringVar(&addr, "addr", "", "addr")

	ethCreateAccountCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	ethCreateAccountCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	ethCreateAccountCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	ethCreateAccountCmd.PersistentFlags().StringVar(&mnemonic, "mnemonic", "", "mnemonic")

	ethListLocalAccountCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")

	ethRecoverAccountCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	ethRecoverAccountCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	ethRecoverAccountCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	ethRecoverAccountCmd.PersistentFlags().StringVar(&mnemonic, "mnemonic", "", "mnemonic")

	ethCmd.AddCommand(ethTransferETHCmd)
	ethCmd.AddCommand(ethTransferERC20Cmd)
	ethCmd.AddCommand(ethCreateAccountCmd)
	ethCmd.AddCommand(ethListLocalAccountCmd)
	ethCmd.AddCommand(ethGetAccountCmd)
	ethCmd.AddCommand(ethRecoverAccountCmd)
}
