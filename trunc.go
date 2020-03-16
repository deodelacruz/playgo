package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//prompt user to input floating number
	fmt.Println("Hi. Please input a floating point number, and I'll give you back the integer value.")

	var userInput string
	if _, err := fmt.Scanln(&userInput); err != nil {
		os.Exit(1)
	} else {
		//convert string type user input to float64
		floatInput, _ := strconv.ParseFloat(userInput, 64)
		//return the integer value of input, truncating the decimal portion
		fmt.Printf("The integer value of floating point number is: %v\n", trunc(floatInput))
	}

}

func trunc(input float64) int {
	return int(input)
}
