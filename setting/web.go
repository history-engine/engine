package setting

import "github.com/spf13/viper"

var (
	Web = struct {
		Port   int
		Addr   string
		Domain string
	}{
		Port:   8080,
		Addr:   "0.0.0.0",
		Domain: "",
	}
)

func loadWeb() {
	v := viper.Sub("web")
	Web.Port = v.GetInt("port")
	Web.Addr = v.GetString("addr")
	Web.Domain = v.GetString("domain")
}
