/*
	Dining Philosopher with following constraints:
	1. There should be 5 philosophers sharing chopsticks, with one chopstick between
	   each adjacent pair of philosophers.
	2. Each philosopher should eat only 3 times (not in an infinite loop)
	3. The philosophers pick up the chopsticks in any order, not lowest-numbered first.
	4. In order to eat, a philosopher must get permission from a host which executes in
	   its own goroutine.
	5. The host allows no more than 2 philosophers to eat concurrently.
	6. Each philosopher is numbered, 1 through 5.
	7. When a philosopher starts eating (after it has obtained necessary locks) it prints
	   “starting to eat <number>” on a line by itself, where <number> is the number of the
	   philosopher.
	8. When a philosopher finishes eating (before it has released its locks) it prints
	   “finishing eating <number>” on a line by itself, where <number> is the number of the
	   philosopher.

	Submitted By: Ayush Thada
*/

package main

import (
	"fmt"
	"sync"
)

type ChopStick struct {
	sync.RWMutex
}

type Philosopher struct {
	leftCS, rightCS *ChopStick
	name            int
}

func (p Philosopher) eat(meals int, add chan int, remove chan int, wg *sync.WaitGroup, mx *sync.Mutex) {
	/*
		This function simulates the eating action for a duration of time.
		The choice varaible allow to recoed the data of every successful attempt to eat
		using chopchopstick in the given duration of time.
	*/

	for i := 0; i < meals; i++ {
		add <- p.name
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("Starting To Eat %d\n", p.name)
		fmt.Printf("Finishing Eating %d\n", p.name)
		remove <- p.name
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		mx.Lock()
	}
	wg.Done()
}

func dinnerOver(abort chan int, wg *sync.WaitGroup) {
	wg.Wait()
	abort <- -1
}

func Host(n_philos int, limit int, meals int) {

	/* Syncronization, Communication and Data storage variables */
	var (
		input     int
		wg        sync.WaitGroup
		count     = 0
		waitlist  = make([]int, 0)
		moderator = make([]sync.Mutex, n_philos)
		philos    = make([]Philosopher, 0)
		chopstick = make([]ChopStick, n_philos)
		add       = make(chan int)
		remove    = make(chan int)
		over      = make(chan int)
	)

	/* Starting the Party */
	wg.Add(n_philos)
	for i := 0; i < n_philos; i++ {
		philos = append(philos, Philosopher{&chopstick[i], &chopstick[(i+1)%n_philos], i})
		moderator = append(moderator, sync.Mutex{})
		moderator[i].Lock()
		go philos[i].eat(meals, add, remove, &wg, &moderator[i])
	}

	/* Managment by the Host */
	go dinnerOver(over, &wg)

	for {
		select {
		case input = <-add:
			if count < limit {
				count += 1
				moderator[input].Unlock()
			} else {
				waitlist = append(waitlist, input)
			}
		case input = <-remove:
			if len(waitlist) != 0 {
				moderator[waitlist[0]].Unlock()
				waitlist = waitlist[1:]
			} else {
				count -= 1
			}
		case <-over:
			fmt.Printf("\nDinner Over!!!")
			return
		default:
		}
	}
}

func main() {
	/* Parameters of the Program */
	n_philos := 5
	limit := 2
	meals := 3

	Host(n_philos, limit, meals)
}
