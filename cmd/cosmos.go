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

	cosmosCmd.AddCommand(cosmosCreateSeedCmd)
	cosmosCmd.AddCommand(cosmosCreateAccountCmd)
	cosmosCmd.AddCommand(cosmosRecoverKeyCmd)
}
