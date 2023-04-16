package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available versions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("cli version list\n")
	},
}
