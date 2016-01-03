package main

import "fmt"

func main() {
	// a note for git test
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)

	}
}
