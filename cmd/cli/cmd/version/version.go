package version

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"github.com/ryurock/cli/config"
	"gopkg.in/yaml.v3"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "バージョンを操作するコマンド",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			versionRun(cmd, args)
			return nil
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
	githubResult := []struct {
		Name string
	}{}

	cliConfig := config.NewCliConfig()
	githubOraganization := cliConfig.GitHub.Repo.Organization
	githubRepoName := cliConfig.GitHub.Repo.Name
	err = client.Get(fmt.Sprintf("repos/%s/%s/tags", githubOraganization, githubRepoName), &githubResult)
	if err != nil {
		log.Fatal(err)
	}

	switch cmd.Parent().Flags().Lookup("format").Value.String() {
	case "json":
		response := map[string]interface{}{"name": githubResult[0].Name}

		jsonData, err := json.Marshal(response)
		if err != nil {
			message := color.RedString(err.Error())
			panic(message)
		}

		fmt.Printf("%s\n", jsonData)
	case "yaml":
		response := map[string]interface{}{"name": githubResult[0].Name}
		yamlData, err := yaml.Marshal(response)
		if err != nil {
			message := color.RedString(err.Error())
			panic(message)
		}
		fmt.Printf("%s", yamlData)

	case "text":
		color.Cyan(githubResult[0].Name)
	case "table":
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Name")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		tbl.AddRow(githubResult[0].Name)
		tbl.Print()
	}
}

func init() {
	VersionCmd.AddCommand(versionListCmd)
	VersionCmd.AddCommand(versionInstallCmd)
}
