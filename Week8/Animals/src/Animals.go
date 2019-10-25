package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

/*
	Methods & Data for Cow
*/
type Cow struct {
	food       string
	locomotive string
	noise      string
}

func (c Cow) Eat()   { fmt.Printf("%s\n", c.food) }
func (c Cow) Move()  { fmt.Printf("%s\n", c.locomotive) }
func (c Cow) Speak() { fmt.Printf("%s\n", c.noise) }

/*
	Methods & Data for Bird
*/
type Bird struct {
	food       string
	locomotive string
	noise      string
}

func (b Bird) Eat()   { fmt.Printf("%s\n", b.food) }
func (b Bird) Move()  { fmt.Printf("%s\n", b.locomotive) }
func (b Bird) Speak() { fmt.Printf("%s\n", b.noise) }

/*
	Methods & Data for Snake
*/
type Snake struct {
	food       string
	locomotive string
	noise      string
}

func (s Snake) Eat()   { fmt.Printf("%s\n", s.food) }
func (s Snake) Move()  { fmt.Printf("%s\n", s.locomotive) }
func (s Snake) Speak() { fmt.Printf("%s\n", s.noise) }

/*
	Driver Function
*/
func main() {
	var action, attr1, attr2 string
	Database := make(map[string]Animal)

	fmt.Println("Enter X to exit.")
	for {
		fmt.Printf("Enter the Command: > ")
		fmt.Scanf("%s %s %s\n", &action, &attr1, &attr2)

		action = strings.ToLower(action)
		attr1 = strings.ToLower(attr1)
		attr2 = strings.ToLower(attr2)

		if strings.Compare(action, "x") == 0 {
			break
		} else {
			switch action {
			case "newanimal":
				switch attr2 {
				case "cow":
					Database[attr1] = Cow{"grass", "walk", "moo"}
					fmt.Println("Created it!")
				case "bird":
					Database[attr1] = Bird{"worms", "fly", "peep"}
					fmt.Println("Created it!")
				case "snake":
					Database[attr1] = Snake{"mice", "slither", "hsss"}
					fmt.Println("Created it!")
				default:
					fmt.Println("Animal Not Found.")
				}
			case "query":
				_, ifExist := Database[attr1]
				if ifExist {
					switch attr2 {
					case "eat":
						Database[attr1].Eat()
					case "move":
						Database[attr1].Move()
					case "speak":
						Database[attr1].Speak()
					default:
						fmt.Println("Action not Found")
					}
				} else {
					fmt.Println("Animal Not Found. Create it First.")
				}
			default:
				fmt.Printf("%s not found.\n", action)
			}
		}
	}
}
