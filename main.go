package main

import (
	"github.com/AlecAivazis/survey/v2"
	"os"
)

const txtTaskGenerateStruct = "生成数据struct"
const txtTaskExit = "退出"

func main() {
	var taskTag string
	prompt := &survey.Select{
		Message: "你想",
		Options: []string{txtTaskGenerateStruct, txtTaskExit},
	}
	err := survey.AskOne(
		prompt,
		&taskTag,
	)
	if err == nil {
		switch taskTag {
		case txtTaskGenerateStruct:
			genStruct()
		case txtTaskExit:
			os.Exit(0)
		}
	}
}
