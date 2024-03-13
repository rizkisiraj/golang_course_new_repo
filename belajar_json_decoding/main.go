package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name": "Fitri Ayu",
		"email": "fitri@gmail.com",
		"age": 23
	}
	`

	var result Student
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full name: ", result.FullName)
	fmt.Println("Email: ", result.Email)
	fmt.Println("Age: ", result.Age)
}
