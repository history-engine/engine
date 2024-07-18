package setting

import (
	"entgo.io/ent/dialect"
	"fmt"
	"history-engine/engine/utils"
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	Database = struct {
		Drive    string
		Path     string
		Host     string
		Port     int
		Name     string
		User     string
		Password string
		Ssl      bool
		Charset  string
		Timeout  time.Duration
	}{
		Drive:    "sqlite",
		Path:     "data/sqlite.db",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
		Ssl:      false,
		Charset:  "utf8mb4",
		Timeout:  500 * time.Millisecond,
	}
)

func loadDatabase() {
	v := viper.Sub("database")
	if v != nil {
		if v.IsSet("drive") {
			Database.Drive = v.GetString("drive")
		}
		Database.Host = v.GetString("host")
		Database.Port = v.GetInt("port")
		Database.Name = v.GetString("name")
		Database.User = v.GetString("user")
		Database.Password = v.GetString("password")
		Database.Ssl = v.GetBool("ssl")
		Database.Charset = v.GetString("charset")
		Database.Timeout = v.GetDuration("timeout")
	}

	if Database.Drive == "sqlite" {
		log.Printf("database drive: %s, path: %s\n", Database.Drive, Database.Path)
	} else {
		log.Printf("database drive: %s, host: %s:%d, name: %s\n", Database.Drive, Database.Host, Database.Port, Database.Name)
	}

	checkSqliteFile()
}

func checkSqliteFile() {
	if Database.Drive != "sqlite" {
		return
	}

	if utils.FileExist(Database.Path) {
		return
	}

	// todo create empty db file
}

func GetDSN() string {
	switch Database.Drive {
	case dialect.SQLite:
		return fmt.Sprintf(
			"file:%s?mode=rwc&cache=shared&_journal_mode=WAL",
			Database.Path,
		)

	case dialect.MySQL:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true",
			Database.User,
			Database.Password,
			Database.Host,
			Database.Port,
			Database.Name,
			Database.Charset,
			Database.Timeout,
		)

	case dialect.Postgres:
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s",
			Database.Host,
			Database.Port,
			Database.User,
			Database.Name,
			Database.Password,
		)
	}

	return ""
}
