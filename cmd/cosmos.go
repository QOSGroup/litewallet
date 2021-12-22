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
)

func init() {
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

	cosmosCmd.AddCommand(cosmosCreateSeedCmd)
	cosmosCmd.AddCommand(cosmosCreateAccountCmd)
	cosmosCmd.AddCommand(cosmosRecoverKeyCmd)
	cosmosCmd.AddCommand(cosmosUpdateKeyCmd)
	cosmosCmd.AddCommand(cosmosWalletAddressCheckCmd)
}
