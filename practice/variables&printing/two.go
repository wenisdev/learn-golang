package main

import "fmt"

func two() {

	// Store your name and age then print them

	fmt.Print("Enter your First Name: ")
	var firstname string
	fmt.Scanln(&firstname)

	fmt.Print("Enter your Last Name: ")
	var lastname string
	fmt.Scanln(&lastname)

	fmt.Printf("Your full name is: %s %s\n", firstname, lastname)

}