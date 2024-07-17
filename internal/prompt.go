package internal

type PromptNode struct {
	Category      string
	SystemMessage string
}

var (
	CategoryMap map[string]*PromptNode
)

func loadAllPrompts() {
	CategoryMap = make(map[string]*PromptNode)

	CategoryMap["log"] = &PromptNode{
		Category:      "log",
		SystemMessage: "analyze logs and find the problem, output the result in markdown format",
	}

	CategoryMap["assistant"] = &PromptNode{
		Category:      "assistant",
		SystemMessage: "you are a useful assistant, help user with the following tasks",
	}

}

func init() {
	// Load all prompts
	loadAllPrompts()
}
