package models

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sklep.db"), &gorm.Config{})

	if err != nil {
		panic("Doesn't work")
	}

	db.AutoMigrate(&User{}, &Category{}, &Product{}, &Purchase{}, &BoughtProduct{})

	return db
}
