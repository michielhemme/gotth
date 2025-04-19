package cmd

import (
	"github.com/michielhemme/gotth/pkg/boilerplate"
	"github.com/michielhemme/gotth/pkg/logger"
	"github.com/spf13/cobra"
)

var subdir bool

var initializeCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize project by supplying project name",
	Long:  `Initialize a new project structure by supplying a project name.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := boilerplate.InitializeProject(args[0], subdir); err != nil {
			logger.Log(1, err)
		}
	},
}

func init() {
	initializeCmd.Flags().BoolVar(&subdir, "subdir", false, "Create project in a subdirectory")
}
