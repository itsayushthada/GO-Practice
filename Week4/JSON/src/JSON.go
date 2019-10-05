package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	var name, addr string
	fmt.Printf("Enter your Name: ")
	_, _ = fmt.Scanln(&name)
	fmt.Printf("Enter the Address: ")
	_, _ = fmt.Scanln(&addr)
	m["name"] = name
	m["address"] = addr

	json_str, _ := json.Marshal(m)
	fmt.Println("JSON String is: ", string(json_str))
}
