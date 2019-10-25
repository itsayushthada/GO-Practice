// Write a program which prompts the user to first enter a name, and then enter
// an address. Your program should create a map and add the name and address to
// the map using the keys “name” and “address”, respectively. Your program
// should use Marshal() to create a JSON object from the map, and then your
// program should print the JSON object.

// Developed by Antonio BOSNJAK

package main
import (
    "encoding/json"
    "bufio"
    "fmt"
    "os"
)

func main() {
  var Name1 string
  var Address1 string

  my_map := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Prease enter a name: ")

	scanner.Scan()
	Name1 = scanner.Text()

	fmt.Print("Please enter an address: ")
	scanner.Scan()
	Address1 = scanner.Text()

  // Save on My map
	my_map["name"] = Name1
	my_map["address"] = Address1

	// Marshal to json
	jsonString, err := json.Marshal(my_map)

	// Print json data
	if err == nil {
		fmt.Println("Contact entered: ", string(jsonString))
	} else {
		fmt.Println("ERROR: Failed to marshal the input provided. ", err)
	}
	fmt.Println("Done.")
  fmt.Printf("\n")
}
