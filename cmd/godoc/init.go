package godoc

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new godocMVP-powered Hugo documentation site",
	Long: `init creates a complete Hugo site with:
- Smart Information Architecture (IA)
- Our premium godocMVP theme
- Pagefind search
- Sample content
- Deploy-ready setup

Example:
  godocMVP init my-docs`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: project name required")
			fmt.Println("Usage: godocMVP init <project-name>")
			return
		}
		projectName := args[0]
		fmt.Printf("🚀 Initializing godocMVP site: %s\n", projectName)
		fmt.Println("✅ Site created! (stub - full implementation coming in next step)")
		// TODO: Call internal/template to embed and copy the Hugo skeleton
	},
}

func init() {
	// This is called automatically by root.go
}
