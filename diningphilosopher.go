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
	"time"
)

var requestForTicketChnl chan int
var grantATicketChnl chan *mealGrant
var receiveMealTicketBackChnl chan int
var mealTickets []*mealTicket
var numPhilosFull int
var wg sync.WaitGroup
var maxPhilos int
var mx sync.Mutex
var philos []*Philo

func main() {

	// create 5 chopstick mutex
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	fmt.Printf("Chopsticks instantiated: %v\n", CSticks)

	philos = make([]*Philo, 5)
	numPhilosFull = 0
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, false, CSticks[i], CSticks[(i+1)%5]}
	}
	fmt.Printf("Philosophers instantiated: %v\n", philos)

	mealTickets = make([]*mealTicket, 2)
	for i := 0; i < 2; i++ {
		//mealTickets[i] = &mealTicket{i, new(sync.Mutex)}
		mealTickets[i] = &mealTicket{i}
	}
	fmt.Printf("Meal tickets instantiated: %v\n", mealTickets)

	requestForTicketChnl = make(chan int)      // philo sends their id to host to request meal ticket
	grantATicketChnl = make(chan *mealGrant)   // host provides mealticket here
	receiveMealTicketBackChnl = make(chan int) // philo sends back meal ticket here when done eating

	fmt.Println("Table host now serving meal tickets.")
	for i := 0; i < 2; i++ { //spawn threads to lease 2 meal tickets
		wg.Add(1)
		go hostLeasesMealTicket(i)
	}

	// spawn threads for x number of philosophers
	maxPhilos = 2
	for i := 0; i < maxPhilos; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
	//time.Sleep(5 * time.Millisecond)
	fmt.Println("Done with func main().")
}

/* investigate use of channel between host/server and clients/philos
func hostProcessMealTickets() {
	go hostGrantsMealTicket()
	go hostGetsBackMealTicket()
} */

/* host lets only 2 philos eat at a time
func hostGrantsMealTicket() {
	fmt.Println("Host now trying to grant meal tickets.")
	// monitor channel of requests for meal ticket from philo
	var requestingPhiloId int
	for {
		requestingPhiloId = <-requestForTicketChnl
		fmt.Printf("Host: Received meal request from philosopher %v.\n", requestingPhiloId)
		// try to grant 1 meal ticket if any avail to requestor
		if isMealTix0Avail { // if mealTix1 is available, lease it out
			mg := &mealGrant{requestingPhiloId, 0}
			fmt.Printf("Host: Leasing out meal ticket 0 to philosopher %v.\n", requestingPhiloId)
			grantATicketChnl <- mg
			isMealTix0Avail = false
			break
		} else if isMealTix1Avail { // else if mealTix2 is availalbe, lease it
			mg := &mealGrant{requestingPhiloId, 1}
			fmt.Printf("Host: Leasing out meal ticket 1 to philosopher %v.\n", requestingPhiloId)
			grantATicketChnl <- mg
			isMealTix1Avail = false
			break
		}
	}
}

func hostGetsBackMealTicket() {
	mealTixid := <-receiveMealTicketBackChnl
	fmt.Printf("Host: Received back meal ticket %v.\n", mealTixid)
	// set mealTix to avail
	if mealTixid == 0 {
		isMealTix0Avail = true
	} else if mealTixid == 1 {
		isMealTix1Avail = true
	}
}
*/

func hostLeasesMealTicket(tixId int) {
	fmt.Printf("Host now serving meal ticket %v.\n", tixId)
	// monitor channel of requests for meal ticket from philo
	isMealTixAvail := true
	var requestingPhiloId, mealTixid int
	var ok, ok2 bool
	for {
		// if all philosophers have fully eaten, terminate this goroutine
		//fmt.Printf("Host for lease %v sees numPhilosFull: %v\n", tixId, numPhilosFull)
		if numPhilosFull < maxPhilos {
			requestingPhiloId, ok = <-requestForTicketChnl // read tix requests from channel while open
			if ok == false {
				fmt.Printf("Request for meal tickets channel was closed: %v\n", requestForTicketChnl)
				break
			}
			//fmt.Printf("Host: Received meal request from Philosopher%v.\n", requestingPhiloId)
			// try to grant 1 meal ticket if any avail to requestor
			// Note: also check here if requesting philoshoper is already full
			if isMealTixAvail && !philos[requestingPhiloId].isAlreadyFull {
				mg := &mealGrant{requestingPhiloId, tixId}
				fmt.Printf("Host: Leasing out meal ticket %v to Philosopher%v.\n", tixId, requestingPhiloId)
				// Don't use shared channel below, just give ticket directly to requesting philosopher
				grantATicketChnl <- mg
				isMealTixAvail = false
				//wait for mealticket to be returned
				mealTixid, ok2 = <-receiveMealTicketBackChnl
				if ok2 == false {
					fmt.Printf("Send meal tickets back channel was closed: %v\n", receiveMealTicketBackChnl)
					break
				}
				if mealTixid == tixId {
					isMealTixAvail = true
					fmt.Printf("Host: Received back meal ticket %v.\n", mealTixid)
				}
			} else {
				fmt.Printf("Host: Sorry, Philosopher%v. Meal ticket %v is not available.\n", requestingPhiloId, tixId)
			}

		} else {
			fmt.Printf("Host: All philos are full, terminating serving of leases for meal ticket %v.\n", tixId)
			close(grantATicketChnl) // no need to continue granting tickets to channel
			break
		} // if all philos
	} // for loop
	wg.Done()
}

// create 5 chopstick mutex
type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id              int
	isAlreadyFull   bool
	leftCS, rightCS *ChopS
}

// create 2 mealticket mutex
type mealTicket struct {
	id int
	//mtx sync.Mutex
}

type mealGrant struct {
	philoId      int
	mealTicketId int
}

func (p Philo) eat() {
	maxTimesEat := 2 // max number of times philosopher can eat before full
	numTimesEat := 0
	for {
		if numTimesEat < maxTimesEat {
			//send request to eat to host via channel
			requestForTicketChnl <- p.id
			//fmt.Printf("Philosopher%v: Sent request to host for meal ticket\n", p.id)
			// check if meal ticket was granted to this philo by host
			//fmt.Printf("Philosopher%v: Waiting for meal ticket\n", p.id)
			myMealGrant, ok := <-grantATicketChnl
			if ok == false {
				fmt.Printf("Grant meal tickets channel was closed: %v\n", receiveMealTicketBackChnl)
				break
			}
			if p.id == myMealGrant.philoId {
				fmt.Printf("Philosopher%v: Received meal ticket %v.\n", p.id, myMealGrant.mealTicketId)
				/* p.leftCS.Lock()
				   p.rightCS.Lock()
				   fmt.Println("eating")
				   p.rightCS.Unlock()
				   p.leftCS.Unlock()
				    p.mealTix.Lock() */
				time.Sleep(1 * time.Microsecond)
				numTimesEat++
				fmt.Printf("Philosopher%v: Done eating. Returning back meal ticket %v.\n", p.id, myMealGrant.mealTicketId)
				// Note: line below can cause deadlock
				receiveMealTicketBackChnl <- myMealGrant.mealTicketId
			} else {
				// philo, immediately return ticket  back to grantATicketChnl if not meant for you
				fmt.Printf("Philosopher%v: Oops, can't eat. Meal ticket %v is not meant for me.\n", p.id, myMealGrant.mealTicketId)
				grantATicketChnl <- myMealGrant
			}
		} else {
			fmt.Printf("Philosopher%v: Done eating max number of times %v.\n", p.id, numTimesEat)
			p.isAlreadyFull = true
			break
		}
		//	time.Sleep(1 * time.Microsecond) // sleep before requesting new ticket
	} //for
	mx.Lock()
	numPhilosFull++
	mx.Unlock()
	//fmt.Printf("Philosopher%v: sees numPhilosFull:%v\n", p.id, numPhilosFull)
	if numPhilosFull == maxPhilos {
		fmt.Printf("Philosopher%v: Closing tix request channel as all philosophers are full: %v\n", p.id, numPhilosFull)
		close(requestForTicketChnl) // close tixRequest and sendMealTixBack channel when all philos are full
		close(receiveMealTicketBackChnl)
	}
	wg.Done()
}
