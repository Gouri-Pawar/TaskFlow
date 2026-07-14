package main

import (

	"fmt"
	"net/http"

	"taskflow/config"
	"taskflow/handlers"
	"taskflow/models"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	fmt.Println("Server Running on port 8080")

	http.ListenAndServe(":8080", nil)

	// load env --> connect DB --> db connected --> automigrate (user) --> users table created 

}