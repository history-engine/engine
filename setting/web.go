package setting

import (
	"github.com/spf13/viper"
)

var (
	Web = struct {
		Port   int
		Addr   string
		Domain string
	}{
		Port:   8080,
		Addr:   "0.0.0.0",
		Domain: "http://localhost",
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
	}
}
