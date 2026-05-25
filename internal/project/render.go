package project

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const templateSuffix = ".tmpl"

// templateData is the only data exposed to embedded templates. Keep
// this list short so templates stay easy to audit.
type templateData struct {
	ProjectName  string
	Title        string
	GodocVersion string
}

// writeSkeleton creates target, copies the embedded skeleton into it,
// and removes the directory if anything fails.
//
// The success-flag defer (rather than an explicit error branch) keeps
// rollback automatic if future code adds new failure paths between the
// Mkdir and the final return.
func writeSkeleton(ctx context.Context, target string, data templateData) error {
	if err := os.Mkdir(target, 0o755); err != nil {
		return fmt.Errorf("create project directory %q: %w", target, err)
	}

	success := false
	defer func() {
		if !success {
			_ = os.RemoveAll(target)
		}
	}()

	if err := copyTree(ctx, target, data); err != nil {
		return err
	}

	success = true
	return nil
}

// copyTree walks the embedded skeleton tree and writes each entry into
// the target directory. Files with the ".tmpl" suffix are rendered with
// text/template; everything else is copied verbatim.
func copyTree(ctx context.Context, target string, data templateData) error {
	return fs.WalkDir(skeletonFS, skeletonRoot, func(srcPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if err := ctx.Err(); err != nil {
			return err
		}

		rel, err := filepath.Rel(skeletonRoot, srcPath)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		outRel := strings.TrimSuffix(rel, templateSuffix)
		// Defense in depth: embedded paths come from our own repo, but
		// reject any traversal we did not intend to ship.
		if strings.Contains(outRel, "..") {
			return fmt.Errorf("embedded template has invalid path: %s", rel)
		}
		outPath := filepath.Join(target, outRel)

		if entry.IsDir() {
			return os.MkdirAll(outPath, 0o755)
		}
		return writeFile(srcPath, outPath, rel, data)
	})
}

// writeFile copies one entry from the embedded skeleton onto disk,
// rendering it as a text/template if the source has the ".tmpl"
// suffix.
func writeFile(srcPath, outPath, rel string, data templateData) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
		return err
	}

	contents, err := fs.ReadFile(skeletonFS, srcPath)
	if err != nil {
		return err
	}

	// O_EXCL guarantees we never silently overwrite a file. Combined
	// with the directory-level overwrite check in Create, this means
	// a half-written project is the worst case — and writeSkeleton's
	// deferred rollback handles that.
	f, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	if !strings.HasSuffix(rel, templateSuffix) {
		_, err = f.Write(contents)
		return err
	}

	tmpl, err := template.New(rel).Option("missingkey=error").Parse(string(contents))
	if err != nil {
		return fmt.Errorf("parse template %s: %w", rel, err)
	}
	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("render template %s: %w", rel, err)
	}
	return nil
}
