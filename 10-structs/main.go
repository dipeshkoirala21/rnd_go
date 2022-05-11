package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	dipesh := User{"Dipesh", "dipes@yopmail.com", true, 22}
	fmt.Println(dipesh)
	fmt.Println("Name is %v and email is %v .", dipesh.Name, dipesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
