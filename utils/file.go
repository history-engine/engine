package utils

import (
	"net/url"
	"os"
	"strings"
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

func FileSuffix(name string) string {
	parsed, err := url.Parse(name)
	if err != nil || parsed == nil {
		return ""
	}

	suffix := strings.Split(parsed.Path, ".")
	if len(suffix) > 1 {
		return suffix[len(suffix)-1]
	}
	return parsed.Path
}
