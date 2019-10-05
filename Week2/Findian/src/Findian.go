package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a String: ")
	input, _ = reader.ReadString('\n')
	input = strings.ToLower(input)
	var length int = len(input);
	
	if input[0] != 'i'|| input[length-1] != 'n'{
		fmt.Printf("Not Found")
	} else {
		for i:=1; i<length-1; i++{
			if input[i] == 'a'{
				fmt.Printf("Found")
				return
			}
		}
		fmt.Printf("Not Found")
	}
}
