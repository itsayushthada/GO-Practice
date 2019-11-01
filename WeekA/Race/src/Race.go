/*
 Race Condition: It's a situation in which the result of a process depends on the ordering of
 		 the instrunction in which they gets executed. Dur to this the the result of the
		 program or machine becaome stochastic when we are expecting a deterministic
		 output. This happends when the steps of a task are depended on each other and
		 they are executed in parallel.

		 The following program demonstrates the race condition. It prints output output
		 of function executed 10 different times.
*/

package main

import (
	"fmt"
	"sync"
)

func increment(a *int){
	/*
	Function reponsible for incrementing the value of an integer variable.
	It's simply updating the value of the integer in the memory.
	*/
	*a = *a + 1
}

func parallel_exec(a *int, wg *sync.WaitGroup){
	/*
	Function reponsible for 10 Parallel increments to an integer variable.
	Due to absense of any syncronization construct there is possibility of
	Race condition and we can expect to see different results on different
	executions of this function.
	*/
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	go increment(a)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	a := 0
	fmt.Printf("Expected Result:\n")
	fmt.Printf("Value of A prior to increment: %d\n", a)
	fmt.Printf("Value of A after increments: %d\n\n", a+10)

	for i:=0; i<10; i++{
		a = 0
		fmt.Printf("Run %d:\n", i+1)
		fmt.Printf("Value of A prior to increment: %d\n", a)

		/* I'm putting syncronization construct to make sure all threads of 
		parallel_exec() function completes there execution and there is no
		pre-emption*/
		wg.Add(1)
		parallel_exec(&a, &wg)
		wg.Wait()

		fmt.Printf("Value of A after increments: %d\n\n", a)
	}
}
