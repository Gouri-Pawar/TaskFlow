package main

import (
	"taskflow/config"
	"taskflow/models"
)

func main() {

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	// load env --> connect DB --> db connected --> automigrate (user) --> users table created 

}