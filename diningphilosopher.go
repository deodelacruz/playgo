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
import ("fmt")

func main() {

    requestForTicketChnl := make(chan int) // philo sends their id to host
    grantATicketChnl := make(chan mealGrant) // host provides mealticket here
    philos := make([]*Philo, 5)
    for i:=0;i<5; i++ {
       philos[i] = &Philo{i,false,CSticks[i],CSticks[(i+1)%5]}
    }
    fmt.Printf("Philosophers instantiated: %v", philos)
    hostProcessMealTickets()
    for i:0;i<5; i++ {
     go philos[i].eat()
    }

}

// create 2 mealticket mutex
type mealTicket struct {
  id int
  mtx sync.Mutex
}

mealTickets := make([]*mealTicket,2)
for i:=0;i<2; i++ {
  mealTickets[i] = new(mealTicket)
}

type mealGrant struct {
  philoId int
  mealTicketId int
}

isMealTix1Avail := true
isMealTix2AVail := true

// investigate use of channel between host/server and clients/philos
func hostProcessMealTickets() {
  hostGrantsMealTicket()
  //hostGetsBackMealTicket()
}

// investigate use of channel between host/server and clients/philos

// host lets only 2 philos eat at a time
func hostGrantsMealTicket() {
  // monitor channel of requests for meal ticket from philo
  philoId <-requestForTicketChnl
  requestingPhiloId := requestForTicketChnl
  fmt.Println("Host: Received meal request from philosopher." + philoId)
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
type ChopS struct {sync.Mutex}
CSticks := make([]*Chops,5)
for i:=0;i<5; i++ {
  CSticks[i] = new(ChopS)
}

type Philo struct {
  id int
  isGrantedMealTicket bool
  leftCS, rightCS *ChopsS
}

func (p Philo) eat() {
 for {
   // check if meal ticket was granted to this philo by host
   mealGrant <- grantATicketChnl
   //
   if p.id == mealGrant.philoId {
     fmt.Println("Meal ticket granted. ")
     /* p.leftCS.Lock()
     p.rightCS.Lock()
     fmt.Println("eating")
     p.rightCS.Unlock()
     p.leftCS.Unlock()
      p.mealTix.Lock() */
    } else {
           fmt.Println("Meal ticket NOT granted to me. philo" + p.id)
    }
}



}
