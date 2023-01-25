package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/proliecan/go_getting-started/webservice-tuorial/controllers"
)

func main() {
	port := 3000
	fmt.Println("Starting server on port", port)
	controllers.RegisterControllers()
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
