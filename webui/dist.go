package webui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var f embed.FS

func Dist(path string) http.FileSystem {
	fs, err := fs.Sub(f, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fs)
}
