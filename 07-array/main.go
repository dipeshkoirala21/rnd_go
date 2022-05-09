package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Dog"
	fruitList[3] = "Tomato"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))

	var vegList = [6]string{"potato", "beans", "mushrooms"}
	fmt.Println("Vegetables list is: ", vegList)
	fmt.Println("Vegetables list is: ", len(vegList))

}
