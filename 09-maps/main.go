package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages ", languages)
	fmt.Println("List of all languages ", languages["JS"])

	//deleting a value

	delete(languages, "RB")
	fmt.Println("List of all languages ", languages)

	//loops
	for _, value := range languages {
		fmt.Printf("For key v, Value is %v\n", value)
	}

}
