package setting

import (
	"history-engine/engine/utils"
	"log"
	"os"

	"github.com/spf13/viper"
)

var Log = struct {
	File  string
	Level string
}{
	Level: "info",
	File:  "data/runtime.log",
}

func loadLogger() {
	v := viper.Sub("log")
	if v != nil {
		if v.IsSet("log_level") {
			Log.Level = v.GetString("log_level")
		}
		if v.IsSet("log_path") {
			Log.File = v.GetString("log_file")
		}
	}

	log.Printf("log level:%s, log file:%s\n", Log.Level, Log.File)
	checkLogFile()
}

func checkLogFile() {
	if utils.FileExist(Log.File) {
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
