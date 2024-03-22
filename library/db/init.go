package db

import "history-engine/engine/library/wait"

var enable = false

func EnableDb() {
	if enable {
		return
	}

	wait.AddWait(1)
	enable = true
	initEngine()
	initSignal()
}
