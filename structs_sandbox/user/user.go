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
	User
}

func NewAdmin(email, password string) (*Admin, error) {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}, nil
}

// using func (x) y, attaches the function to the struct.
func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// a pointer is needed if function wants to mutate values of struct value otherwise a copy is made
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate must be provided")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
