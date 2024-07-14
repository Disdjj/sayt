package main

import (
	"github.com/Disdjj/sayt/internal"
	"github.com/guonaihong/clop"
)

func main() {
	clop.SetVersion("v0.0.1")
	clop.SetVersionOption("v", "version")
	clop.ShowUsageDefault = false
	clop.SetAbout("Say it: LLM CLI Tool")
	err := clop.Bind(internal.Arg)
	if err != nil {
		panic(err)
		return
	}

	internal.InitConfig()
	internal.InitClient()

	if internal.Arg.File == "" {
		println("File is required")
		return
	}
	response, err := internal.CompleteWithFile(internal.Arg.File)
	if err != nil {
		println(err)
		return
	}
	println(response)
}
