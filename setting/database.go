package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Database = struct {
		Drive    string
		Host     string
		Port     int
		Name     string
		User     string
		Password string
		Ssl      bool
		Charset  string
		Timeout  int
	}{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
		Ssl:      false,
		Charset:  "utf8mb4",
		Timeout:  500,
	}
)

func loadDatabase() {
	v := viper.Sub("database")
	if v == nil {
		panic("database setting not found")
	}

	Database.Drive = v.GetString("drive")
	Database.Host = v.GetString("host")
	Database.Port = v.GetInt("port")
	Database.Name = v.GetString("name")
	Database.User = v.GetString("user")
	Database.Password = v.GetString("password")
	Database.Ssl = v.GetBool("ssl")
	Database.Charset = v.GetString("charset")
	Database.Timeout = v.GetInt("timeout")
}

func GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%dms&parseTime=true",
		Database.User,
		Database.Password,
		Database.Host,
		Database.Port,
		Database.Name,
		Database.Charset,
		Database.Timeout,
	)
}
