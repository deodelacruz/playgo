package main

import (
	"fmt"
)

func main() {
	a, b := 5, 10
	// call swap here
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	var tmpInt int
	tmpInt = *a
	*a = *b
	*b = tmpInt
	fmt.Printf("a: %v , b: %v \n", *a, *b)
}
