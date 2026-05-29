package main

import "fmt"

func main() {
	// Ask for name

	name := ""

	fmt.Print("Please enter your name: ")
	fmt.Scanln(&name)

	fmt.Printf("Hello, %s!", name)
}