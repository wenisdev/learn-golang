package main

import "fmt"

func main() {
	// Convert int to float

	var intValue int = 42
	var intToFloat float64 = float64(intValue)

	fmt.Printf("The integer value is: %d\n", intValue)
	fmt.Printf("The float value is: %f\n", intToFloat)
}