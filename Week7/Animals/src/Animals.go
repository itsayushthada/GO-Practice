package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotive string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotive
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	var class, action string

	// Knowledge Base
	kb := make(map[string]Animal)
	kb["cow"] = Animal{"grass", "walk", "moo"}
	kb["bird"] = Animal{"worms", "fly", "peep"}
	kb["snake"] = Animal{"mice", "slither", "hsss"}

	// Driver Code
	fmt.Println("Enter X to exit.")
	for {
		fmt.Printf("\n> ")
		fmt.Scanf("%s %s\n", &class, &action)

		if strings.Compare(class, "x") == 0 || strings.Compare(action, "x") == 0 {
			break
		} else {
			fmt.Printf("User Requested '%s' action for '%s' results: ", action, class)
			action = strings.ToLower(action)
			class = strings.ToLower(class)

			_, ifExist := kb[class]
			if ifExist {
				switch action {
				case "eat":
					fmt.Printf("%s \n", kb[class].Eat())
				case "move":
					fmt.Printf("%s \n", kb[class].Move())
				case "speak":
					fmt.Printf("%s \n", kb[class].Speak())
				default:
					fmt.Printf("Action Not found in Knowledge Base.\n")
				}
			} else {
				fmt.Printf("Animal Not Found in knowledge Base.\n")
			}
		}
	}
}
