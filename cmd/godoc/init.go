package godoc

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new godoc-powered Hugo documentation site",
	Long: `init creates a complete Hugo site with:
- Smart Information Architecture (IA)
- Our premium godoc theme
- Pagefind search
- Sample content
- Deploy-ready setup

Example:
  godoc init my-docs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: project name required")
			fmt.Println("Usage: godoc init <project-name>")
			return
		}
		projectName := args[0]
		fmt.Printf("🚀 Initializing godoc site: %s\n", projectName)
		fmt.Println("✅ Site created! (stub - full implementation coming in next step)")
		// TODO: Call internal/template to embed and copy the Hugo skeleton
	},
}

func init() {
	// This is called automatically by root.go
}
