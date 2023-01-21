package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "John"
	u.LastName = "Doe"
	fmt.Println(u)
	fmt.Println(u.FirstName)

	u2 := user{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Anderson",
	}
	fmt.Println(u2)

	u3 := user{3, "Bill", "Smith"}
	fmt.Println(u3)
}
