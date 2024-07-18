package setting

import (
	"errors"
	"github.com/spf13/viper"
	"history-engine/engine/utils"
	"log"
	"os"
)

var (
	SingleFile = struct {
		HtmlPath             string
		MaxVersion           int // todo 可以按天周月年保留
		MinVersionInterval   int // todo 可针对性配置，可动态
		VersionCheckInterval int
		VersionCheckLimit    int
		IgnoreHost           []string
	}{
		HtmlPath:             "data/html",
		MaxVersion:           5,
		MinVersionInterval:   86400,
		VersionCheckInterval: 300,
		VersionCheckLimit:    100,
		IgnoreHost:           []string{},
	}
)

func loadSingleFile() {
	v := viper.Sub("singlefile")
	if v != nil {
		if v.IsSet("html_path") {
			SingleFile.HtmlPath = v.GetString("html_path")
		}
		if v.IsSet("max_version") {
			SingleFile.MaxVersion = v.GetInt("max_version")
		}
		if v.IsSet("min_version_interval") {
			SingleFile.MinVersionInterval = v.GetInt("min_version_interval")
		}
		if v.IsSet("version_check_interval") {
			SingleFile.VersionCheckInterval = v.GetInt("version_check_interval")
		}
		if v.IsSet("version_check_limit") {
			SingleFile.VersionCheckLimit = v.GetInt("version_check_limit")
		}
		if v.IsSet("ignore_host") {
			SingleFile.IgnoreHost = v.GetStringSlice("ignore_host")
		}
	}

	checkStoragePath()
}

func checkStoragePath() {
	if utils.PathExist(SingleFile.HtmlPath) {
		log.Printf("html stroage paht: %s\n", SingleFile.HtmlPath)
		return
	}

	if _, err := os.Stat(SingleFile.HtmlPath); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(SingleFile.HtmlPath, 0755)
		if err != nil {
			log.Fatalf("create html storage path err:%v\n", err)
		}
	}
}
