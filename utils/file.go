package utils

import "os"

func FileExists(path string) bool {
	if fi, err := os.Stat(path); fi != nil && err != os.ErrNotExist {
		return !fi.IsDir()
	}

	return false
}
