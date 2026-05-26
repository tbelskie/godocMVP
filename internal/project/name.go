package project

import (
	"fmt"
	"strings"
)

const maxNameLen = 64

// reservedNames are forbidden on Windows. We reject them everywhere so
// generated projects stay portable across operating systems.
var reservedNames = map[string]struct{}{
	"con": {}, "prn": {}, "aux": {}, "nul": {},
	"com1": {}, "com2": {}, "com3": {}, "com4": {}, "com5": {},
	"com6": {}, "com7": {}, "com8": {}, "com9": {},
	"lpt1": {}, "lpt2": {}, "lpt3": {}, "lpt4": {}, "lpt5": {},
	"lpt6": {}, "lpt7": {}, "lpt8": {}, "lpt9": {},
}

// validateName enforces a strict, conservative allowlist for project
// names. ASCII-only is intentional: it sidesteps Unicode normalization
// surprises, filesystem case-folding, and shell-quoting pitfalls. The
// rule can be relaxed later if real users need it.
//
// Errors echo the offending input so users can see what they typed and
// which rule it tripped — the most common UX win for CLI validation.
func validateName(raw string) (string, error) {
	name := strings.TrimSpace(raw)
	switch {
	case name == "":
		return "", fmt.Errorf("project name is required")
	case len(name) > maxNameLen:
		return "", fmt.Errorf("project name %q is too long (%d characters; limit is %d)", name, len(name), maxNameLen)
	case strings.ContainsAny(name, `/\`):
		return "", fmt.Errorf("project name %q may not contain path separators ('/' or '\\')", name)
	case strings.ContainsRune(name, 0):
		return "", fmt.Errorf("project name may not contain NUL bytes")
	case name == "." || name == ".." || strings.Contains(name, ".."):
		return "", fmt.Errorf("project name %q may not contain '.' or '..'", name)
	case strings.HasPrefix(name, "-"):
		return "", fmt.Errorf("project name %q may not start with '-' (looks like a flag)", name)
	case strings.HasPrefix(name, "."):
		return "", fmt.Errorf("project name %q may not start with '.' (would create a hidden directory)", name)
	}
	if _, reserved := reservedNames[strings.ToLower(name)]; reserved {
		return "", fmt.Errorf("project name %q is reserved on Windows; pick another", name)
	}
	if bad, ok := firstDisallowedRune(name); ok {
		return "", fmt.Errorf("project name %q contains invalid character %q; allowed: ASCII letters, digits, '-', '_'", name, bad)
	}
	return name, nil
}

// firstDisallowedRune returns the first rune in name that is not in the
// allowlist, plus a bool indicating whether one was found.
func firstDisallowedRune(name string) (rune, bool) {
	for _, r := range name {
		if !isAllowedNameRune(r) {
			return r, true
		}
	}
	return 0, false
}

func isAllowedNameRune(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z':
		return true
	case r >= 'A' && r <= 'Z':
		return true
	case r >= '0' && r <= '9':
		return true
	case r == '-' || r == '_':
		return true
	}
	return false
}

// titleFromName turns "my-cool_project" into "My Cool Project" for use
// as a default site title. Safe to do with ASCII byte slicing because
// validateName already restricted the input to that set.
func titleFromName(name string) string {
	cleaned := strings.NewReplacer("-", " ", "_", " ").Replace(name)
	words := strings.Fields(cleaned)
	for i, w := range words {
		words[i] = capitalizeASCII(w)
	}
	return strings.Join(words, " ")
}

func capitalizeASCII(w string) string {
	if w == "" {
		return w
	}
	return strings.ToUpper(w[:1]) + w[1:]
}
