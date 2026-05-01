package main

import (
	"fmt"
	"strings"
)

type user struct {
	firstName string
	lastName  string
	email     string
	age       int
	password  []string
}

var user1 = user{
	firstName: "Jeff",
	lastName:  "Alex",
	email:     "jeff.alex@example.com",
	age:       21,
	password:  []string{"password123"},
}

var user2 = user{
	firstName: "Tobi",
	lastName:  "Alabi",
	email:     "tobi.alabi@gmail.com",
	age:       21,
	password:  []string{"password001"},
}

var user3 = user{
	firstName: "John",
	lastName:  "Sanjay",
	email:     "john.sanjay@apple.com",
	age:       21,
	password:  []string{"password202"},
}

func main() {

	users := createUsers()
	printUsers(users)
	filtered := filterByDomain(users, "@apple.com")
	fmt.Println("\nApple employees:")
	printUsers(filtered)
}

func createUsers() []user {
	users := []user{user1, user2, user3}
	return users
}

func printUsers(users []user) {
	for _, value := range users {
		fmt.Printf("Name: %s %s, Email: %s, Age: %d\n", value.firstName, value.lastName, value.email, value.age)
	}
}

func filterByDomain(users []user, domain string) []user {
	var filtered []user
	for _, value := range users {
		if strings.Contains(value.email, domain) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}
func printAppleUsers(users []user) {
	for _, value := range users {
		fmt.Printf("user with email %s is an apple employee\n", value.email)
	}
}
