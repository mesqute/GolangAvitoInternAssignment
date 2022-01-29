package main

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

// инициализация
func init() {
	conn, err := sql.Open("mysql", "root:0000@/cashdb")
	if err != nil {
		log.Fatal(err)
	}
	db = conn

	// задаем параметры пула соединений БД
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
}

func GetDB() *sql.DB {
	return db
}
