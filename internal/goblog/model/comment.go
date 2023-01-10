package model

import "github.com/google/uuid"

type Comment struct {
	Model
	Body   string `gorm:"type:TEXT"`
	UserID uuid.UUID
	PostID uuid.UUID
}
