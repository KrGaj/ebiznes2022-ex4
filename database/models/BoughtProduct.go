package models

import "gorm.io/gorm"

type BoughtProduct struct {
	gorm.Model

	product Product
	amount  int
}
