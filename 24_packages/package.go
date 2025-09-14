package main

import (
	"fmt"
	"github.com/fatih/color"

	"github.com/Shashivarunreddy/learning_Go/auth"
	"github.com/Shashivarunreddy/learning_Go/user"
)

func main() {
	// Using the auth package
	auth.Login("admin", "password123")
	session := auth.Session()
	fmt.Println(session)
	user := user.User{
		Username: "shashivarun",
		Email:    "sha@gmail.com",
	}

	fmt.Println(user)
	fmt.Println("User:", user.Username, "Email:", user.Email)

	color.Green("This is a red message")
	
}
