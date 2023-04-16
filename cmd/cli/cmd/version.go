package cmd

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	Long:  `All software has versions. This is cli's`,
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Println(response)

		fmt.Println("version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
