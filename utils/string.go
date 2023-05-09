package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5str(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Str2Int(s string) int {
	i := 0
	for _, c := range s {
		i += int(c)
	}
	return i
}
