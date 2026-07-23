package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"taskflow/config"
	"taskflow/models"
	"time"
)

// validPriority makes sure only known priority values get stored.
// Anything empty or unrecognized quietly falls back to "medium"
// instead of failing the request outright.
func validPriority(p string) string {
	switch p {
	case "low", "medium", "high":
		return p
	default:
		return "medium"
	}
}

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

	task.Priority = validPriority(task.Priority)

	// A task can technically be created already-completed;
	// if so, stamp CompletedAt now instead of leaving it nil.
	if task.Completed {
		now := time.Now()
		task.CompletedAt = &now
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

	// invalidate cached task list so the next GET reflects the new task
	config.CacheDelete(fmt.Sprintf("tasks:%d", task.UserID))

	json.NewEncoder(w).Encode(task)
}

// get tasks function written below

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	userID := uint(
		r.Context().
			Value("user_id").
			(float64),
	)

	cacheKey := fmt.Sprintf("tasks:%d", userID)

	// try cache first — skip the DB entirely on a hit
	if err := config.CacheGet(cacheKey, &tasks); err == nil {
		json.NewEncoder(w).Encode(tasks)
		return
	}

	err := config.DB.
		Where("user_id = ?", userID).
		Find(&tasks).Error

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	// cache the result for 5 minutes
	config.CacheSet(cacheKey, tasks, 5*time.Minute)

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

	// invalidate cache
	config.CacheDelete(fmt.Sprintf("tasks:%d", userID))

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Task Deleted Successfully",
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
				Value("user_id").(float64))

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

	task.Priority =
		validPriority(updatedTask.Priority)

	task.DueDate =
		updatedTask.DueDate

	// Only touch CompletedAt when the completed status actually
	// changes — this is the key fix that makes streaks reliable.
	// Editing a title/description no longer disturbs it.

	if updatedTask.Completed && !task.Completed {
		now := time.Now()
		task.CompletedAt = &now
	} else if !updatedTask.Completed && task.Completed {
		task.CompletedAt = nil
	}

	task.Completed =
		updatedTask.Completed

	config.DB.Save(&task)

	// invalidate cache
	config.CacheDelete(fmt.Sprintf("tasks:%d", userID))

	json.NewEncoder(w).Encode(task)
}