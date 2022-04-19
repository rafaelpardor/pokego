package pokemon

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Id     uint
	Name   string
	Height int
}

func MigrateModel(db *gorm.DB) error {
	return db.AutoMigrate(&Pokemon{})
}
