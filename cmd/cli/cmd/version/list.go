package version

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/ryurock/cli/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var versionListCmd = &cobra.Command{
	Use:   "list",
	Short: "バージョンの一覧を表示します",
	Run: func(cmd *cobra.Command, args []string) {
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
			message := color.RedString(err.Error())
			panic(message)
		}

		switch cmd.Parent().Flags().Lookup("format").Value.String() {
		case "json":
			response := []string{}
			for _, tag := range githubResult {
				response = append(response, tag.Name)
			}
			jsonData, err := json.Marshal(response)
			if err != nil {
				message := color.RedString(err.Error())
				panic(message)
			}

			fmt.Printf("%s\n", jsonData)
		case "yaml":
			response := []string{}
			for _, tag := range githubResult {
				response = append(response, tag.Name)
			}
			yamlData, err := yaml.Marshal(response)
			if err != nil {
				message := color.RedString(err.Error())
				panic(message)
			}
			fmt.Printf("%s", yamlData)

		case "text":
			for _, tag := range githubResult {
				color.Cyan(tag.Name)
			}
		case "table":
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl := table.New("Name")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, tag := range githubResult {
				tbl.AddRow(tag.Name)
			}
			tbl.Print()
		}
	},
}
