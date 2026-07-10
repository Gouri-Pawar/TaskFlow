package models

import "gorm.io/gorm"

type User struct {
	gorm.Model   		// automatically adds id createdAt updatedAt and DeletedAt

	Name string  `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}