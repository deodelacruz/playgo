package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Println("count:", count)
	// copy of count passed, orig unchanged after func call
	add5Value(count)
	fmt.Println("count:", count)
	// reference to orig passed, changed after func call
	add5Point(&count)
	fmt.Println("count:", count)
}

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point", *count)
}
