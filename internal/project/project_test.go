package project

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
)

func TestValidateName_AcceptsValidNames(t *testing.T) {
	t.Parallel()
	valid := []string{"my-docs", "docs_v2", "Abc123", "x"}
	for _, n := range valid {
		n := n
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			if _, err := validateName(n); err != nil {
				t.Errorf("expected %q to be valid, got error: %v", n, err)
			}
		})
	}
}

func TestValidateName_RejectsInvalidNames(t *testing.T) {
	t.Parallel()
	cases := []struct {
		why   string
		input string
	}{
		{"empty", ""},
		{"whitespace only", "   "},
		{"single dot", "."},
		{"double dot", ".."},
		{"path traversal", "../etc"},
		{"absolute path", "/abs"},
		{"contains forward slash", "a/b"},
		{"contains backslash", `a\b`},
		{"leading dash looks like flag", "-foo"},
		{"leading dot is hidden", ".hidden"},
		{"contains space", "foo bar"},
		{"embedded traversal", "a/../b"},
		{"windows reserved con", "con"},
		{"windows reserved PRN", "PRN"},
		{"windows reserved NUL", "NUL"},
		{"windows reserved COM1", "COM1"},
		{"contains NUL byte", "hello\x00world"},
		{"oversize", strings.Repeat("a", maxNameLen+1)},
		{"non-ascii emoji", "emoji-🎉"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.why, func(t *testing.T) {
			t.Parallel()
			if _, err := validateName(tc.input); err == nil {
				t.Errorf("expected %q to be rejected, got nil error", tc.input)
			}
		})
	}
}

func TestValidateName_ErrorMessagesIncludeOffender(t *testing.T) {
	t.Parallel()
	cases := []struct {
		why      string
		input    string
		expected []string
	}{
		{
			why:      "shows the name and the offending character",
			input:    "my-project!",
			expected: []string{`"my-project!"`, `'!'`, "letters", "digits"},
		},
		{
			why:      "explains why a leading dash is rejected",
			input:    "-foo",
			expected: []string{`"-foo"`, "flag"},
		},
		{
			why:      "explains why a leading dot is rejected",
			input:    ".hidden",
			expected: []string{`".hidden"`, "hidden"},
		},
		{
			why:      "shows the limit on oversize names",
			input:    strings.Repeat("a", maxNameLen+1),
			expected: []string{"too long", "limit is 64"},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.why, func(t *testing.T) {
			t.Parallel()
			_, err := validateName(tc.input)
			if err == nil {
				t.Fatalf("expected error for %q, got nil", tc.input)
			}
			for _, want := range tc.expected {
				if !strings.Contains(err.Error(), want) {
					t.Errorf("error %q missing %q", err.Error(), want)
				}
			}
		})
	}
}

func TestTitleFromName_HumanizesSlugs(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"my-docs":   "My Docs",
		"docs_v2":   "Docs V2",
		"abc":       "Abc",
		"Already":   "Already",
		"foo-bar_v": "Foo Bar V",
	}
	for in, want := range cases {
		if got := titleFromName(in); got != want {
			t.Errorf("titleFromName(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCreate_WritesExpectedSkeleton(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	target, err := Create(context.Background(), Options{
		Name:      "my-docs",
		ParentDir: parent,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if want := filepath.Join(parent, "my-docs"); target != want {
		t.Errorf("target = %s, want %s", target, want)
	}

	expected := []string{
		"hugo.toml",
		"godoc.yaml",
		"llms.txt",
		".gitignore",
		"content/_index.md",
		"content/docs/_index.md",
		"content/docs/getting-started.md",
		"content/guides/_index.md",
		"content/api/_index.md",
		"content/changelog.md",
		"content/contributing.md",
		"archetypes/default.md",
		"layouts/index.html",
		"layouts/_default/baseof.html",
		"layouts/_default/single.html",
		"layouts/_default/list.html",
		"layouts/partials/head.html",
		"layouts/partials/header.html",
		"layouts/partials/footer.html",
		"layouts/partials/sidebar.html",
		"layouts/partials/helpful.html",
		"layouts/partials/godoc-mark.html",
		"assets/css/main.css",
		"assets/js/theme.js",
		"assets/img/godoc-mark.svg",
		"static/.gitkeep",
		"data/.gitkeep",
	}
	for _, rel := range expected {
		info, err := os.Stat(filepath.Join(target, rel))
		if err != nil {
			t.Errorf("missing file %s: %v", rel, err)
			continue
		}
		if info.IsDir() {
			t.Errorf("expected file at %s, got directory", rel)
		}
	}

	// .tmpl files should never appear in the output.
	for _, p := range []string{"hugo.toml.tmpl", "content/_index.md.tmpl"} {
		if _, err := os.Stat(filepath.Join(target, p)); err == nil {
			t.Errorf("unexpected .tmpl file present at %s", p)
		}
	}
}

func TestCreate_InterpolatesTitleIntoConfig(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	target, err := Create(context.Background(), Options{
		Name:      "my-docs",
		ParentDir: parent,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(target, "hugo.toml"))
	if err != nil {
		t.Fatalf("read hugo.toml: %v", err)
	}
	if !strings.Contains(string(raw), `title = "My Docs"`) {
		t.Errorf("hugo.toml missing humanized title; got:\n%s", raw)
	}
}

func TestCreate_RefusesToOverwriteExistingDirectory(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	existing := filepath.Join(parent, "site")
	if err := os.Mkdir(existing, 0o755); err != nil {
		t.Fatal(err)
	}
	_, err := Create(context.Background(), Options{
		Name:      "site",
		ParentDir: parent,
	})
	if err == nil {
		t.Fatal("expected refusal, got nil error")
	}
	if !strings.Contains(err.Error(), "refusing to overwrite existing directory") {
		t.Errorf("expected 'refusing to overwrite existing directory' in error, got: %v", err)
	}
}

func TestCreate_RefusesToOverwriteSymlink(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	other := t.TempDir()
	link := filepath.Join(parent, "site")
	if err := os.Symlink(other, link); err != nil {
		t.Skipf("symlink unsupported: %v", err)
	}
	_, err := Create(context.Background(), Options{
		Name:      "site",
		ParentDir: parent,
	})
	if err == nil {
		t.Fatal("expected refusal, got nil error")
	}
	if !strings.Contains(err.Error(), "refusing to overwrite existing symlink") {
		t.Errorf("expected 'refusing to overwrite existing symlink' in error, got: %v", err)
	}
}

func TestCreate_RejectsInvalidName_LeavesParentEmpty(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	_, err := Create(context.Background(), Options{
		Name:      "../escape",
		ParentDir: parent,
	})
	if err == nil {
		t.Fatal("expected validation error")
	}
	entries, _ := os.ReadDir(parent)
	if len(entries) != 0 {
		t.Errorf("parent should remain empty on rejection, got %v", entries)
	}
}

func TestCreate_FailsWhenParentIsNotADirectory(t *testing.T) {
	t.Parallel()
	parent := t.TempDir()
	blocker := filepath.Join(parent, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := Create(context.Background(), Options{
		Name:      "child",
		ParentDir: blocker,
	})
	if err == nil {
		t.Fatal("expected error when parent is a file")
	}
}

func TestCreate_RollsBackOnRenderFailure(t *testing.T) {
	t.Parallel()
	// Cancelling the context after directory creation but before any
	// file write forces renderInto to return early. The deferred
	// rollback in Create should then remove the target directory.
	parent := t.TempDir()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := Create(ctx, Options{
		Name:      "rollback",
		ParentDir: parent,
	})
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
	if _, statErr := os.Stat(filepath.Join(parent, "rollback")); !errors.Is(statErr, fs.ErrNotExist) {
		t.Errorf("expected target directory to be rolled back, stat err = %v", statErr)
	}
}

func TestSkeletonFS_IncludesDotfiles(t *testing.T) {
	t.Parallel()
	// Sanity check that the `all:` prefix on //go:embed is still in
	// place; without it, .gitignore.tmpl and .gitkeep files would be
	// silently dropped.
	if _, err := fs.ReadFile(skeletonFS, "templates/.gitignore.tmpl"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			t.Error(".gitignore.tmpl not embedded — did you drop the 'all:' prefix on //go:embed?")
		} else {
			t.Error(err)
		}
	}
}

func TestSkeletonFS_HasEnoughFiles(t *testing.T) {
	t.Parallel()
	var count int
	err := fs.WalkDir(skeletonFS, skeletonRoot, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if count < 12 {
		t.Errorf("expected at least 12 embedded skeleton files, got %d", count)
	}
}

// TestEmbeddedLayouts_ParseAsTemplates parses every embedded Hugo
// layout/partial with Go's text/template parser. Hugo's template
// syntax is a strict superset of Go's, so successful parsing here
// rules out the most common typos (unbalanced braces, unclosed
// blocks, malformed actions) at unit-test time even when Hugo is
// not available on PATH.
//
// Hugo-specific functions are registered as no-op stubs so the
// parser recognizes them as function calls; their return values
// are never used (parse-only check).
func TestEmbeddedLayouts_ParseAsTemplates(t *testing.T) {
	t.Parallel()
	stub := func(args ...interface{}) interface{} { return nil }
	funcs := template.FuncMap{
		"partial":   stub,
		"anchorize": stub,
		"now":       stub,
		"site":      stub,
		"resources": stub,
		"default":   stub,
	}
	layoutsRoot := skeletonRoot + "/layouts"
	walkErr := fs.WalkDir(skeletonFS, layoutsRoot, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}
		body, err := fs.ReadFile(skeletonFS, path)
		if err != nil {
			return err
		}
		if _, err := template.New(path).Funcs(funcs).Parse(string(body)); err != nil {
			t.Errorf("parse %s: %v", path, err)
		}
		return nil
	})
	if walkErr != nil {
		t.Fatal(walkErr)
	}
}

// TestScaffoldBuildsWithHugo is the end-to-end check that what we
// scaffold actually renders through real Hugo. It is skipped when
// hugo is not on PATH so the unit-test build remains hermetic; on
// developer machines and any CI runner with Hugo installed, it
// guards against template regressions that the parse-only test
// cannot catch (missing partials, bad asset pipeline calls, etc.).
func TestScaffoldBuildsWithHugo(t *testing.T) {
	t.Parallel()
	hugoBin, err := exec.LookPath("hugo")
	if err != nil {
		t.Skip("hugo not on PATH; skipping real-Hugo integration build")
	}

	parent := t.TempDir()
	target, err := Create(context.Background(), Options{
		Name:      "demo-site",
		ParentDir: parent,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	cmd := exec.Command(hugoBin, "--minify", "--quiet")
	cmd.Dir = target
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("hugo build failed: %v\noutput:\n%s", err, out)
	}

	checks := []string{
		"public/index.html",
		"public/docs/index.html",
		"public/docs/getting-started/index.html",
		"public/guides/index.html",
		"public/api/index.html",
		"public/changelog/index.html",
		"public/contributing/index.html",
	}
	for _, rel := range checks {
		body, err := os.ReadFile(filepath.Join(target, rel))
		if err != nil {
			t.Errorf("missing %s: %v", rel, err)
			continue
		}
		if len(body) < 200 {
			t.Errorf("%s body too short (%d bytes); did the template render empty?", rel, len(body))
		}
	}

	home, err := os.ReadFile(filepath.Join(target, "public/index.html"))
	if err != nil {
		t.Fatal(err)
	}
	homeStr := string(home)
	if !strings.Contains(homeStr, "Demo Site") {
		t.Errorf("homepage missing humanized title 'Demo Site'; got first 400 bytes:\n%s", truncate(homeStr, 400))
	}
	if !strings.Contains(homeStr, "godoc-mark") {
		t.Errorf("homepage missing brand mark partial output")
	}
	if !strings.Contains(homeStr, "data-theme-toggle") {
		t.Errorf("homepage missing theme toggle button")
	}
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
