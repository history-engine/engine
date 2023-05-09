package setting

import "github.com/spf13/viper"

var (
	Readability = struct {
		Parser string
		Path   string
	}{
		Parser: "mozilla",
		Path:   "",
	}
)

func loadReadability() {
	v := viper.Sub("readability")
	Readability.Parser = v.GetString("parser")
	Readability.Path = v.GetString("path")
}
