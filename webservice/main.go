package main

import (
	"fmt"

	"github.com/deodelacruz/playgo/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Deo",
		LastName:  "Dela Cruz",
	}
	fmt.Println(u)
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(port, err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting web server...")
	// do stuff
	fmt.Println("Server started on port.", port)
	fmt.Println("Number of retries", 2)
	return port, nil

}
