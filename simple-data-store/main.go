package main

import (
	"fmt"
	"strings"
)

func main() {
	type user struct {
		firstName string
		lastName  string
		email     string
		age       int
		password  []string
	}

	user1 := user{
		"Jeff",
		"Alex",
		"jeff.alex@example.com",
		21,
		[]string{"password123"},
	}
	// fmt.Printf("%+v\n", user1)

	user2 := user{
		"Tobi",
		"Alabi",
		"tobi.alabi@gmail.com",
		21,
		[]string{"password001"},
	}
	// fmt.Printf("%+v\n", user2)

	user3 := user{
		"John",
		"Sanjay",
		"john.sanjay@apple.com",
		21,
		[]string{"password202"},
	}
	// fmt.Printf("%+v\n", user3)
	users := []user{user1, user2, user3}

	for index, value := range users {
		fmt.Printf("index: %d, value: %v\n", index, value)
		if strings.Contains(value.email, "@apple.com") {
			fmt.Printf("user with email %s is an apple employee \n", value.email)
		}
	}
}
