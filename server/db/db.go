package db

import (
	"fcompressor/env"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func DB() *gorm.DB {
	return instance
}

func Open() {
	var (
		host     = env.Get("DB_HOST").String("localhost")
		port     = env.Get("DB_PORT").Int(5432)
		username = env.Get("DB_USERNAME").String("postgres")
		password = env.Get("DB_PASSWORD").String()
		db       = env.Get("DB_NAME").String("postgres")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, db, port)
	_db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("DB ERROR", err)
	}

	instance = _db
}
