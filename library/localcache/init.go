package localcache

import "history-engine/engine/library/wait"

var enable = false

func EnableLocalCache() {
	if enable {
		return
	}

	wait.AddWait(1)
	enable = true
	initGoCache()
	initSignal()
}
