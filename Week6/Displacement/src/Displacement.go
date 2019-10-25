package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InputData(u_init *float64, s_init *float64, acc *float64, met int) {
	/*
		This function iitate the value variables.
		Methods 0: Initiate variables with random numbers.
		Methods 1: initiate variables array with user defined inputs.
	*/
	if met == 0 {
		rand.Seed(time.Now().UnixNano())
		*u_init = (rand.Float64() - 0.5) * 200.0
		*s_init = (rand.Float64() - 0.5) * 200.0
		*acc = (rand.Float64() - 0.5) * 200.0
	} else if met == 1 {
		fmt.Printf("Enter Initial Velocity: ")
		fmt.Scanf("%f\n", u_init)
		fmt.Printf("Enter Initial Displacement: ")
		fmt.Scanf("%f\n", s_init)
		fmt.Printf("Enter Accelaration: ")
		fmt.Scanf("%f\n", acc)
	}
}

func GenerateDisplaceFn(u_init, s_init, acc float64) func(float64) float64 {
	/*
		This function return the function to calculate the displacment.
	*/
	fn := func(t float64) float64 {
		return s_init + u_init*t + 0.5*acc*t*t
	}
	return fn
}

func main() {
	/*
		This is our driver function which encapsualte all the functional blocks in program.
	*/
	var u_init, s_init, acc float64
	InputData(&u_init, &s_init, &acc, 1)
	fn := GenerateDisplaceFn(u_init, s_init, acc)

	var time float64
	fmt.Print("Enter the time for which displacement has to be estimated: ")
	fmt.Scanf("%f\n", &time)
	
	/*
		Final Output of the Program.
	*/
	fmt.Printf("\nInitial Velocoty: %f m/s", u_init)
	fmt.Printf("\nInitial Displacement: %f m", s_init)
	fmt.Printf("\nAccelaration: %f m/s^2", acc)
	fmt.Printf("\nTime of Interest: %f s", time)
	fmt.Printf("\nFinal Displacment: %f m", fn(time))
}
