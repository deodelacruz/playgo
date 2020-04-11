/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
)

var requestForTicketChnl chan int
var grantATicketChnl chan mealGrant

func main() {

	// create 5 chopstick mutex
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	fmt.Printf("Philosophers instantiated: %v", CSticks)

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, false, CSticks[i], CSticks[(i+1)%5]}
	}
	fmt.Printf("Philosophers instantiated: %v", philos)

	/*mealTickets := make([]*mealTicket, 2)
	for i := 0; i < 2; i++ {
		mealTickets[i] = &mealTicket{i, new(sync.Mutex)}
	}
	fmt.Printf("Meal tickets instantiated: %v", mealTickets)*/

	isMealTix1Avail := true
	isMealTix2AVail := true

	requestForTicketChnl := make(chan int)   // philo sends their id to host
	grantATicketChnl := make(chan mealGrant) // host provides mealticket here

	/*
	   hostProcessMealTickets()
	   for i:0;i<5; i++ {
	    go philos[i].eat()
	   }
	*/
}

// investigate use of channel between host/server and clients/philos
func hostProcessMealTickets() {
	hostGrantsMealTicket()
	//hostGetsBackMealTicket()
}

// investigate use of channel between host/server and clients/philos

// host lets only 2 philos eat at a time
func hostGrantsMealTicket() {
	// monitor channel of requests for meal ticket from philo
	var requestingPhiloId int
	requestingPhiloId = <-requestForTicketChnl
	fmt.Printf("Host: Received meal request from philosopher %v.\n", requestingPhiloId)
	/* try to grant ticket if avail to requestor
	for _,mealTicket := range mealTickets {
	  if isMealTix1Avail {   // if mealTix1 is available, lease it out
	    grantATicketChnl <- mealTickets[0]
	    fmt.Println("Host: Leasing out meal ticket 1.")
	    isMealTix1Avail = false
	    return mealTickets[0]
	  } else isMealTix2Avail {   // else if mealTix2 is availalbe, lease it
	    grantATicketChnl <- mealTickets[0]
	    fmt.Println("Host: Leasing out meal ticket 2.")
	    isMealTix1Avail = false
	    return mealTickets[1]
	  }
	} */
}

/*
func hostGetsBackMealTicket(mealTix *sync.Mutext)  {
  // set mealTix to avail
  mealTix.Lock()
} */

// create 5 chopstick mutex
type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id                  int
	isGrantedMealTicket bool
	leftCS, rightCS     *ChopS
}

// create 2 mealticket mutex
type mealTicket struct {
	id  int
	mtx sync.Mutex
}

type mealGrant struct {
	philoId      int
	mealTicketId int
}

func (p Philo) eat() {
	for {
		// check if meal ticket was granted to this philo by host
		myMealGrant := <-grantATicketChnl
		//
		if p.id == myMealGrant.philoId {
			fmt.Printf("Meal ticket %v granted to me. philo%v\n", myMealGrant, p.id)
			/* p.leftCS.Lock()
			   p.rightCS.Lock()
			   fmt.Println("eating")
			   p.rightCS.Unlock()
			   p.leftCS.Unlock()
			    p.mealTix.Lock() */
		} else {
			fmt.Printf("Meal ticket %v NOT granted to me. philo%v\n", myMealGrant, p.id)
		}
	}

}
