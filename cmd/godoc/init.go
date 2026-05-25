package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tbelskie/godocMVP/internal/project"
)

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new godoc-powered Hugo documentation site",
	Long: `init scaffolds a complete Hugo project layout with smart defaults:

  - hugo.toml + godoc.yaml configuration
  - content/ with Docs, Guides, and API sections
  - llms.txt at the project root for AI/machine-readable consumers
  - archetypes/, layouts/, assets/, static/, data/ ready for use

The target directory must not already exist.

Example:
  godoc init my-docs`,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		target, err := project.Create(cmd.Context(), project.Options{
			Name:    name,
			Version: version,
		})
		if err != nil {
			return err
		}
		out := cmd.OutOrStdout()
		fmt.Fprintf(out, "Created godoc project at %s\n\n", target)
		fmt.Fprintln(out, "Next steps:")
		fmt.Fprintf(out, "  cd %s\n", name)
		fmt.Fprintln(out, "  hugo server")
		return nil
	},
}
