package main

import (
	"fmt"
	us "User/user"
)

type User struct {
	Name       string
	Age        uint
	Occupation string
}

func main() {
	users:=us.USER()
	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("Name : %s\nAge : %d\nOccupation : %s\n\n", user.Name, user.Age, user.Occupation)
	}
}
