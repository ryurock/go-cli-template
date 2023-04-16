package config

type CliConfig struct {
	Version string `default:"0.0.4"`
	GitHub  CliConfigGithub
}

type CliConfigGithub struct {
	Repo CliConfigGithubRepo
}

type CliConfigGithubRepo struct {
	Organization string `default:"ryurock"`
	Name         string `default:"cli"`
}

func NewCliConfig() *CliConfig {
	config := &CliConfig{
		Version: "0.0.4",
		GitHub: CliConfigGithub{
			Repo: CliConfigGithubRepo{
				Organization: "ryurock",
				Name:         "go-cli-template",
			},
		},
	}

	return config
}
