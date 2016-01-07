package main

import "fmt"

func main() {
	result := 0
	for i := 1; i < 1000; i++ {
		switch {
		case i%3 == 0 || i%5 == 0:
			result += i
		}
	}

	fmt.Println("Result:", result) // Result: 233168
}
