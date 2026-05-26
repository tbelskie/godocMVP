package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// version is overridable at build time via -ldflags.
var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pendragon",
	Short: "Pendragon — The AI-powered DocOps assistant",
	Long: `Pendragon is the AI-powered DocOps assistant for documentation teams.

Create beautiful Hugo documentation sites in seconds, then keep them healthy
with audit, fix, and polish workflows designed for docs-as-code repos.`,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "Pendragon — The AI-powered DocOps assistant")
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
