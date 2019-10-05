package main

import (
	"strings"
	"fmt"
)

func IsPattern1(data string) bool{
	input := []rune(data)
	var length int = len(input);
	
	if (input[0] != 'i'|| input[length-1] != 'n'){
		return false
	}
	
	for i:=1; i<length-1; i++{
		if input[i] == 'a'{
			return true
		}
	}
	return false
}

func main() {
	var input string
	
	fmt.Printf("Enter a String: ")
	val, err := fmt.Scanln(&input)
	
	if err != nil{
		fmt.Printf("[ERROR]: Cannot parse your input: %s", err)
	} else if val == 0 {
		fmt.Printf("[ERROR]: please type your string")
	} else {
		input = strings.ToLower(input)
		if IsPattern1(input){
			fmt.Printf("Found")
		} else {
			fmt.Printf("Not Found")
		}
	}
}
