package webui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var f embed.FS

func Dist(path string) http.FileSystem {
	fsSub, err := fs.Sub(f, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fsSub)
}
