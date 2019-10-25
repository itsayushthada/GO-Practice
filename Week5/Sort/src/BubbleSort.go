package main

import (
	    "fmt"
	        "math/rand"
		    "time"
	)

func InputData(slice *[]int, met int){
	/*
	This function populates the slice with different methods.
	Methods 0: Populate array with random numbers.
	Methods 1: Populate array with user defined inputs.
	*/
	var temp int

	if met == 0{
		rand.Seed(time.Now().UnixNano())
		for i:=0; i<10 ; i++{
			temp = rand.Intn(10000)
			temp = temp - 2000 /* It'll generate 20% -ve numbers and 80 % +ve numbers. */
			*slice = append(*slice, temp)
		}
	} else if met == 1{
		for i:=0; i<10 ; i++{
			fmt.Scanf("%d\n", &temp)
			*slice = append(*slice, temp)
		}
	}
}

func Swap(slice []int, index int){
	/*
	This function swaps the value of two integer type variables.
	*/
	c := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = c
}

func BubbleSort(slice []int){
	/*
	This function sorts the value of a slice in ascending order.
	Time complexity of the function is O(n^2).
	Space complexity if function is O(n).
	*/
	for i:=0; i<len(slice)-1; i++{
		for j:=0; j<len(slice)-i-1; j++{
			if (slice[j] > slice[j+1]){
				Swap(slice, j)
			}
		}
	}
}

func main(){
	/*
	This is our driver function which encapsualte all the functional blocks in program.
	*/
	fmt.Printf("Enter list of number to be sorted:\n")
	var data []int
	InputData(&data, 0)
	BubbleSort(data)
	fmt.Println(data)
}
