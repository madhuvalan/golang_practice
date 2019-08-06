package main

import (
	"fmt"
	
	)
func main() {
	userInterfaceScreen1()

	}

func userInterfaceScreen1() {
	var username,password string
	fmt.Println("New User Registration, press 1:")
	fmt.Println("Enter your username:")
	fmt.Scanln(&username)
	if username != "1" {
		fmt.Println("Enter your password")
		fmt.Scanln(&password)
		fmt.Printf("username %s", username)
		fmt.Printf("password %s", password)
		return
	}
	new_user_reg()
}

func usernamePasswordProcess() {


}

func new_user_reg() {
   var username,password string	
   fmt.Println("Came to new user reg screen")
   fmt.Println("Enter your username:")
   fmt.Scanln(&username)
   fmt.Println("Enter your password")
   fmt.Scanln(&password)

}

func plan() {
	fmt.Println("Program to create User login")
	fmt.Println("Step1: Create a user login screen - Monday")
	fmt.Println("Step2 Program a new user scenario - Monday")
	fmt.Println("Step3 Exisitng user scenario - Tuesday")
	fmt.Println("Step4 Test for concurrency - Tuesday/Wednesday")
}