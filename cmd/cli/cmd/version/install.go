package version

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/ryurock/cli/pkg/config"
	"github.com/spf13/cobra"
)

var versionInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "指定したバージョンでコマンドをインストールします",
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

		tags := []string{}
		for _, tag := range githubResult {
			tags = append(tags, tag.Name)
		}

		prompt := promptui.Select{
			Label: color.CyanString("ダウンロードするバージョンを選択してください"),
			Items: tags,
		}

		_, version, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", version)
		ghArgs := []string{"release", "download", version, "--repo", fmt.Sprintf("%s/%s", githubOraganization, githubRepoName), "--pattern", "macOS-arm64.zip"}
		stdOut, stdErr, err := gh.Exec(ghArgs...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(stdOut.String())
		fmt.Println(stdErr.String())
	},
}
