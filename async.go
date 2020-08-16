package main

import (
	"fmt"
	"time"
)

func myAsyncFunction(s int32) <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		fmt.Println("Doing something.")
		time.Sleep(time.Second * 2)
		r <- s
	}()
	return r
}

func main() {
	var r int32
	select {
	case r = <-myAsyncFunction(2):
	case r = <-myAsyncFunction(3):
	}
	// outputs `2` after three seconds
	fmt.Println(r)
}
