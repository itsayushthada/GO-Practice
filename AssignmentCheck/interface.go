// by Ardavan Izadiyar @ardawanizadi
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ food, locomtion, speak string }
type Bird struct{ food, locomtion, speak string }
type Snake struct{ food, locomtion, speak string }

var name string
var request string

func printHelp() {
	fmt.Println("\n=====================")
	fmt.Println("Help : Follow one of the options")
	fmt.Println("> NewAnimal or new | AnimalName - Food - Locomtion - Sound")
	fmt.Println("> Query | AnimalName - Request")
	fmt.Println("> Type Help for more information.")
	fmt.Println("> Type Exit or x to end the program.")
	fmt.Println("=====================\n")
}

func main() {
	var info = make(map[string]Animal)

	fmt.Println("Type Help for more info")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		input := strings.Fields(strings.ToLower(scanner.Text()))

		switch input[0] {
		case "new", "newanimal":
			if len(input) < 5 {
				fmt.Println("ERROR: incomplete inputs")
				printHelp()
				continue
			}

			switch strings.ToLower(input[1]) {
			case "cow":
				info[input[1]] = Cow{input[2], input[3], input[4]}
			case "bird":
				info[input[1]] = Bird{input[2], input[3], input[4]}
			case "snake":
				info[input[1]] = Snake{input[2], input[3], input[4]}
			}

			fmt.Println(info)
		case "query":
			if len(input) < 3 {
				fmt.Println("ERROR: incomplete inputs")
				printHelp()
				continue
			} else if len(info) < 1 {
				fmt.Println("ERROR: No data found")
				printHelp()
				continue
			}

			switch input[2] {
			case "eat":
				info[input[1]].Eat()
			case "move":
				info[input[1]].Move()
			case "speak":
				info[input[1]].Speak()
			}

		case "help":
			printHelp()

		case "x", "exit":
			return

		default:
			fmt.Println("ERROR: Wrong option keyword. Choose newAnimal(new) or query only!")
			fmt.Println("Type Help for more info")

		}
	}

}

func (a Cow) Eat()   { fmt.Println(a.food) }
func (a Cow) Move()  { fmt.Println(a.locomtion) }
func (a Cow) Speak() { fmt.Println(a.speak) }

func (a Bird) Eat()   { fmt.Println(a.food) }
func (a Bird) Move()  { fmt.Println(a.locomtion) }
func (a Bird) Speak() { fmt.Println(a.speak) }

func (a Snake) Eat()   { fmt.Println(a.food) }
func (a Snake) Move()  { fmt.Println(a.locomtion) }
func (a Snake) Speak() { fmt.Println(a.speak) }
