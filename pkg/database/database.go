package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func initDB() *gorm.DB {
	var logLevel = logger.Info
	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	if DB == nil {
		db, err := gorm.Open(sqlite.Open("pokebase.db"), &config)
		if err != nil {
			panic(err)
		}
		DB = db
		return DB
	}
	return DB
}

func Get() *gorm.DB {
	return initDB()
}
