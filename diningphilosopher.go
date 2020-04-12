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
	//fmt.Printf("Chopsticks instantiated: %v\n", CSticks)

	maxPhilos = 5
	philos = make([]*Philo, maxPhilos)
	numPhilosFull = 0
	for i := 0; i < maxPhilos; i++ {
		philos[i] = &Philo{i, false, CSticks[i], CSticks[(i+1)%5], make(chan int)}
	}
	//fmt.Printf("Philosophers instantiated: %v\n", philos)

	requestForTicketChnl = make(chan int, 1) // philo sends their id to host to request meal ticket

	//fmt.Println("Table host now serving meal tickets.")
	maxMealTickets := 2
	for i := 0; i < maxMealTickets; i++ { //spawn threads to lease 2 meal tickets
		wg.Add(1)
		go hostLeasesMealTicket(i)
	}
	time.Sleep(1 * time.Millisecond)

	// spawn threads for x number of philosophers
	for i := 0; i < maxPhilos; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
	//time.Sleep(5 * time.Millisecond)
	fmt.Println("Done with func main().")
}

func hostLeasesMealTicket(tixId int) {
	fmt.Printf("Host%v now serving meal ticket %v.\n", tixId, tixId)
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
				fmt.Printf("Host%v: Request for meal tickets channel was already closed: %v\n", tixId, requestForTicketChnl)
				break
			}
			fmt.Printf("Host%v: Received meal request from Philosopher%v.\n", tixId, requestingPhiloId)
			// try to grant 1 meal ticket if any avail to requestor
			// Note: also check here if requesting philoshoper is already full
			if isMealTixAvail && !philos[requestingPhiloId].isAlreadyFull {
				fmt.Printf("Host%v: Leasing out meal ticket %v to Philosopher%v.\n", tixId, tixId, requestingPhiloId)
				philos[requestingPhiloId].hostCommsCh <- tixId
				fmt.Printf("Host%v: Successfully sent out meal ticket %v via channel %v to Philosopher%v.\n", tixId, tixId, philos[requestingPhiloId].hostCommsCh, requestingPhiloId)
				isMealTixAvail = false
				//wait for mealticket to be returned
				fmt.Printf("Host%v: Waiting for meal ticket %v to be returned on channel %v by Philosopher%v.\n", tixId, tixId, philos[requestingPhiloId].hostCommsCh, requestingPhiloId)
				mealTixid, ok2 = <-philos[requestingPhiloId].hostCommsCh
				if ok2 == false {
					fmt.Printf("Host%v: Send meal tickets back channel was closed: %v\n", tixId, philos[requestingPhiloId].hostCommsCh)
					break
				}
				if mealTixid == tixId {
					isMealTixAvail = true
					fmt.Printf("Host%v: Received back meal ticket %v.\n", tixId, mealTixid)
				}
			} else {
				fmt.Printf("Host%v: Sorry, Philosopher%v. Meal ticket %v is not available.\n", tixId, requestingPhiloId, tixId)
			}

		} else {
			fmt.Printf("Host%v: All philos are full, terminating serving of leases for meal ticket %v.\n", tixId, tixId)
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
	hostCommsCh     chan int
}

func (p Philo) eat() {
	maxTimesEat := 3 // max number of times philosopher can eat before full
	numTimesEat := 0
	for {
		if numTimesEat < maxTimesEat {
			//send request to eat to host via common channel
			requestForTicketChnl <- p.id
			fmt.Printf("Philosopher%v: Sent request to host for meal ticket.\n", p.id)

			// check if meal ticket was granted to this philo by host
			fmt.Printf("Philosopher%v: Waiting for meal ticket on channel %v...\n", p.id, p.hostCommsCh)
			grantedMealTixId, ok := <-p.hostCommsCh //dedicated channel to host here
			if ok == false {
				fmt.Printf("Grant meal tickets channel was closed: %v\n", p.hostCommsCh)
				break //hmm, check this
			}
			fmt.Printf("Philosopher%v: Received meal ticket %v.\n", p.id, grantedMealTixId)
			/* p.leftCS.Lock()
			   p.rightCS.Lock()
			   fmt.Println("eating")
			   p.rightCS.Unlock()
			   p.leftCS.Unlock()
			    p.mealTix.Lock() */
			time.Sleep(1 * time.Microsecond)
			numTimesEat++
			fmt.Printf("Philosopher%v: Done eating. Returning back meal ticket %v on channel %v....\n", p.id, grantedMealTixId, p.hostCommsCh)
			p.hostCommsCh <- grantedMealTixId
			fmt.Printf("Philosopher%v: Successfully sent back meal ticket %v on channel %v....\n", p.id, grantedMealTixId, p.hostCommsCh)
		} else {
			fmt.Printf("Philosopher%v: Done eating max number of times %v.\n", p.id, numTimesEat)
			p.isAlreadyFull = true
			break
		}
		time.Sleep(1 * time.Microsecond) // sleep a bit before requesting new ticket again
	} //for
	mx.Lock()
	numPhilosFull++
	mx.Unlock()
	//fmt.Printf("Philosopher%v: sees numPhilosFull:%v\n", p.id, numPhilosFull)
	if numPhilosFull == maxPhilos {
		fmt.Printf("Philosopher%v: Closing tix request channel as all philosophers are now full: %v\n", p.id, numPhilosFull)
		close(requestForTicketChnl) // close tixRequest channel when all philos are full
	}
	wg.Done()
}
