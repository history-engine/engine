package setting

import (
	"github.com/spf13/viper"
	"history-engine/engine/utils"
	"log"
)

var (
	Web = struct {
		Port   int
		Addr   string
		Domain string
		UiPath string
	}{
		Port:   8080,
		Addr:   "0.0.0.0",
		Domain: "localhost",
		UiPath: "data/webui",
	}
)

func loadWeb() {
	v := viper.Sub("web")
	if v != nil {
		if v.IsSet("port") {
			Web.Port = v.GetInt("port")
		}
		if v.IsSet("addr") {
			Web.Addr = v.GetString("addr")
		}
		if v.IsSet("domain") {
			Web.Domain = v.GetString("domain")
		}
		if v.IsSet("ui_path") {
			Web.UiPath = v.GetString("ui_path")
		}
	}

	log.Printf("webui path: %s\n", Web.UiPath)
	checkWebUiPath()
}

func checkWebUiPath() {
	if !utils.FileExist(Web.UiPath + "/index.html") {
		log.Fatalf("webui path may empty")
	}
}
