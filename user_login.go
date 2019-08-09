package main
import (
	"fmt"
	"os"
	"os/exec"
	)

func main() {
	//username,password := userInterfaceScreen1()
	var username = "test"
	var password = "test"
	usernamePasswordProcess(username,password)
	//fmt.Printf("username %s", username)
	//fmt.Printf("password %s", password)
}
func clrScr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func userInterfaceScreen1() (username,password string){
	clrScr()
	fmt.Println("For new User registration, Press key 1:")
	fmt.Println("Enter your username:")
	fmt.Scanln(&username)
	if username != "1" {
		fmt.Println("Enter your password")
		fmt.Scanln(&password)
		return
	}
	username,password = new_user_reg()
	return
}

func usernamePasswordProcess(uname,pword string) {
    fmt.Printf("username %s", uname)
	fmt.Printf("password %s", pword)
}

func new_user_reg() (username,password string){
   clrScr()
   fmt.Println("Came to new user reg screen")
   fmt.Println("Enter your username:")
   fmt.Scanln(&username)
   fmt.Println("Enter your password")
   fmt.Scanln(&password)
   return 

}

func plan() {
	fmt.Println("Program to create User login")
	fmt.Println("Step1: Create a user login screen - Monday")
	fmt.Println("Step2 Program a new user scenario - Monday")
	fmt.Println("Step3 Exisitng user scenario - Tuesday")
	fmt.Println("Step4 Test for concurrency - Tuesday/Wednesday")
}