package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	
	
	for i  := 1; i <= 5; i++ {
		fmt.Print(i)
	}
	
	fmt.Println()

	for	 i := 0; i < 5; i++ {
		fmt.Print(i+1)
	}
}