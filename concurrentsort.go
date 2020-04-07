/*

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	getUserInput()

}

func getUserInput() {
	fmt.Println("Please type in a series of integers and I will perform concurrent sorting on those. Type . to end user input.")
	inputScanner := bufio.NewScanner(os.Stdin)
	tmp := ""
	for inputScanner.Scan() {
		tmp = inputScanner.Text()
		fmt.Printf("Input: %v\n", tmp)
	}
	if inputScanner.Err() != nil {
		fmt.Println(inputScanner.Err())
		os.Exit(1)
	}

}
