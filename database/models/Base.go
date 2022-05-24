package models

import (
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.NewV4()

	if err != nil {
		return err
	}

	return nil
}
