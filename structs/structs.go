package main

import (
	"fmt"
	"go-sandbox/structs/custom"
	"go-sandbox/structs/user"
)

func main() {
	// this
	if false {
		userFirstName := getUserData("Please enter your first name: ")
		userLastName := getUserData("Please enter your last name: ")
		userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

		var appUser *user.User
		appUser, err := user.New(userFirstName, userLastName, userBirthdate)

		if err != nil {
			fmt.Println(err)
			return
		}

		admin := user.NewAdmin("test@example.com", "123456")

		// admin.User.OutputUserDetails()
		// admin.User.ClearUserName()
		// admin.User.OutputUserDetails()

		admin.OutputUserDetails()
		admin.ClearUserName()
		admin.OutputUserDetails()

		appUser.OutputUserDetails()
		appUser.ClearUserName()
		appUser.OutputUserDetails()
	}

	// custom
	if false {
		custom.Run()
	}
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
