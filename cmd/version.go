package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0"

// versionCmd  版本信息
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
