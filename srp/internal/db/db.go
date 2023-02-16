package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Create(username, password string) (*sql.DB, error) {
	creds := fmt.Sprintf("%s:%s@/payment", username, password)
	db, err := sql.Open("mysql", creds)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
