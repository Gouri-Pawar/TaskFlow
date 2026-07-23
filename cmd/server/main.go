package main

import (
	"fmt"
	"net/http"

	"taskflow/config"
	"taskflow/handlers"
	"taskflow/middleware"
	"taskflow/models"
)

// CORS Middleware is added to access data to frontend from backend
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		w.Header().Set(
			"Access-Control-Allow-Origin",
			"*",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		// Browser sends OPTIONS request first
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	http.HandleFunc(
		"/register",
		handlers.Register,
	)

	http.HandleFunc(
		"/login",
		handlers.Login,
	)

	http.HandleFunc(
		"/get_tasks",
		middleware.JWTMiddleware(
			handlers.GetTasks,
		),
	)

	http.HandleFunc(
		"/tasks",
		middleware.JWTMiddleware(
			handlers.CreateTask,
		),
	)

	http.HandleFunc(
		"/tasks/delete",
		middleware.JWTMiddleware(
			handlers.DeleteTask,
		),
	)
	
	http.HandleFunc(
		"/tasks/update",
		middleware.JWTMiddleware(
		handlers.UpdateTask,
		),
	)

	config.ConnectRedis()

	fmt.Println("Server Running on port 8080")

	err := http.ListenAndServe(
		":8080",
		enableCors(http.DefaultServeMux),
	)

	if err != nil {
		fmt.Println(err)
	}
}