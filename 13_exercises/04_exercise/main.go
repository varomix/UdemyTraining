package main

import "fmt"

func main() {
	var min int
	var max int

	fmt.Println("Please enter a small number:")
	fmt.Scan(&min)
	fmt.Println("Please enter a BIGGER number:")
	fmt.Scan(&max)

	var remainder = max % min

	// fmt.Println("The remainder of", min, "divided by", max, "is", remainder)
	fmt.Printf("The remainder of %d divided by %d is %d\n", min, max, remainder)
}
