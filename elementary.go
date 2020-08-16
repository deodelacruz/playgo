/*
Write a program that prints ‘Hello World’ to the screen.
Write a program that asks the user for their name and greets them with their name.
Modify the previous program such that only the users Alice and Bob are greeted with their names.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World.")
	//scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hi, what is your name?")
	//var input string
	//for scanner.Scan() {
	//	input = scanner.Text()
	//}

	allowedNames := []string{"Alice", "Bob"}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)

	for _, allowedName := range allowedNames {
		//fmt.Println("allowedName", allowedName)
		if name == allowedName {
			fmt.Println("Hello, ", input)
		} else {
			//fmt.Println("Sorry, I dont' know you", name)
		}
	}
}
