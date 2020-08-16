/*
Print a random number between 1 and 5 of stars (*) to the console
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// use stdlib to generate random number
	// make the range be between 1 and 5
	//seed rand to avoid deterministic behavior
	rand.Seed(time.Now().UnixNano())
	myRandomInt := rand.Intn(5) + 1
	// print it out

	numStars := strings.Repeat("*", myRandomInt)
	fmt.Println("My random  integer between 1 and 5 is: ", numStars)

}
