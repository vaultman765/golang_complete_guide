package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // Embedded type
}

// Optional pointer
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// Must be a pointer receiver to modify the struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Utility function to create a new user
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// Utility function to create a new admin
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}
