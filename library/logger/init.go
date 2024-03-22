package logger

import "history-engine/engine/library/wait"

var enable = false

func EnableLogger() {
	if enable {
		return
	}

	wait.AddWait(1)
	enable = true
	initZap()
	initSignal()
}
