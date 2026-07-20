package handlers

import (
	"encoding/json"
	"net/http"

	"taskflow/config"
	"taskflow/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Only POST allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)
		return
	}

	// Get user_id from JWT middleware
	userID := r.Context().Value("user_id")

	task.UserID = uint(userID.(float64))

	err = config.DB.Create(&task).Error

	if err != nil {
		http.Error(
			w,
			"Failed to create task",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// get tasks function written below

func GetTasks(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task

	userID := uint(
		r.Context().
		Value("user_id").
		(float64),
	)

	err := config.DB.
	Where("user_id = ?", userID).
	Find(&tasks).Error

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}