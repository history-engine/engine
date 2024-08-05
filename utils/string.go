package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Md5str(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Sha1Str(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func Str2Int(s string) int {
	i := 0
	for _, c := range s {
		i += int(c)
	}
	return i
}

func Ternary[T any](cond bool, true, false T) T {
	if cond {
		return true
	}
	return false
}
