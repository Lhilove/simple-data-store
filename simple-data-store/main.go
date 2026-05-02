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

	updated := deleteUser(users, "jeff.alex@example.com")
	fmt.Println("\nthe users left:")
	printUsers(updated)

	readded := updateUser(users, "jeff.alex@example.com", "jeff.okodua@example.com")
	fmt.Println("\nthe new user has been added:")
	printUsers(readded)

	names := []string{"Jeff", "Tobi", "John", "Ajala"}
	for _, name := range names {
		foundUser, err := findByName(users, name)
		if err != nil {
			fmt.Printf("\nUser with name %s not found\n", name)
		} else {
			fmt.Printf("\nUser with name %s found: %s %s, Email: %s, Age: %d\n", name, foundUser.firstName, foundUser.lastName, foundUser.email, foundUser.age)
		}
	}

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

// func printAppleUsers(users []user) {
// 	for _, value := range users {
// 		fmt.Printf("user with email %s is an apple employee\n", value.email)
// 	}
// }

func deleteUser(users []user, email string) []user {
	var updated []user
	for _, u := range users {
		if u.email != email {
			updated = append(updated, u)
		}
	}
	return updated
}

func updateUser(users []user, email string, newEmail string) []user {
	var readded []user
	for _, x := range users {
		if x.email == email {
			x.email = newEmail
		}
		readded = append(readded, x)
	}
	return readded
}

func findByName(users []user, name string) (user, error) {
	var names []user
	for _, value := range users {
		if strings.Contains(value.firstName, name) || strings.Contains(value.lastName, name) {
			names = append(names, value)
		}
	}
	if len(names) == 0 {
		return user{}, fmt.Errorf("user not found")
	}
	return names[0], nil
}
