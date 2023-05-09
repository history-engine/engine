package setting

import (
	"github.com/spf13/viper"
	"os"
)

var (
	CustomFile = "setting.toml"
)

func Load(file string) error {
	if err := initViper(file); err != nil {
		return err
	}

	loadWeb()
	loadSingleFile()
	loadJwt()
	loadDatabase()
	loadZincSearch()

	return nil
}

func initViper(file string) error {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("HE")
	viper.SetConfigType("toml")
	viper.SetConfigFile(file)

	return viper.ReadInConfig()
}
