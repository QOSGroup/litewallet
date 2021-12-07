package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = `
 _      _ _    __          __   _ _      _   
| |    (_) |   \ \        / /  | | |    | |  
| |     _| |_ __\ \  /\  / /_ _| | | ___| |_ 
| |    | | __/ _ \ \/  \/ / _ | | |/ _ \ __|
| |____| | ||  __/\  /\  / (_| | | |  __/ |_
|______|_|\__\___| \/  \/ \__,_|_|_|\___|\__|

`

// rootCmd 主命令
var rootCmd = &cobra.Command{
	Use:   "litewallet",
	Short: "litewallet 命令行工具",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		println(banner)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
