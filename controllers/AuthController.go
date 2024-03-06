package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"server/config"
	"server/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func Hello (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    message:= map[string] string{"Hello": "world"}
    json.NewEncoder(w).Encode(message)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Unmarshal request body
    var user User
    if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Connect to the DB
    config.ConnectToDB()
	db := config.DB 
    // defer db.Close() // Close the database connection at the end

    // Check if user with provided email already exists
    existingUser := model.User{}
    if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
        http.Error(w, "User with this email already exists", http.StatusConflict)
        return
    }

    // Hash the password
    hashedPassword, err := hashPassword(user.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Create new user
    if err := model.CreateUser(user.Username, user.Email, hashedPassword); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{"message": "SUCCESSFULLY SIGNED UP"}
    json.NewEncoder(w).Encode(response)
    w.WriteHeader(http.StatusCreated)
}

func hashPassword(password string)(string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}