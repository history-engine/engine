package setting

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	SingleFile = struct {
		Path                 string
		MaxVersion           int // todo 可以按天周月年保留
		VersionCheckInterval int
		VersionCheckLimit    int
		IgnoreHost           []string
	}{
		Path:                 "",
		MaxVersion:           5,
		VersionCheckInterval: 300,
		VersionCheckLimit:    100,
		IgnoreHost:           []string{},
	}
)

func init() {
	home, err := os.UserConfigDir()
	if err == nil {
		SingleFile.Path = home + "/history-engine/html"
	}
}

func loadSingleFile() {
	v := viper.Sub("singlefile")
	SingleFile.Path = v.GetString("path")
	SingleFile.MaxVersion = v.GetInt("max_version")
	SingleFile.VersionCheckInterval = v.GetInt("version_check_interval")
	SingleFile.VersionCheckLimit = v.GetInt("version_check_limit")
	SingleFile.IgnoreHost = v.GetStringSlice("ignore_host")
	checkStoragePath()
}

func checkStoragePath() {
	if _, err := os.Stat(SingleFile.Path); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(SingleFile.Path, 0755)
		if err != nil {
			log.Fatalf("initialize storage path err:%v", err)
		}
	}
}
