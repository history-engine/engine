package setting

import "github.com/spf13/viper"

var Search = struct {
	Engine string
	Prefix string
}{
	Engine: "meili",
	Prefix: "history_engine_index",
}

func loadSearch() {
	v := viper.Sub("search")
	if v != nil {
		if v.IsSet("search_engine") {
			Search.Engine = v.GetString("engine")
		}
		if v.IsSet("prefix") {
			Search.Prefix = v.GetString("prefix")
		}
	}
}
