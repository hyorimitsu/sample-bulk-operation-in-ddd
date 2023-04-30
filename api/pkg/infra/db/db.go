package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConn(dataSourceName string) (*sql.DB, error) {
	return sql.Open("mysql", dataSourceName)
}

func NewDB(conn *sql.DB) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}))
}
