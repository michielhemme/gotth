package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long:  `Display version of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s", Version)
	},
}
