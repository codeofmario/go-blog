package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *Model) BeforeCreate(tx *gorm.DB) (err error) {
	entity.ID = uuid.New()
	return
}
