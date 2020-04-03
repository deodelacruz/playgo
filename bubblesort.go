/**

Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// request up to 10 integers from the user
	intSlice := getUserInput()
	fmt.Println("Sorting your integer set via bubble sort algorithm.")
	BubbleSort(intSlice)
	fmt.Printf("sorted slice: %v %s", intSlice, "\n")
}

func getUserInput() []int {
	fmt.Println("Hi. Please type in up to 10 integers (press Enter after each integer and . to end input).")
	reader := bufio.NewReader(os.Stdin)
	inputSlice := []int{}
	for i := 0; i < 10; i++ {
		inputString, _ := reader.ReadString('\n')
		if strings.TrimSpace(inputString) != "." {
			inputInt, err := strconv.Atoi(strings.TrimSpace(inputString))
			if err != nil {
				log.Fatal(err)
			}
			inputSlice = append(inputSlice, inputInt)
		} else {
			break
		}
	}
	fmt.Printf("input: %v %s", inputSlice, "\n")
	return inputSlice
}

// BubbleSort -  sorts a slice of integers via the bubble sort algorith
func BubbleSort(intSlice []int) {
	j := 0
	for {
		swapOccurred := false
		for i := 0; i < (len(intSlice) - 1); i++ {
			if intSlice[i] > intSlice[i+1] { //swap
				Swap(intSlice, i)
				swapOccurred = true
			}
		} // single pass
		fmt.Printf("pass %v: %v %s", j, intSlice, "\n")
		if swapOccurred == false {
			break
		}
		j++
	}
	//fmt.Printf("sorted slice: %v %s", intSlice, "\n")
	//return
}

// Swap - swaps two adjacent slice elements starting at index i
func Swap(intSlice []int, i int) {
	tmpElem := intSlice[i]
	intSlice[i] = intSlice[i+1]
	intSlice[i+1] = tmpElem
}
