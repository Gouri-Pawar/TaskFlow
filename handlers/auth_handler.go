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

func Login(w http.ResponseWriter, r *http.Request) {

	// Check request method
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// temp struct only needed email and pass that's why User struct not used
	var loginData struct {
		Email	string `json:"email"`
		Password	string `json:"password"`
	}

	// decode json
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//empty user 
	var user models.User

	/*
	 find user by email  ----> 
	SELECT * 
	FROM users 
	WHERE email='gouri@gmail.com'
	LIMIT 1; 
	*/
	result := config.DB.Where("email = ?", loginData.Email).First(&user)

	if result.Error != nil {
		http.Error(w, "User Not Found", http.StatusUnauthorized)
		return
	}

	//Password Comparison  ---> bcrypt internally does hash of entered pass again and comapre
	err = bcrypt.CompareHashAndPassword (
		[]byte(user.Password),
		[]byte(loginData.Password),
	)

	if err != nil {
		http.Error(w,
			"Invalid Password",
			http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Login Successful",
		},
	)

}