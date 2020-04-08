/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

My explanation of a race condition:
-----------------------------------
In general, the order of execution and completion of concurrent/interleaved tasks (such as in the go runtime) is non-deterministic. If the outcome of the program 'varies' depending on the order that concurrent tasks are executed (interleaving), then that is a race condition. Race conditions can occur due to inter-task communication and co-dependence, such as when multiple goroutines modify the same data in a shared variable.

*/

package playgo

import (
	"fmt"
	"time"
)

func main() {
	var isChildMalePtr *bool
	isChildMale := true
	isChildMalePtr = &isChildMale

	go momWantsFemale(isChildMalePtr)
	go dadWantsMale(isChildMalePtr)
	time.Sleep(4 * time.Millisecond)
	showChildGender(isChildMalePtr)
}

func momWantsFemale(isChildMalePtr *bool) {
	for i := 0; i < 3; i++ {
		fmt.Println("Nope. I'm the Mom, and I want my child to be a female!")
		*isChildMalePtr = false
		time.Sleep(time.Millisecond)
	}
}

func dadWantsMale(isChildMalePtr *bool) {
	for i := 0; i < 3; i++ {
		fmt.Println("No way. I'm the Dad, and I want my child to be a male!")
		*isChildMalePtr = true
		time.Sleep(time.Millisecond)
	}
}

func showChildGender(isChildMalePtr *bool) {
	gender := "male"
	if *isChildMalePtr {
		gender = "male"
	} else {
		gender = "female"
	}
	fmt.Println("Finally, after much bickering, the child's gender turns out to be: " + gender + "! :)")
}
