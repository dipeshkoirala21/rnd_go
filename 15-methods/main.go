package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	dipesh := User{"dipesh", "dipesh@go.dev", true, 16}
	fmt.Println(dipesh)
	fmt.Printf("dipesh details are: %+v\n", dipesh)
	fmt.Printf("Name is %v and email is %v.\n", dipesh.Name, dipesh.Email)
	dipesh.GetStatus()
	dipesh.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", dipesh.Name, dipesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
