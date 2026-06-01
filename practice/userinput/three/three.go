package main

import "fmt"

func main() {

	//Ask user for age and print if adult
	
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}
}