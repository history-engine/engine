package db

import (
	"database/sql"
	"database/sql/driver"
	"modernc.org/sqlite"
)

type sqlite3Driver struct {
	*sqlite.Driver
}

type sqlite3DriverConn interface {
	Exec(string, []driver.Value) (driver.Result, error)
}

func (d sqlite3Driver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	_, err = conn.(sqlite3DriverConn).Exec("PRAGMA foreign_keys = ON;", nil)
	if err != nil {
		_ = conn.Close()
	}
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqlite3Driver{Driver: &sqlite.Driver{}})
}
