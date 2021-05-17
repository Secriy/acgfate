package model

import (
	"time"

	"gorm.io/gorm"
)

type Words struct {
	WID       uint64 `gorm:"primaryKey;unique;autoIncrement;comment:'发言ID'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Publisher uint64
	Content   string
}
