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

func DeleteTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodDelete {

		http.Error(
			w,
			"Only DELETE allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	taskID := r.URL.Query().Get("id")

	userID := uint(
			r.Context().
				Value("user_id").(float64),
		)

	var task models.Task

	err := config.DB.Where(
			"id = ? AND user_id = ?",
			taskID,
			userID,
		).
		First(&task).Error

	if err != nil {

		http.Error(
			w,
			"Task not found",
			http.StatusNotFound,
		)
		return
	}

	config.DB.Delete(&task)

	json.NewEncoder(w).Encode(
		map[string]string{
			"message":
				"Task Deleted Successfully",
		},
	)
}

func UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPut {

		http.Error(
			w,
			"Only PUT allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	taskID :=
		r.URL.Query().Get("id")

	userID :=
		uint(
			r.Context().
				Value("user_id").(float64),)

	var task models.Task

	err := config.DB.Where(
			"id = ? AND user_id = ?",
			taskID,
			userID,
		).
		First(&task).Error

	if err != nil {

		http.Error(
			w,
			"Task not found",
			http.StatusNotFound,
		)
		return
	}

	var updatedTask models.Task

	err = json.NewDecoder(
			r.Body,
		).Decode(
			&updatedTask)

	if err != nil {

		http.Error(
			w,
			"Invalid JSON",
			http.StatusBadRequest,
		)
		return
	}

	task.Title =
		updatedTask.Title

	task.Description =
		updatedTask.Description

	task.Completed =
		updatedTask.Completed

	config.DB.Save(&task)

	json.NewEncoder(w).Encode(task)
}