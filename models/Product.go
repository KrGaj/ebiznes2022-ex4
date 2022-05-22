package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	category Category

	name  string
	price int
}
