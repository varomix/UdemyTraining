package main

import "fmt"

func half(x int) (int, bool) {
	even := x%2 == 0
	return x / 2, even
}

func main() {

	fmt.Println(half(5))
}
