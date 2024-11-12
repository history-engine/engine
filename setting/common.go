package setting

import (
	_ "embed"
	"github.com/spf13/viper"
	"history-engine/engine/data"
	"history-engine/engine/utils"
	"log"
	"os"
)

var Common = struct {
	Env            string
	EnableRegister bool
	Lang           string
	DataPath       string
	TimeZone       string
	IconPath       string
	HtmlPath       string
}{
	Env:            "dev",
	EnableRegister: false,
	Lang:           "zh-CN",
	DataPath:       "data",
	TimeZone:       "Asia/Shanghai",
	IconPath:       "data/icon",
	HtmlPath:       "data/html",
}

func loadCommon() {
	v := viper.Sub("common")
	if v != nil {
		if v.IsSet("env") {
			Common.Env = v.GetString("env")
		}
		if v.IsSet("enable_register") {
			Common.EnableRegister = v.GetBool("enable_register")
		}
		if v.IsSet("lang") {
			Common.Lang = v.GetString("lang")
		}
		if v.IsSet("data_path") {
			Common.DataPath = v.GetString("data_path")
		}
		if v.IsSet("time_zone") {
			Common.TimeZone = v.GetString("time_zone")
		}
		if v.IsSet("icon_path") {
			Common.IconPath = v.GetString("icon_path")
		}
		if v.IsSet("html_path") {
			Common.HtmlPath = v.GetString("html_path")
		}
	}

	log.Printf("run env: %s\n", Common.Env)
	log.Printf("common data path: %s\n", Common.DataPath)
	log.Printf("common icon path: %s\n", Common.IconPath)
	checkDataPath()
	checkIconPath()
}

func checkDataPath() {
	if utils.PathExist(Common.DataPath) {
		return
	}

	if err := os.MkdirAll(Common.DataPath, 0755); err != nil {
		log.Fatalf("create data path err:%v\n", err)
	}
}

func checkIconPath() {
	if !utils.PathExist(Common.IconPath) {
		if err := os.MkdirAll(Common.IconPath, 0755); err != nil {
			log.Fatalf("create icon path err:%v\n", err)
		}
	}

	defaultSvg := Common.IconPath + "/default.svg"
	if !utils.FileExist(defaultSvg) {
		if err := os.WriteFile(defaultSvg, data.DefaultSvg, 0644); err != nil {
			log.Fatalf("create default svg icon err:%v\n", err)
		}
	}
}
