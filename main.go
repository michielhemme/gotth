package main

import (
	"github.com/michielhemme/gotth/cmd"
	"github.com/michielhemme/gotth/pkg/tools"
)

func main() {
	tools.InitializeTools()
	tools.InitializeConfiguration()
	cmd.Execute()
}
