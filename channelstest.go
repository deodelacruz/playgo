package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var stringCh chan string

func main() {
	maxThreads := 4
	stringCh = make(chan string)
	// spawn thread that writes to channel
	go writeToChan()
	// spwan multiple threads that read from channel
	for i := 0; i < maxThreads; i++ {
		wg.Add(1)
		go readFromChan(i)
	}
	wg.Wait()
	fmt.Println("Done with main().")
}

func writeToChan() {
	sli := []string{"a", "b"}
	for _, elem := range sli {
		stringCh <- elem
	}
	close(stringCh)
}

func readFromChan(threadId int) {
	val := <-stringCh
	fmt.Printf("Thread%v read from channel this value: %s\n", threadId, val)
	wg.Done()
}
