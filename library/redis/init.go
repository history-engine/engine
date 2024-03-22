package redis

import "history-engine/engine/library/wait"

var enable = false

func EnableRedis() {
	if enable {
		return
	}

	wait.AddWait(1)
	enable = true
	initEngine()
	initSignal()
}
