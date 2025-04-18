package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize project by supplying project name",
	Long:  `Initialize a new project structure by supplying a project name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Project initialized!")
	},
}
