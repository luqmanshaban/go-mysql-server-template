package config

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var DB *gorm.DB // Declare DB variable as a global variable to be used throughout the package

func ConnectToDB() (*gorm.DB, error) {
    dsn := "root:password@tcp(127.0.0.1:3306)/practice?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err // Return error if connection fails
    }
    
    DB = db // Assign the opened DB to the global variable
	println("DATABASE CONNECTED")
    return db, nil // Return DB connection and nil error if successful
}
