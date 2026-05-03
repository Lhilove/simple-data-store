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
	age:       29,
	password:  []string{"password001"},
}

var user3 = user{
	firstName: "John",
	lastName:  "Sanjay",
	email:     "john.sanjay@apple.com",
	age:       18,
	password:  []string{"password202"},
}

func main() {
	// this is to create users and print them out
	users := createUsers()
	printUsers(users)

	// filter by domain
	filtered := filter(users, func(u user) bool {
		return strings.Contains(u.email, "@apple.com")

	})
	fmt.Println("\nApple employees:")
	printUsers(filtered)

	// filter by age
	filtered = filter(users, func(u user) bool {
		return u.age < 20
	})
	fmt.Println("\nUsers younger than 20:")
	printUsers(filtered)

	// filter by name (findByName already does this ones job, this is just to demostrate firstclass function)
	// filtered = filter(users, func(u user) bool {
	// 	return u.firstName == "Jeff"
	// })
	// fmt.Println("\nUsers with the first name Jeff:")
	// printUsers(filtered)

	// this is to delete user from the store
	updated := deleteUser(users, "jeff.alex@example.com")
	fmt.Println("\nthe users left:")
	printUsers(updated)

	// this is to update user email

	readded := updateUser(users, "jeff.alex@example.com", "jeff.okodua@example.com")
	fmt.Println("\nthe new user has been added:")
	printUsers(readded)

	// this is to find users by their first name or last name
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

// this is a firstclass function that allows you filter by name, domain or age instead of writing every filter function
func filter(users []user, fn func(user) bool) []user {
	var filtered []user
	// using range instead of index loop since we do not need the postion of each slice
	for _, value := range users {
		if fn(value) {
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

// updates the email of an existing user identified by their current email
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

// returns the user if found, returns an error if not found
// caller must check err before using the returned user
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
