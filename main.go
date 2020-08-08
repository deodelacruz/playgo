package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter First Floating Point: ")
	var first float64
	fmt.Scanln(&first)
	fmt.Println(int64(first))

	fmt.Print("Enter Second Floating Point: ")
	var second float64
	fmt.Scanln(&second)
	fmt.Println(int64(second))
}
