package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotth",
	Short: "A CLI tool for project automation",
	Long:  "gotth is a developer tool that helps initialize and manage projects",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	// Subcommands are added here
	rootCmd.AddCommand(initializeCmd)
	rootCmd.AddCommand(airCmd)
	rootCmd.AddCommand(versionCmd)
}
