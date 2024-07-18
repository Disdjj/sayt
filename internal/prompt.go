package internal

import (
	"io"
	"os"
)

type PromptNode struct {
	Repo     string
	Category string
	Help     string
	Path     PromptPathString
}

type PromptPathString string

func (p PromptPathString) GetSystemMessage() string {
	finalPath := GetProjectPath() + "/prompts/" + string(p) + "/system.md"
	// Read the file
	reader, err := os.Open(finalPath)
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	sysStr, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return string(sysStr)
}

var (
	CategoryMap map[string]*PromptNode
)

func loadAllPrompts() {
	CategoryMap = make(map[string]*PromptNode)
	for _, repoNode := range RepoConfig {
		for _, category := range repoNode.Categories {
			promptNode := &PromptNode{
				Repo:     repoNode.Repo,
				Category: category,
				Path:     PromptPathString(repoNode.Repo + "/" + category),
			}
			CategoryMap[category] = promptNode
		}
	}

}
