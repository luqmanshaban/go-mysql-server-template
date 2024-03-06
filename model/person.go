package model

import (
	"server/config"
)


type User struct {
	ID       uint  
	Username string 
	Email string 
	Password string 
}

func CreateUser(username, email, password string) error{
	db, err := config.ConnectToDB()
	if err != nil {
		return err
	}

	newUser := User{
		Username: username,
		Email: email,
		Password: password,
	}

	result := db.Create(&newUser)

	if result.Error != nil {
		println("Error creating new user: ", result.Error)
		return result.Error
	}

	println("New User created: ", newUser.Username)
	return nil
}

func GetAllUsers() ([]User, error) {
    var users []User
    db := config.DB // Access the global DB variable from the config package
    
    // Retrieve all users from the database
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }

    return users, nil
}