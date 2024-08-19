package setting

import (
	"os"

	"github.com/spf13/viper"
)

func Load(file string) error {
	if err := initViper(file); err != nil {
		return err
	}

	loadCommon()
	loadLogger()
	loadJwt()
	loadWeb()
	loadRedis()
	loadDatabase()
	loadSearch()
	loadMeiliSearch()
	loadZincSearch()
	loadReadability()
	loadSingleFile()

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
