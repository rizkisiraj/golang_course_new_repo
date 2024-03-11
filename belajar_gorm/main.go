package main

import (
	"fmt"
	"gorm-sql/database"
	"gorm-sql/models"
)

func main() {
	database.StartDB()

	createUser("sijar123@gmail.com")
}

func createUser(email string) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: database connection is nil")
		return
	}

	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("Error create user data", err)
		return
	}

	fmt.Println("New user Data", user)
}
