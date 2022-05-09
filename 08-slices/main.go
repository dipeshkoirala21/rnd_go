package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to exercises on slices")
	var fruitList = []string{"Apple", "Banana", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 567
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

}