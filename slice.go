package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 3)
	for {
		fmt.Println("Please enter an integer to add to the slice. Press 'X' to exit.")
		var userInput string
		if _, err := fmt.Scanln(&userInput); err != nil {
			os.Exit(1)
		} else {
			if userInput == "X" {
				break
			}

			// if input cannot be parsed successfully to int, just ignore, don't add to slice of integers
			parsedInt, _ := strconv.ParseInt(userInput, 0, 64)
			sli = append(sli, int(parsedInt))
			sort.Ints(sli)
			fmt.Printf("Sorted slice: %v\n", sli)
		}
	}
	return
}

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
