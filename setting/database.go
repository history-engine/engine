package setting

import (
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"time"
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
		PrintSql bool
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
		PrintSql: false,
	}
)

func loadDatabase() {
	v := viper.Sub("database")
	if v != nil {
		if v.IsSet("drive") {
			Database.Drive = v.GetString("drive")
		}
		if v.IsSet("db_path") {
			Database.Path = v.GetString("db_path")
		}
		if v.IsSet("host") {
			Database.Host = v.GetString("host")
		}
		if v.IsSet("port") {
			Database.Port = v.GetInt("port")
		}
		if v.IsSet("name") {
			Database.Name = v.GetString("name")
		}
		if v.IsSet("user") {
			Database.User = v.GetString("user")
		}
		if v.IsSet("password") {
			Database.Password = v.GetString("password")
		}
		if v.IsSet("ssl") {
			Database.Ssl = v.GetBool("ssl")
		}
		if v.IsSet("charset") {
			Database.Charset = v.GetString("charset")
		}
		if v.IsSet("timeout") {
			Database.Timeout = v.GetDuration("timeout")
		}
		if v.IsSet("print_sql") {
			Database.PrintSql = v.GetBool("print_sql")
		}
	}

	if Database.Drive == dialect.SQLite {
		log.Printf("database drive: %s, path: %s\n", Database.Drive, Database.Path)
	} else {
		log.Printf("database drive: %s, host: %s:%d, name: %s\n", Database.Drive, Database.Host, Database.Port, Database.Name)
	}
}

func GetDSN() string {
	loc := url.QueryEscape(Common.TimeZone)

	switch Database.Drive {
	case dialect.SQLite:
		return fmt.Sprintf(
			"file:%s?mode=rwc&cache=shared&_journal_mode=WAL&_fk=1&loc=%s",
			Database.Path,
			loc,
		)

	case dialect.MySQL:
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true&loc=%s",
			Database.User,
			Database.Password,
			Database.Host,
			Database.Port,
			Database.Name,
			Database.Charset,
			Database.Timeout,
			loc,
		)

	case dialect.Postgres:
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s TimeZone=%s",
			Database.Host,
			Database.Port,
			Database.User,
			Database.Name,
			Database.Password,
			loc,
		)
	}

	return ""
}
