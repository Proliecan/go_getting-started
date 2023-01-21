package main

import (
	"fmt"

	"github.com/proliecan/go_getting-started/webservice-tuorial/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Bob",
		LastName:  "Smith",
	}
	fmt.Println(u)
}
