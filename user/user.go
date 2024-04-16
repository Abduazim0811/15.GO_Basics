package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name       string
	Age        uint
	Occupation string
}

func USER() (users []User){
	file, err := os.Open("../sample.txt")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rs := strings.Split(line, ":")
		switch strings.TrimSpace(rs[0]) {
		case "Name":
			users = append(users, User{Name: strings.TrimSpace(rs[1])})
		case "Age":
			var age uint
			fmt.Sscanf(strings.TrimSpace(rs[1]), "%d", &age)
			users[len(users)-1].Age = uint(age)
		case "Occupation":
			users[len(users)-1].Occupation = strings.TrimSpace(rs[1])
		}
	}
	return users
}
