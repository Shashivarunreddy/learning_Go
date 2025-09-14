package auth

import "fmt"

func Login(username, password string){
	fmt.Println("Authenticating user:", username)
	fmt.Println("Password:", password)
}

