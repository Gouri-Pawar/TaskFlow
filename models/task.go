package models

import (
	"gorm.io/gorm"
	"time"
) 

type Task struct {
	gorm.Model

	Title       			string              `json:"title"`
	Description  	 string              `json:"description"`
	Completed       bool                `json:"completed"`
	Priority  		     string               `json:"priority"`
	DueDate		     *time.Time		`json:"due_date"`
	CompletedAt	 *time.Time 	`json:"completed_at"`
	UserID               uint                 `json:"user_id"`
}