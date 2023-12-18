package main

import (
	"github.com/yangjerry110/tool/internal/cmd/command"
	"github.com/yangjerry110/tool/internal/cmd/config"
)

func main() {

	// set projectPath
	if err := config.CreateConfig(&config.ProjectPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// set projectImportPath
	if err := config.CreateConfig(&config.ProjectImportPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// set cli app
	if err := command.CreateCommand(&command.CliNewApp{}).New(); err != nil {
		panic(err)
	}
}
