/*
	Dining Philosopher Problem With Some Additional Restrictions.
	Submitted By: Ayush Thada
*/
package main

import(
	"fmt"
	"sync"
)
var nPHILOS_MAP = make(map[int]float64)

type Chopstick struct{
	sync.RWMutex
}

type Philosopher struct{
	id int
	leftCS, rightCS *Chopstick
}

func (p Philosopher) Eat(meals int, include chan int, exclude chan int, mx *sync.Mutex, wg *sync.WaitGroup){
	/*
	This function simulates the eating action for a duration of time.
	The choice varaible allow to recoed the data of every successful attempt to eat
	using chopsticks in the given duration of time.
	*/
	for i:=0; i<meals; i++{
		include <- p.id
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("Starting to Eat %d\n", p.id)
		fmt.Printf("Finishing Eating %d\n", p.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		exclude <- p.id
		mx.Lock()
	}
	wg.Done()
}

func PartyOver(dinner_over chan int, wg *sync.WaitGroup){
	wg.Wait()
	dinner_over <- 1
}

func Host(n_philos int, meals int, max_limit int){
	
	/* Important Slices Initialization */
	CS := make([]*Chopstick, n_people)
	Philo := make([]*Philosopher, n_people)

	for i:=0 ; i<n_people ; i++{
		CS[i] = new(Chopstick)
	}

	for i:=0 ; i<n_people ; i++{
		Philo[i] = &Philosopher{i+1, CS[i], CS[(i+1)%n_people]}
	}
	
	/* Syncronization and Communication variables*/
	var wg sync.WaitGroup
	include := make(chan int)
	exclude := make(chan int)
	dinner_over := make(chan int)
	Moderator := make([]*sync.Mutex, n_people)
	
	/* Starting the Party */
	wg.Add(n_philos)
	for i:=0; i<n_philos ; i++ {
		go Philo[i].Eat(meals, include, exclude, &Moderator[i], &wg)
	}
	
	/* Managment by the Host */
	go PartyOver(dinner_over, &wg)
}

func main(){
	/* Hyper-Parameters of the Program */
	n_philos := 5
	meals := 3
	max_limit := 2
	
	host(n_philos, meals, max_limit)
}
