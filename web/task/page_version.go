package task

import (
	"history-engine/engine/setting"
	"time"
)

// RunPageVersionCheck 页面版本检查, 每5分钟检查一次
func RunPageVersionCheck() {
	t := time.NewTicker(time.Duration(setting.SingleFile.VersionCheckInterval) * time.Second)
	for {
		select {
		case <-t.C:
			for i := 0; i < setting.SingleFile.VersionCheckLimit; i++ {
				// todo
			}
		}
	}
}
