/*
	Dining Philosopher Problem
	Submitted By: Ayush Thada
*/
package main

import(
	"fmt"
	"sync"
	"time"
	"math"
)
var nPHILOS_MAP = make(map[int]float64)

type Chopstick struct{
	id int
	sync.RWMutex
}

type Philosopher struct{
	id int
	leftCS, rightCS *Chopstick
}

func (p Philosopher) Eat(clk time.Duration, choice int, wg *sync.WaitGroup){
	/*
	This function simulates the eating action for a duration of time.
	The choice varaible allow to recoed the data of every successful attempt to eat
	using chopsticks in the given duration of time.
	*/

	for start := time.Now(); time.Since(start) < clk; {
		p.leftCS.Lock()
		p.rightCS.Lock()
		if choice == 1{
			nPHILOS_MAP[p.id] = nPHILOS_MAP[p.id] + 1
		}
		fmt.Printf("Philosopher %d is Eating With Chopstics %d and %d\n", p.id, p.leftCS.id, p.rightCS.id)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func Strategy(p_num int, n_philos, choice int) (int, int){
	/*
	This function allow us to swich between various strategies in which chopstics
	has to be picked by Philosophers so that there is no Deadlock in the situation.
	The secondry motive of the task is to reduce the starvation of philosophers.
	[Note]: Other strategies will be added soon.
	*/

	switch(choice){
		/* 
		Naive Strategy: Pick Left and then Right Chopstick.
		*/
		case 1:
			return p_num, (p_num+1)%n_philos

		/*
		Dijkstra Method: Pick the Chopstic with smallest id First.
		*/
		case 2:
			if p_num == n_philos-1{
				return (p_num+1)%n_philos, p_num
			}else{
				return p_num, (p_num+1)%n_philos
			}

		default:
			return p_num, (p_num+1)%n_philos
	}
}

func Initilialize(n_events int, strategy int) []*Philosopher{
	/*
	This function creates the Philosophers, Chopstics and the table to record the
	data of the experiment for studying starvation.
	*/

	CS := make([]*Chopstick, n_events)
	Philo := make([]*Philosopher, n_events)

	for i:=0 ; i<n_events ; i++{
		CS[i] = new(Chopstick)
	}

	for i:=0 ; i<n_events ; i++{
		nPHILOS_MAP[i] = 0
		idx1, idx2 := Strategy(i, n_events, strategy)
		CS[idx1].id = idx1
		CS[idx2].id = idx2
		Philo[i] = &Philosopher{i, CS[idx1], CS[idx2]}
	}

	return Philo
}

func normalize(dict map[int]float64, precision float64){
	/*
	This function normalizes the count of successfull attempt to eat by Philosophers
	over a duration of time.
	*/

	var n_total float64 = 0
	fact := math.Pow(10, precision)

	for i:=0 ; i<len(dict) ; i++{
		n_total += dict[i]
	}:w

	for i:=0 ; i<len(dict) ; i++ {
		dict[i] = 100*(dict[i]/n_total)
		dict[i] = math.Floor(dict[i]*fact)/fact
	}
}

func Summary(prec float64){
	/*
	This function prints the Summary or we can say Results of the Experiments.
	*/

	normalize(nPHILOS_MAP, prec)
	fmt.Println("\nFraction of Time Consumed in Eating by Philosophers:")
	for key, val := range nPHILOS_MAP{
		fmt.Printf("Philosopher %d has taken %f %% of Total Time.\n", key, val)
	}
}

func main(){
	var CHOICE int

	for {
		fmt.Println("\n############################################################")
		fmt.Println("\nChoice 1: Starvation Summary.")
		fmt.Println("Choice 2: Infinite Run for 5 Philosophers.")
		fmt.Printf("\nEnter the Choice in Integer: ")
		fmt.Scanf("%d\n", &CHOICE)

		switch CHOICE{

		case 1:
				n_philos := 10
				strategy := 2
				duration := time.Second*5
				var prec float64 = 4

				fmt.Printf("It'll take atleast %d seconds...", n_philos*5)
				time.Sleep(5*time.Second)

				Philos := Initilialize(n_philos, strategy)
				var wg sync.WaitGroup
				wg.Add(n_philos)
				for i:=0 ; i<n_philos ; i++{
					go Philos[i].Eat(duration, choice, &wg)
				}
				wg.Wait()
				Summary(prec)

		case 2:
				n_philos := 5
				strategy := 2
				var duration time.Duration = time.Hour

				Philos := Initilialize(n_philos, strategy)
				var wg sync.WaitGroup

				wg.Add(n_philos)
				for i:=0 ; i<n_philos ; i++{
					go Philos[i].Eat(duration, &wg)
				}
				wg.Wait()

		default:
				fmt.Println("End")
				return
		}
	}
}
