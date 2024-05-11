package setting

import "github.com/spf13/viper"

var Common = struct {
	Env            string
	EnableRegister bool
	Lang           string
}{
	Env:            "dev",
	EnableRegister: false,
	Lang:           "zh-CN",
}

func loadCommon() {
	v := viper.Sub("common")
	if v == nil {
		panic("common setting not found")
	}

	if v.IsSet("env") {
		Common.Env = v.GetString("env")
	}
	if v.IsSet("enable_register") {
		Common.EnableRegister = v.GetBool("enable_register")
	}
	if v.IsSet("lang") {
		Common.Lang = v.GetString("lang")
	}
}
