package app

import (
	"belajar-golang-rest/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_rest")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
