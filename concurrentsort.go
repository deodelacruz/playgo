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
		inputString := strings.TrimSpace(inputScanner.Text())
		if inputString == "." {
			break
		}
		myInt, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("Oops. Please type in an integer.")
			//fmt.Printf("Err: %v\n", err)
		} else {
			// append to slice
			inputSlice = append(inputSlice, myInt)
			fmt.Printf("inputSlice: %v\n", inputSlice)
		}
	}
	if inputScanner.Err() != nil {
		fmt.Println(inputScanner.Err())
		os.Exit(1)
	}
	fmt.Printf("inputSlice: %v\n", inputSlice)
	fmt.Printf("inputSlicePtr: %v\n", &inputSlice)
	processInputSlice(inputSlice)
}

func processInputSlice(inputSlice []int) {
	fmt.Printf("inputSlicePtr: %v\n", &inputSlice)
	fmt.Printf("inputSlice: %v\n", inputSlice)
	fmt.Printf("inputSliceLength: %v\n", len(inputSlice))
	// divide slice into 4 equal parts
	slice1 := []int{}
	slice2 := []int{}
	slice3 := []int{}
	slice4 := []int{}
	for i, elem := range inputSlice {
		remainder := i % 4 // modulo or remainder
		fmt.Printf("remainder of div 4: %v\n", remainder)
		fmt.Printf("elem: %v\n", elem)
		if remainder == 0 {
			slice1 = append(slice1, elem)
		}
		if remainder == 1 {
			slice2 = append(slice2, elem)
		}
		if remainder == 2 {
			slice3 = append(slice3, elem)
		}
		if remainder == 3 {
			slice4 = append(slice4, elem)
		}
	}
	fmt.Printf("slices: %v %v %v %v\n", slice1, slice2, slice3, slice4)
}
