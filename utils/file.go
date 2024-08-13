package utils

import (
	"fmt"
	"math"
	"net/url"
	"os"
	"strings"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

var sizeUnits = []string{" Bytes", " KB", " MB", " GB", " TB", " PB", " EB", " ZB", " YB"}

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

// SizeFormat 文件大小格式化
func SizeFormat[T Number](size T) string {
	if size == 0 {
		return "0 Bytes"
	}

	i := int(math.Floor(math.Log(float64(size)) / math.Log(1024)))
	sizeFormatted := float64(size) / math.Pow(1024, float64(i))
	if sizeFormatted == float64(int(sizeFormatted)) {
		return fmt.Sprintf("%d%s", int(sizeFormatted), sizeUnits[i])
	}
	return fmt.Sprintf("%.2f%s", sizeFormatted, sizeUnits[i])
}
