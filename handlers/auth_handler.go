package handlers

import (
	"encoding/json"
	"net/http"
	"taskflow/config"
	"taskflow/models"

	"golang.org/x/crypto/bcrypt"
)

// w - used to send data to client
// r - contains everything sent by the client

func Register(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User  			// creates an empty struct

	// Decode() - converts json --> Go struct  | Decode() needs memory address to modify original struct
 
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	// bycrypt - works on bytes
	// [] bytes(user.Pass) - converts string to bytes
	// bcrypt.DefaultCost - control hashing difficulty 
	// bycrypt takes password and hash it using GenerateFromPassword() function

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte (user.Password), 
		bcrypt.DefaultCost,
	)

	if err != nil {
		http.Error(w, "Error hashing Password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	// Create() - GORM method to insert data

	result := config.DB.Create(&user)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string {
		"message" : "User Registered Successfully",
	}) 

}