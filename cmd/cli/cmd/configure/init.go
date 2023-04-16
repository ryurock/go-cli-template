package configure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/ryurock/cli/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Configure struct {
	WorkSpaces []string
	Versions   []ConfigureVersion
	Email      string
	ACM        ConfigureACM
}

type ConfigureVersion struct {
	Activate bool
	Path     string
}

type ConfigureACM struct {
	Type string
}

type ACM struct {
	Name string
	Role string
}

var configureInitCmd = &cobra.Command{
	Use:   "init",
	Short: "CLI の設定を初期化するコマンド",
	Run: func(cmd *cobra.Command, args []string) {
		xdgConfigDir, err := os.UserConfigDir()
		if err != nil {
			panic(color.RedString(err.Error()))
		}

		cliConfig := config.NewCliConfig()
		xdgConfigPath := filepath.Join(xdgConfigDir, cliConfig.GitHub.Repo.Organization, "config.yaml")
		_, err = os.Stat(xdgConfigPath)
		if err != nil && os.IsNotExist(err) {
			err, isCreatedFile := confirmCreateConfigFile(xdgConfigPath)
			if err != nil {
				panic(color.RedString(err.Error()))
			}
			email := inputEmail()
			acm := selectAccessControlManagement()
			fmt.Println(acm.Role)

			if isCreatedFile {
				os.MkdirAll(filepath.Dir(xdgConfigPath), 0755)
				configYaml := Configure{
					WorkSpaces: []string{},
					Versions:   []ConfigureVersion{},
					Email:      email,
					ACM:        ConfigureACM{Type: acm.Role},
				}

				yamlData, err := yaml.Marshal(&configYaml)
				if err != nil {
					panic(color.RedString(err.Error()))
				}
				fmt.Println(string(yamlData))
				fp, err := os.Create(xdgConfigPath)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer fp.Close()

				fp.WriteString(string(yamlData))
			}
		}
	},
}

func selectAccessControlManagement() ACM {
	acms := []ACM{
		{Name: "正社員", Role: "role/regular.employee"},
		{Name: "業務委託", Role: "role/outsourcing"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Role | cyan }} ({{ .Name | yellow }})",
		Inactive: "{{ .Role | cyan }} ({{ .Name | faint }})",
		Selected: "{{ .Role | cyan | red }}",
		Details: `
--------- Access Control Management ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Role:" | faint }}	{{ .Role }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := acms[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     color.YellowString("契約体系を選択してください"),
		Items:     acms,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err.Error())
	}

	return acms[i]
}

/**
 * Email を入力する
 *
 * @return string
 **/
func inputEmail() string {
	prompt := promptui.Prompt{
		Label: "Email を入力してください",
	}

	result, err := prompt.Run()
	if err != nil {
		panic(color.RedString(err.Error()))
	}

	return result
}

func confirmCreateConfigFile(xdgConfigPath string) (error, bool) {
	message := color.YellowString(fmt.Sprintf("%s に設定が存在しませんでした。ファイルを作成します。", xdgConfigPath))
	prompt := promptui.Select{
		Label: message,
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		panic(color.RedString(err.Error()))
	}

	if result == "Yes" {
		return nil, true
	}

	return nil, false
}
