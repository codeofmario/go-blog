package model

import "github.com/google/uuid"

type Post struct {
	Model
	Title   string `gorm:"type:VARCHAR"`
	Body    string `gorm:"type:TEXT"`
	ImageID uuid.UUID
	UserID  uuid.UUID

	Comments []Comment
}
