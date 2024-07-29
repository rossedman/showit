package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// top level instantion of the CLI
var rootCmd = &cobra.Command{
	Use:   "showit",
	Short: "showit is a simple example of a go test parser",
}

func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
