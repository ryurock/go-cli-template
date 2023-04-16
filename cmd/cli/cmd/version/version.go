package version

import (
	"errors"
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			versionRun(cmd, args)
			return nil
		}
		// subCommands := []string{"install", "list"}
		// result := slices.Contains(subCommands, args[0])
		// fmt.Println(result) // true

		fmt.Println(args)
		err := cmd.Help()
		if err != nil {
			return err
		}
		message := color.RedString("a valid subcommand is required")
		return errors.New(message)
	},
}

func versionRun(cmd *cobra.Command, args []string) {
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}
	response := []struct {
		Name string
	}{}
	err = client.Get("repos/ryurock/go-cli-template/tags", &response)
	if err != nil {
		log.Fatal(err)
	}

	color.Cyan(response[0].Name)
}

func init() {
	VersionCmd.AddCommand(versionListCmd)
}
