package main

import (
	"fmt"
	"structs_sandbox/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, _ = user.New(userFirstName, userLastName, userBirthdate)

	//... do something awesome with that gathered data!
	//outputUserData1(&appUser)
	fmt.Println("*****")
	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

	appUser2, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser2.OutputUserData()

	admin, _ := user.NewAdmin("test@email.com", "blah")
	admin.OutputUserData()

}

//func outputUserData1(u *user) {
//	// if using pointers with structs, you don't need to dereference, go has some syntactic sugar at last!
//	fmt.Println(u.firstName, u.lastName, u.birthDate)
//}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
