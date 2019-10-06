package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	first_name string
	last_name  string
}

func main() {
	var fname string
	var database []Person
	fmt.Printf("Enter File Name: ")
	_, err := fmt.Scanln(&fname)
	if err != nil {
		fmt.Println("[Error]: ", err)
	} else {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println("[ERROR]: ", err)
		} else {
			values := strings.Split(string(data), "\n")
			for _, v := range values {
				name := strings.Split(v, " ")
				database = append(database, Person{first_name: name[0], last_name: name[1]})
			}
		}
	}
	for idx, val := range database {
		fmt.Println(idx+1, val.first_name, val.last_name)
	}

}
