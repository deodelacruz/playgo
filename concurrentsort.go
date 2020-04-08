/*

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	getUserInput()

}

func getUserInput() {
	fmt.Println("Please type in a series of integers and I will perform concurrent sorting on those. Type . to end user input.")
	inputSlice := []int{}
	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		myInt, err := strconv.Atoi(strings.TrimSpace(inputScanner.Text()))
		if err != nil {
			fmt.Println("Oops. Please type in an integer.")
			//fmt.Printf("Err: %v\n", err)
		} else {
			fmt.Printf("Input: %i\n", myInt)
			// append to slice
			inputSlice = append(inputSlice, myInt)
			fmt.Printf("inputSlice: %v\n", inputSlice)
		}
	}
	if inputScanner.Err() != nil {
		fmt.Println(inputScanner.Err())
		os.Exit(1)
	}

}
