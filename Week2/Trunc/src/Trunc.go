package main

import "fmt"

func main(){
	var digit float64
	fmt.Printf("Enter any number: ")
	_, err := fmt.Scanf("%f", &digit)
	if err != nil{
		fmt.Printf("[ERROR]: %d", err)
	} else{
		fmt.Printf("Truncated number is %d", int64(digit))
	}
}

