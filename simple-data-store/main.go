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

// this variable is for findByName, deleteUser
var name string
var email string

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
	fmt.Println("Enter email to delete: ")
	fmt.Scan(&email)
	record, err := deleteUser(users, email)
	if err != nil {
		fmt.Printf("Invalid email format\n")
	} else {
		fmt.Printf("User with email %s deleted: %s %s, Email: %s, Age: %d\n", email, record.firstName, record.lastName, record.email, record.age)
	}

	// this is to update user email

	readded := updateUser(users, "jeff.alex@example.com", "jeff.okodua@example.com")
	fmt.Println("\nthe new user has been added:")
	printUsers(readded)

	// this is to demonstrate user input for name search
	fmt.Println("Enter name to search: ")
	fmt.Scan(&name)
	result, err := findByName(users, name)
	if err != nil {
		fmt.Printf("User with name %s not found\n", name)
	} else {
		fmt.Printf("User with name %s found: %s %s, Email: %s, Age: %d\n", name, result.firstName, result.lastName, result.email, result.age)
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

// this is to delete user from the store, it returns the deleted user if found, otherwise it returns an error
func deleteUser(users []user, email string) (user, error) {

	//user input validation to prevent auth bypass via email manipulation.
	isValidEmail := strings.Contains(email, "@")
	if !isValidEmail {
		return user{}, fmt.Errorf("invalid email format")
	}
	for _, u := range users {
		if u.email != email {
		} else {
			return u, nil // return the deleted user
		}
	}
	return user{}, fmt.Errorf("user not found")
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

// added pointer receiver to updateUser function to modify the original user in the slice instead of creating a new slice with updated email
func updateUserPointer(users []user, email string, newEmail string) {
	for i := range users {
		if users[i].email == email {
			users[i].email = newEmail
		}
	}
}

// returns the user if found, returns an error if not found
// caller must check err before using the returned user
func findByName(users []user, name string) (user, error) {
	var names []user
	for _, value := range users {
		// fix case insensitive name search, prevents auth bypass via case manipulation
		if strings.EqualFold(value.firstName, name) || strings.EqualFold(value.lastName, name) {
			names = append(names, value)
		}
	}

	//user input validation to verify the inputed name is not empty
	if len(names) == 0 {
		return user{}, fmt.Errorf("user not found")
	}
	return names[0], nil // return the first user found, you can modify this to return all users found if needed
}
