package utils

import (
	"time"
)

func CheckVersionInterval(minVersionInterval int, lastTime time.Time) bool {
	end := lastTime.Add(time.Duration(minVersionInterval) * time.Second)
	comp := time.Now().Compare(end)
	return comp == -1
}
