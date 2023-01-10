package model

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/util"
	"gorm.io/gorm"
)

type User struct {
	Model
	Email    string `gorm:"type:VARCHAR"`
	Username string `gorm:"type:VARCHAR"`
	Password string `gorm:"type:VARCHAR"`

	Posts []Post
}

func (entity *User) BeforeCreate(tx *gorm.DB) (err error) {
	entity.ID = uuid.New()
	entity.Password, _ = util.HashPassword(entity.Password)
	return
}
