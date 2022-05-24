package models

import (
	_ "github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique"`
	CreatedAt time.Time      `gorm:"default:autoCreateTime"`
	UpdatedAt time.Time      `gorm:"default:autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
