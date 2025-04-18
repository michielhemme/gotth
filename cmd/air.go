package cmd

import (
	"fmt"
	"path"

	"github.com/michielhemme/gotth/pkg/lib"
	"github.com/michielhemme/gotth/pkg/logger"
	"github.com/michielhemme/gotth/pkg/tools"
	"github.com/spf13/cobra"
)

var airCmd = &cobra.Command{
	Use:   "air",
	Short: "Run the air service",
	Long:  `Start the air service to auto-reload your application during development.`,
	Run: func(cmd *cobra.Command, args []string) {
		executable := tools.GetExecutable("air")
		err := tools.ExecuteCommand(executable, "-c", path.Join(lib.GetCacheDir(), "air.toml"))
		if err != nil {
			logger.Log(1, fmt.Sprintf("Could not execute command for air service: %v", err))
		}
	},
}
