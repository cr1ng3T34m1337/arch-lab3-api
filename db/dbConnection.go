package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
}

func (c *Connection) Open() (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@%s/%s", c.User, c.Password, c.Host, c.DbName)
	return sql.Open("mysql", url)
}

func DbProvider(conn *Connection) (*sql.DB, error) {
	return conn.Open()
}
