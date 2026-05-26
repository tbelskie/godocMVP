// Package project creates new Pendragon projects from an embedded skeleton.
//
// The package is small on purpose: no shelling out, no symlink
// traversal, no overwriting of existing paths, and every file write
// goes through O_EXCL. The caller (the CLI) owns user-facing output;
// this package just returns the absolute path it produced or an error
// describing exactly what went wrong.
package project

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Options configures a single Create call.
type Options struct {
	// Name is the user-supplied project name. It will be validated and
	// used as both the destination directory name and the default site
	// title.
	Name string

	// ParentDir is the directory under which the project directory is
	// created. Defaults to the process working directory.
	ParentDir string

	// Version is the godoc version string interpolated into generated
	// files (e.g. llms.txt). Empty string is fine.
	Version string
}

// Create writes a new project tree and returns the absolute path it
// created.
//
// The flow is intentionally linear and ordered cheap-checks-first so
// that bad input is rejected before we touch the filesystem:
//
//  1. validateName       — pure-CPU check on the user-supplied name.
//  2. resolveTarget      — compute the absolute target path.
//  3. ensureNothingAt    — refuse to clobber any existing path.
//  4. writeSkeleton      — create the directory and fill it; rolled
//                          back atomically on any failure.
//
// The caller therefore never sees a half-written project.
func Create(ctx context.Context, opts Options) (string, error) {
	name, err := validateName(opts.Name)
	if err != nil {
		return "", err
	}

	target, err := resolveTarget(opts.ParentDir, name)
	if err != nil {
		return "", err
	}

	if err := ensureNothingAt(target); err != nil {
		return "", err
	}

	data := templateData{
		ProjectName:  name,
		Title:        titleFromName(name),
		PendragonVersion: opts.Version,
	}
	if err := writeSkeleton(ctx, target, data); err != nil {
		return "", err
	}
	return target, nil
}

// resolveTarget computes the absolute path where the project should be
// created and confirms that path stays inside the parent directory.
// The subpath check is defense in depth — validateName already rejects
// path-traversal in `name`.
func resolveTarget(parentDir, name string) (string, error) {
	parent := parentDir
	if parent == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("determine current directory: %w", err)
		}
		parent = cwd
	}

	parentAbs, err := filepath.Abs(parent)
	if err != nil {
		return "", fmt.Errorf("resolve parent directory %q: %w", parent, err)
	}

	target := filepath.Join(parentAbs, name)
	if !isInside(parentAbs, target) {
		return "", fmt.Errorf("invalid target path: would escape parent directory %q", parentAbs)
	}
	return target, nil
}

// ensureNothingAt refuses to overwrite an existing path. Using Lstat
// (not Stat) means symlinks are rejected without being followed, so we
// can never write into a directory the user did not directly intend.
func ensureNothingAt(target string) error {
	info, err := os.Lstat(target)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("check target path %q: %w", target, err)
	}
	return fmt.Errorf("refusing to overwrite existing %s: %s", describePathKind(info), target)
}

// describePathKind returns a user-friendly noun for the file mode bits.
// Avoids surfacing raw mode strings like "d---------" in errors.
func describePathKind(info os.FileInfo) string {
	mode := info.Mode()
	switch {
	case mode&os.ModeSymlink != 0:
		return "symlink"
	case mode.IsDir():
		return "directory"
	case mode.IsRegular():
		return "file"
	default:
		return "path"
	}
}

func isInside(parent, child string) bool {
	rel, err := filepath.Rel(parent, child)
	if err != nil {
		return false
	}
	if rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return false
	}
	return true
}
