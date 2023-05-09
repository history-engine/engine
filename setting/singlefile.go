package setting

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	SingleFile = struct {
		Path                 string
		MaxVersion           int
		VersionCheckInterval int
		VersionCheckLimit    int
	}{
		Path:                 "",
		MaxVersion:           5,
		VersionCheckInterval: 300,
		VersionCheckLimit:    100,
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
	checkStoragePath()
}

func checkStoragePath() {
	log.Printf("singlefile storage path:%s\n", SingleFile.Path)
	if _, err := os.Stat(SingleFile.Path); err == os.ErrNotExist {
		err = os.MkdirAll(SingleFile.Path, 0755)
		if err != nil {
			log.Fatalf("initialize storage path err:%v", err)
		}
	}
}
