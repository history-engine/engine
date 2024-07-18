package setting

import (
	"github.com/spf13/viper"
	"history-engine/engine/utils"
	"log"
	"os"
)

var Common = struct {
	Env            string
	EnableRegister bool
	Lang           string
	DataPath       string
}{
	Env:            "dev",
	EnableRegister: false,
	Lang:           "zh-CN",
	DataPath:       "data",
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
	}

	log.Printf("run env: %s\n", Common.Env)
	log.Printf("common data path: %s\n", Common.DataPath)
	checkDataPath()
}

func checkDataPath() {
	if utils.PathExist(Common.DataPath) {
		return
	}

	if err := os.MkdirAll(Common.DataPath, 0755); err != nil {
		log.Fatalf("create data path err:%v\n", err)
	}
}
