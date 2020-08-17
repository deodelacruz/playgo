/*
Write a function that returns the largest element in a list.
Write function that reverses a list, preferably in place.
Write a function that checks whether an element occurs in a list.
*/

package main

import (
	"fmt"
)

func main() {
	myElements := []int{0, 1, 2, 3, 4}
	largestElem := getLargestElem(myElements)
	fmt.Printf("Largest element is : %v", largestElem)
	myReversedElems := reverseList(myElements)
	fmt.Printf("Reversed element list is : %v", myReversedElems)
	fmt.Println(isElemInList(4, myElements))
}

func getLargestElem(elements []int) int {
	// return largest
	largest := 0
	for _, elem := range elements {
		//fmt.Printf("index %v, element %v \n", i, elem)
		if largest < elem {
			largest = elem
		}
	} // for
	return largest
}

func reverseList(elements []int) []int {
	reversedList := []int{}
	fmt.Println("")
	for i := len(elements) - 1; i >= 0; i-- {
		//fmt.Println("i: ", i, "element", elements[i])
		reversedList = append(reversedList, elements[i])
	}
	return reversedList
}

func isElemInList(myElem int, myList []int) bool {
	isElemInList := false
	for _, elem := range myList {
		if elem == myElem {
			isElemInList = true
		}
	}
	return isElemInList
}
