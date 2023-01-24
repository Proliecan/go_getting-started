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

	port := 3000
	_, err := startWebserver(port, 2)
	fmt.Println(err)
}

// returns success or not (error) AND the port (int)
func startWebserver(port, numberOfRetries int) (int, error) { // <name> <type>, ... or <name>, <name>, ... <type> if same type.
	fmt.Println("Starting server...")
	// do smth
	fmt.Println("Server running on port", port)
	fmt.Println("Number of retries:", numberOfRetries)
	// return errors.New("Something went wrong when starting the server (test)")
	return port, nil
}
