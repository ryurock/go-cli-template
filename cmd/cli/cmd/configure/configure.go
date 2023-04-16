package configure

import (
	"errors"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ConfugureCmd = &cobra.Command{
	Use:   "configure",
	Short: "CLI の設定を操作するコマンド",
	RunE: func(cmd *cobra.Command, args []string) error {

		message := color.RedString("a valid subcommand is required")
		return errors.New(message)
	},
}

func init() {
	ConfugureCmd.AddCommand(configureInitCmd)
}
