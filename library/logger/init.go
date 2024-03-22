package logger

var enable = false

func EnableLogger() {
	if enable {
		return
	}

	if err := initZap(); err != nil {
		return
	}

	enable = true
}
