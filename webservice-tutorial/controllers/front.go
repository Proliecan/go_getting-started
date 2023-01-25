package controllers

import (
	"fmt"
	"net/http"
)

func RegisterControllers() {
	uc := NewUserController()

	http.Handle("/users/", *uc)
	fmt.Println("/users/ was registered")
}
