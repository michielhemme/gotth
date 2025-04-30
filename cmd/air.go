package cmd

import (
	"path"
	"path/filepath"

	"github.com/michielhemme/gotth/pkg/lib"
	"github.com/michielhemme/gotth/pkg/logger"
	"github.com/spf13/cobra"
)

var airCmd = &cobra.Command{
	Use:   "air",
	Short: "Run the air service",
	Long:  `Start the air service to auto-reload your application during development.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.RunCommand(lib.Command{
			WorkingDir: "",
			Program:    "go",
			Args:       []string{"build", "-o", filepath.Join("tmp", lib.AppendIfExe("main")), "."},
		}); err != nil {
			logger.Log(1, err)
		}
		cacheDir, err := lib.GetCacheDir()
		if err != nil {
			logger.Log(1, err)
		}
		if err := lib.RunCommand(lib.Command{
			WorkingDir: "",
			Program:    "air",
			Args:       []string{"-c", path.Join(cacheDir, "air.toml")},
		}); err != nil {
			logger.Log(1, err)
		}
	},
}
