package main

import "fmt"

func main() {

	// Swap two variables with temporary variable

	var a int = 5
	var b int = 10

	temp := a
	a = b
	b = temp
	
	fmt.Printf("a: %d, b: %d\n", a, b)

}

