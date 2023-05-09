package setting

import "github.com/spf13/viper"

var (
	ZincSearch = struct {
		Host        string
		Index       string
		IndexPrefix string
		SharedNum   int
		User        string
		Password    string
	}{
		Host:        "http://localhost:8080",
		Index:       "history",
		IndexPrefix: "he",
		SharedNum:   3,
		User:        "",
		Password:    "",
	}
)

func loadZincSearch() {
	v := viper.Sub("zincsearch")
	ZincSearch.Host = v.GetString("host")
	ZincSearch.Index = v.GetString("index")
	ZincSearch.IndexPrefix = v.GetString("index_prefix")
	ZincSearch.SharedNum = v.GetInt("shared_num")
	ZincSearch.User = v.GetString("user")
	ZincSearch.Password = v.GetString("password")
}
