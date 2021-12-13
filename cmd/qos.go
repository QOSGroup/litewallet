package cmd

import (
	"github.com/QOSGroup/litewallet/litewallet/slim"
	"github.com/spf13/cobra"
)

// ethCmd
var (
	mncode string
	qosCmd = &cobra.Command{
		Use:   "qos",
		Short: "qos cli command",
	}

	qosAccountCreateCmd = &cobra.Command{
		Use:   "accountCreate",
		Short: "accountCreate cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := slim.AccountCreate(password)
			cmd.Println(res)
		},
	}

	qosAccountRecoverCmd = &cobra.Command{
		Use:   "accountRecover",
		Short: "accountRecover cli command",
		Run: func(cmd *cobra.Command, args []string) {
			res := slim.AccountRecoverStr(mncode, password)
			cmd.Println(res)
		},
	}
)

func init() {
	qosAccountCreateCmd.PersistentFlags().StringVar(&password, "password", "", "password")

	qosAccountRecoverCmd.PersistentFlags().StringVar(&mncode, "mncode", "", "mncode")
	qosAccountRecoverCmd.PersistentFlags().StringVar(&password, "password", "", "password")

	qosCmd.AddCommand(qosAccountCreateCmd)
	qosCmd.AddCommand(qosAccountRecoverCmd)
}
