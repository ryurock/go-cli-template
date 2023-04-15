package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	Long:  `All software has versions. This is cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
