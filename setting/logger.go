package setting

import (
	"history-engine/engine/utils"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Log = struct {
	File   string
	Level  string
	Format string
}{
	Level:  "info",
	File:   "",
	Format: "console",
}

func loadLogger() {
	v := viper.Sub("log")
	if v != nil {
		if v.IsSet("level") {
			Log.Level = v.GetString("level")
		}
		if v.IsSet("file") {
			Log.File = v.GetString("file")
		}
		if v.IsSet("format") {
			Log.Format = v.GetString("format")
		}
	}

	log.Printf("log level: %s, format: %s, log file: %s\n", Log.Level, Log.Format, Log.File)
	checkLogFile()
}

func checkLogFile() {
	if Log.File == "" || utils.FileExist(Log.File) {
		return
	}

	file, err := os.OpenFile(Log.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("create log file err:%v\n", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("close log file err:%v\n", err)
		}
	}(file)
}
