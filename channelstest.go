package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 2; i++ {
		readFromChan(i)
	}
	fmt.Println("Done with main().")
}

func readFromChan(threadId int) {
	fmt.Printf("Thread %v\n", threadId)
}
