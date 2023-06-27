package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "Brahma CLI",
		Long: `Brahma CLI`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
