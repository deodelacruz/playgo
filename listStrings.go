/*
Write a function that returns the largest element in a list.
Write function that reverses a list, preferably in place.
Write a function that checks whether an element occurs in a list.
Write a function that returns the elements on odd positions in a list.

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
	fmt.Println("Odd elements:", getElementsInOddPositions(myElements))
	//fmt.Println(getDigitsOfNumber(9876))
	addToSlice(&myElements)
	fmt.Println(myElements)
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

func getElementsInOddPositions(elements []int) []int {
	// even if modulus of division by 2 is zero
	oddElems := []int{}
	for i, elem := range elements {
		if i%2 == 0 { //index is even
			// skip
		} else { //index is odd
			oddElems = append(oddElems, elem)
		}
	}
	return oddElems
}

/* Write a function that takes a number and returns a list of its digits. So for 2342 it should return [2,3,4,2].
func getDigitsOfNumber(myNumber int) []int {
	nunmberAsString := strconv.Itoa(myNumber)
	digitsAsString := strings.Split(nunmberAsString)
	digits := []int{}
	for _, strDigits := range digitsAsString {
		digits = append(digits, strconv.AToi(strDigits))
	}
	return digits
} */

// pass slice by reference
func addToSlice(mySlicePtr *[]int) {
	tmpSlice := []int{8, 9}
	*mySlicePtr = append(*mySlicePtr, tmpSlice...)

}
