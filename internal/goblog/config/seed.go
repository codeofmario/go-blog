package config

import (
	"goblog.com/goblog/internal/goblog/model"
	"gorm.io/gorm"
)

func SeedDemoData(DB *gorm.DB) {
	// Create john
	john := &model.User{}
	result := DB.First(john, "email = ?", "john@goblog.com")
	if result.RowsAffected == 0 {
		john = &model.User{
			Username: "john",
			Email:    "john@goblog.com",
			Password: "password",
		}

		DB.Create(&john)
	}

	// Create jane
	jane := &model.User{}
	result = DB.First(&jane, "email = ?", "jane@goblog.com")
	if result.RowsAffected == 0 {
		jane = &model.User{
			Username: "jane",
			Email:    "jane@goblog.com",
			Password: "password",
		}

		DB.Create(jane)
	}
}
