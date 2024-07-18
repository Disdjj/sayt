package internal

import (
	"encoding/json"
	"github.com/go-git/go-git/v5"
	"os"
	"path/filepath"
)

var RepoConfig []PromptConfig

type PromptConfig struct {
	Repo       string   `json:"repo"`
	Categories []string `json:"categories"`
}

func InstallRemoteRepo(name, url string) {
	// url like https://github.com/user/repo
	// folder name will be user-repo

	// Clone the given repository to the given directory
	_, err := git.PlainClone(
		GetProjectPath()+"/prompts/"+name, false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		},
	)
	if err != nil {
		panic(err)
	}
	BuildRepoConfig()
}

func RemoveRepo(name string) {
	// url like https://github.com/user/repo
	// folder name will be user-repo

	// Clone the given repository to the given directory

	path := GetProjectPath() + "/prompts/" + name
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
	BuildRepoConfig()
}

func BuildRepoConfig() []PromptConfig {
	path := GetProjectPath() + "/prompts/"
	// walk Path like: /prompts/<repoName>/<promptName>
	result := make([]PromptConfig, 0)
	err := filepath.Walk(
		path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && filepath.Base(filepath.Dir(filepath.Dir(path))) == "prompts" {
				// check if it has system.md
				_, err := os.Stat(path + "/system.md")
				if err == nil {
					// get the repo name
					repoName := filepath.Base(filepath.Dir(path))
					// get the prompt name
					promptName := filepath.Base(path)
					// check if the repo is already in the result
					var found bool
					for i := range result {
						if result[i].Repo == repoName {
							result[i].Categories = append(result[i].Categories, promptName)
							found = true
							break
						}
					}
					if !found {
						result = append(
							result, PromptConfig{
								Repo:       repoName,
								Categories: []string{promptName},
							},
						)
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}
	// save the result to the config file

	path = GetProjectPath() + "/prompts.json"
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(result)
	if err != nil {
		panic(err)
	}
	return result
}

func ListRepo() []string {
	result := make([]string, 0)
	for _, config := range RepoConfig {
		result = append(result, config.Repo)
	}
	return result
}

func ListCategories(repo string) []string {
	result := make([]string, 0)
	if repo == "*" {
		for _, config := range RepoConfig {
			result = append(result, config.Categories...)
		}
		return result
	}

	for _, config := range RepoConfig {
		if config.Repo == repo {
			result = config.Categories
			return result
		}
	}
	return result
}

func init() {
	// Load all prompts
	RepoConfig = BuildRepoConfig()
}
