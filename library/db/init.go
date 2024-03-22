package db

var enable = false

func EnableDb() {
	if enable {
		return
	}

	if err := initEngine(); err != nil {
		return
	}

	enable = true
}
