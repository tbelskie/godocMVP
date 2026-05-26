package project

import "embed"

// skeletonFS holds the embedded project template tree.
//
// The `all:` prefix is required so dotfiles (e.g. .gitignore, .gitkeep)
// are included; Go's default //go:embed otherwise skips paths whose
// names begin with `.` or `_`.
//
//go:embed all:templates
var skeletonFS embed.FS

const skeletonRoot = "templates"
