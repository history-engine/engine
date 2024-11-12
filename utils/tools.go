package utils

func Ternary[T any](cond bool, true, false T) T {
	if cond {
		return true
	}
	return false
}
