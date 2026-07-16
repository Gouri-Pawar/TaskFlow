package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title 				    string  `json:"title"`
	Description 	 string `json: "description"`
	Completed		bool `json: "completed"`

	userID uint `json: "user_id"`
}