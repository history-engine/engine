package setting

import (
	"os"

	"github.com/spf13/viper"
)

var Log = struct {
	Path  string
	Level string
}{
	Level: "info",
	Path:  "",
}

func loadLogger() {
	v := viper.Sub("log")
	Log.Level = v.GetString("log_level")
	Log.Path = v.GetString("log_path")
	checkFileExist()
}

func checkFileExist() {
	if Log.Path == "" {
		return
	}
	if !fileExist(Log.Path) {
		if err := createFile(Log.Path); err != nil {
			panic(err)
		}
	}
}

func createFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	return nil
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
