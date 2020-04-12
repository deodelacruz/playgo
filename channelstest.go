/*
Major problem here is that there are 2 threads that write to the same channel.
When does writing thread close the channel?
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var stringCh chan string

func main() {
	maxReadThreads := 3
	maxWriteThreads := 2
	stringCh = make(chan string, 1)

	// spawn thread that write to channel
	for i := 0; i < maxWriteThreads; i++ {
		wg.Add(1)
		go writeToChan(i)
	}
	//wg.Wait()
	//close(stringCh)

	// spawn multiple threads that read from channel
	for i := 0; i < maxReadThreads; i++ {
		wg.Add(1)
		go readFromChan(i)
	}
	wg.Wait()
	fmt.Println("Done with main().")
}

/*
func provideSessionChannel() chan string { //provide channel back to requestor for comms
	return make(chan str)
} */

func writeToChan(threadId int) {
	sli := []string{"a", "b"}
	for _, elem := range sli {
		elemStr := strconv.Itoa(threadId) + elem
		//fmt.Printf("WriteThread%v attempting channel send: %v\n", threadId, elemStr)

		/* select {
		case stringCh <- elemStr:
			// won't be reached until something is read from ch
			fmt.Printf("WriteThread%v successfully sent to channel: %v\n", threadId, elemStr)
		default:
			fmt.Printf("WriteThread%v was not able to send to channel: %v\n", threadId, elemStr)
		} */

		stringCh <- elemStr
		// won't be reached until something  reads from ch
		fmt.Printf("WriteThread%v successfully sent to channel: %v\n", threadId, elemStr)
	}
	wg.Done()
}

func readFromChan(threadId int) {
	for readTries := 0; readTries < 2; readTries++ {
		time.Sleep(time.Millisecond)
		//fmt.Printf("ReadThread%v attempting to read from channel.\n", threadId)
		select {
		case val := <-stringCh:
			// won't be reached until something is read from ch
			fmt.Printf("ReadThread%v successfully read from channel: %v\n", threadId, val)
			break
		default:
			fmt.Printf("ReadThread%v was not able to read from channel.\n", threadId)
		}
	} //for
	wg.Done()
}
