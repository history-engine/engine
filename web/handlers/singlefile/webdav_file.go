package singlefile

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/webdav"
)

type Dir struct {
	webdav.Dir
}

func (d Dir) MkdirAll(ctx context.Context, name string, perm os.FileMode) error {
	if name = d.resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.MkdirAll(name, perm)
}

func (d Dir) resolve(name string) string {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) || strings.Contains(name, "\x00") {
		return ""
	}
	dir := string(d.Dir)
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, filepath.FromSlash(slashClean(name)))
}

func slashClean(name string) string {
	if name == "" || name[0] != '/' {
		name = "/" + name
	}
	return path.Clean(name)
}
