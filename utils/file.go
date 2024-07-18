package utils

import (
	"os"
)

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	if path == "" {
		return false
	}

	fi, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return !fi.IsDir()
}

// PathExist 判断目录是否存在
func PathExist(path string) bool {
	if path == "" {
		return false
	}

	fi, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return fi.IsDir()
}
