package main

import (

	"fmt"
	"net/http"

	"taskflow/config"
	"taskflow/handlers"
	"taskflow/models"
	"taskflow/middleware"
)

func main() {

	config.ConnectDB()

	//creates task table
	config.DB.AutoMigrate(
		&models.User{},
		&models.Task{},
	)

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc(
	"/get_tasks",
	middleware.JWTMiddleware(
		handlers.GetTasks,
	),
)
	http.HandleFunc(
		"/tasks",
		middleware.JWTMiddleware(handlers.CreateTask,
		),
	)


	fmt.Println("Server Running on port 8080")

	http.ListenAndServe(":8080", nil)

	// load env --> connect DB --> db connected --> automigrate (user) --> users table created 

}