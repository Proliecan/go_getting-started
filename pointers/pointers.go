package main

import "fmt"

func main() {
	var firstName *string
	firstName = new(string)
	*firstName = "John"
	fmt.Println(firstName, *firstName)

	lastName := "Doe"
	fmt.Println(lastName)
	fmt.Println(&lastName)

	lastNamePtr := &lastName
	fmt.Println(lastNamePtr, *lastNamePtr)

	lastName = "Smith"
	fmt.Println(lastNamePtr, *lastNamePtr)

	*lastNamePtr = "Anderson"
	fmt.Println(lastNamePtr, *lastNamePtr)
}
