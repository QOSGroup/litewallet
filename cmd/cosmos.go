package cmd

import (
	"github.com/QOSGroup/litewallet/litewallet/sdksource"
	"github.com/spf13/cobra"
)

// cosmosCmd
var (
	seed      string
	cosmosCmd = &cobra.Command{
		Use:   "cosmos",
		Short: "cosmos cli command",
	}

	cosmosGetAccountCmd = &cobra.Command{
		Use:   "getAccount",
		Short: "getAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.GetAccount(rootDir, node, chainID, addr)
			cmd.Println(res)
		},
	}

	cosmosCreateSeedCmd = &cobra.Command{
		Use:   "createSeed",
		Short: "createSeed cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.CreateSeed()
			cmd.Println(res)
		},
	}

	cosmosCreateAccountCmd = &cobra.Command{
		Use:   "createAccount",
		Short: "createAccount cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.CreateAccount(rootDir, name, password, seed)
			cmd.Println(res)
		},
	}

	cosmosRecoverKeyCmd = &cobra.Command{
		Use:   "recoverKey",
		Short: "recoverKey cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.RecoverKey(rootDir, name, password, seed)
			cmd.Println(res)
		},
	}

	cosmosUpdateKeyCmd = &cobra.Command{
		Use:   "updateKey",
		Short: "updateKey cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.UpdateKey(rootDir, name, oldpass, newpass)
			cmd.Println(res)
		},
	}

	cosmosWalletAddressCheckCmd = &cobra.Command{
		Use:   "walletAddressCheck",
		Short: "walletAddressCheck cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.WalletAddressCheck(addr)
			cmd.Println(res)
		},
	}

	cosmosTransferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "transfer cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.Transfer(rootDir, node, chainID, fromName, password, toStr, coinStr, feeStr, broadcastMode)
			cmd.Println(res)
		},
	}

	cosmosDelegateCmd = &cobra.Command{
		Use:   "delegate",
		Short: "delegate cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.Delegate(rootDir, node, chainID, delegatorName, password, delegatorAddr, validatorAddr,
				delegationCoinStr, feeStr, broadcastMode)
			cmd.Println(res)
		},
	}

	cosmosGetDelegationSharesCmd = &cobra.Command{
		Use:   "getDelegationShares",
		Short: "getDelegationShares cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := sdksource.GetDelegationShares(rootDir, node, chainID, delegatorAddr, validatorAddr)
			cmd.Println(res)
		},
	}
)

func init() {
	cosmosGetAccountCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosGetAccountCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	cosmosGetAccountCmd.PersistentFlags().StringVar(&chainID, "chainID", "", "chainID")
	cosmosGetAccountCmd.PersistentFlags().StringVar(&addr, "addr", "", "addr")

	cosmosCreateAccountCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosCreateAccountCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	cosmosCreateAccountCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	cosmosCreateAccountCmd.PersistentFlags().StringVar(&seed, "seed", "", "seed")

	cosmosRecoverKeyCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosRecoverKeyCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	cosmosRecoverKeyCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	cosmosRecoverKeyCmd.PersistentFlags().StringVar(&seed, "seed", "", "seed")

	cosmosUpdateKeyCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosUpdateKeyCmd.PersistentFlags().StringVar(&name, "name", "", "name")
	cosmosUpdateKeyCmd.PersistentFlags().StringVar(&password, "oldpass", "", "oldpass")
	cosmosUpdateKeyCmd.PersistentFlags().StringVar(&seed, "newpass", "", "newpass")

	cosmosWalletAddressCheckCmd.PersistentFlags().StringVar(&seed, "addr", "", "addr")

	cosmosTransferCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosTransferCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	cosmosTransferCmd.PersistentFlags().StringVar(&chainID, "chainID", "", "chainID")
	cosmosTransferCmd.PersistentFlags().StringVar(&fromName, "fromName", "", "fromName")
	cosmosTransferCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	cosmosTransferCmd.PersistentFlags().StringVar(&toStr, "toStr", "", "toStr")
	cosmosTransferCmd.PersistentFlags().StringVar(&coinStr, "coinStr", "", "coinStr")
	cosmosTransferCmd.PersistentFlags().StringVar(&feeStr, "feeStr", "", "feeStr")
	cosmosTransferCmd.PersistentFlags().StringVar(&broadcastMode, "broadcastMode", "", "broadcastMode")

	cosmosDelegateCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosDelegateCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	cosmosDelegateCmd.PersistentFlags().StringVar(&chainID, "chainID", "", "chainID")
	cosmosDelegateCmd.PersistentFlags().StringVar(&delegatorName, "delegatorName", "", "delegatorName")
	cosmosDelegateCmd.PersistentFlags().StringVar(&password, "password", "", "password")
	cosmosDelegateCmd.PersistentFlags().StringVar(&delegatorAddr, "delegatorAddr", "", "delegatorAddr")
	cosmosDelegateCmd.PersistentFlags().StringVar(&validatorAddr, "validatorAddr", "", "validatorAddr")
	cosmosDelegateCmd.PersistentFlags().StringVar(&delegationCoinStr, "delegationCoinStr", "", "delegationCoinStr")
	cosmosDelegateCmd.PersistentFlags().StringVar(&feeStr, "feeStr", "", "feeStr")
	cosmosDelegateCmd.PersistentFlags().StringVar(&broadcastMode, "broadcastMode", "", "broadcastMode")

	cosmosGetDelegationSharesCmd.PersistentFlags().StringVar(&rootDir, "rootDir", "", "rootDir")
	cosmosGetDelegationSharesCmd.PersistentFlags().StringVar(&node, "node", "", "node")
	cosmosGetDelegationSharesCmd.PersistentFlags().StringVar(&chainID, "chainID", "", "chainID")
	cosmosGetDelegationSharesCmd.PersistentFlags().StringVar(&delegatorAddr, "delegatorAddr", "", "delegatorAddr")
	cosmosGetDelegationSharesCmd.PersistentFlags().StringVar(&validatorAddr, "validatorAddr", "", "validatorAddr")

	cosmosCmd.AddCommand(cosmosGetAccountCmd)
	cosmosCmd.AddCommand(cosmosCreateSeedCmd)
	cosmosCmd.AddCommand(cosmosCreateAccountCmd)
	cosmosCmd.AddCommand(cosmosRecoverKeyCmd)
	cosmosCmd.AddCommand(cosmosUpdateKeyCmd)
	cosmosCmd.AddCommand(cosmosWalletAddressCheckCmd)
	cosmosCmd.AddCommand(cosmosTransferCmd)
	cosmosCmd.AddCommand(cosmosDelegateCmd)
	cosmosCmd.AddCommand(cosmosGetDelegationSharesCmd)
}
