package godoc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godoc",
	Short: "godoc — Instant beautiful Hugo documentation sites",
	Long: `godoc is the fastest way to create professional Hugo-based documentation sites.
One command turns nothing into a full, beautiful, searchable docs site.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("✨ godoc — Docs as Code, done right")
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(initCmd)
}
